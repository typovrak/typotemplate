package html

import (
	"bytes"
	"slices"
)

// INFO:
// Golang
// pas de regex
// pas de dépendances* (sauf connecteurs type DB ou Prometheus)
// pas de tokenisation/lexer/parsing, tout en une boucle si possible
// pas d'IA dans l'IDE/éditeur de texte

//func Header() string {
//	return "<a title=\"test\">test</a>"
//}

// all state for managing <style></style>
const (
	StyleTagOutside int = iota
	StyleTagOpening
	StyleTagInCSS
)

// all state for managing <script></script>
const (
	ScriptTagOutside int = iota
	ScriptTagOpening
	ScriptTagInJS
)

type TupleInt [2]int

// INFO: HTML comments in HTML comments are forbidden
// INFO: only < are allowed in attribute without separator
// INFO: \' in URL attributes with ' separators are forbidden
// INFO: don't touch to href, href= href="" because this gives different and inconsistent output in JS
// INFO: auto closing tags can be attribute values, chrome doesn't render />, only >: <img src=/> will be <img src="/">
// INFO: > is considered  as a closing HTML tag if no separator is defined: <a href= > ></a> will be <a href> ></a>
func Minifier(html string) string {
	var (
		buf      bytes.Buffer
		char     byte
		lastChar byte

		bufAttrSeparator byte
		repeatedSpaces   TupleInt

		isInComment bool

		isBufInTag     bool
		isBufInAttr    bool
		isBufInURLAttr bool

		bufBytes []byte
		bufLen   int

		// entity encoded representation of "
		// INFO: single quote (') are valid in URL attributes
		entityEncodedDoubleQuote = "&quot;"
		// url encoded representation of "
		// INFO: single quote (') are valid in URL attributes
		urlEncodedDoubleQuote = "%22"

		// all state for managing <style></style>
		// INFO: </style> in CSS tag is forbidden, must be escaped
		styleTagState = StyleTagOutside

		// all state for managing <script></script>
		// INFO: </script> in JS tag is forbidden, must be escaped
		scriptTagState = ScriptTagOutside
		isScriptTagSrc = false

		styleTagOpeningSuffix  = []byte("<style")
		scriptTagOpeningSuffix = []byte("<script")

		hrefAttrSuffix   = []byte("href")
		srcAttrSuffix    = []byte("src")
		actionAttrSuffix = []byte("action")
		dataAttrSuffix   = []byte("data")
	)

	buf.Grow(len(html))

	// TODO: gérer les chevrons dans le contenu?
	// TODO: gérer les balises auto fermante

	// TODO: vu que je retire tous les \n etc, est-ce que je garde les trailing space du contenu ? les doubles espaces ?
	// TODO: retirer les espaces entre les balises html ou pas si aucun texte?

	// TODO: tout paramétrer en bool

	// TODO: benchmark
	// TODO: tester avec vite ou webpack

	// INFO: previous characters = use buf
	// INFO: next characters = use html
	for i := 0; i < len(html); i++ {
		char = html[i]
		bufBytes = buf.Bytes()
		bufLen = buf.Len()

		// count repeated spaces for href attributes
		repeatedSpaces[0] = repeatedSpaces[1]
		if char == ' ' {
			repeatedSpaces[1]++
		} else {
			repeatedSpaces[1] = 0
		}

		// remove HTML comments <!-- -->
		if !isInComment {
			// <!--
			// INFO: check if i+N < len(html) -> with N as maximal number added (this rule prevent every overflow)
			if i+3 < len(html) && char == '<' && html[i+1] == '!' && html[i+2] == '-' && html[i+3] == '-' {
				isInComment = true
				continue
			}
		} else {
			// -->
			// INFO: check if i >= |N| && i < len(html) -> with N as minimal number added (this rule prevent every overflow)
			if i >= 2 && i < len(html) && html[i-2] == '-' && html[i-1] == '-' && char == '>' {
				isInComment = false
			}

			continue
		}

		// manage <style></style>
		switch styleTagState {
		case StyleTagOutside:
			// <style
			if bytes.HasSuffix(bufBytes, styleTagOpeningSuffix) {
				styleTagState = StyleTagOpening
			}

		case StyleTagOpening:
			if !isBufInTag {
				styleTagState = StyleTagInCSS
			}

		case StyleTagInCSS:
			// INFO: manage all CSS code here
			handleStyleInCSS(&buf, &lastChar, char, html, &i, &isBufInTag, &styleTagState)
			continue
		}

		// manage <script></script>
		if scriptTagState == ScriptTagOutside {
			// <script
			if bytes.HasSuffix(bufBytes, scriptTagOpeningSuffix) {
				scriptTagState = ScriptTagOpening
			}
		} else if scriptTagState == ScriptTagOpening {
			// src="x"
			// TODO: rewrite with hasSuffix?
			if bufLen > 6 && bufBytes[bufLen-6] == 's' && bufBytes[bufLen-5] == 'r' && bufBytes[bufLen-4] == 'c' &&
				bufBytes[bufLen-3] == '=' && bufBytes[bufLen-2] == '"' && bufBytes[bufLen-1] != ' ' && bufBytes[bufLen-1] != '"' {
				isScriptTagSrc = true
			}

			if !isBufInTag {
				scriptTagState = ScriptTagInJS

				if isScriptTagSrc {
					// prevent adding a JS character
					continue
				}
			}

		} else if scriptTagState == ScriptTagInJS {
			// INFO: manage all JS code here
			handleScriptInJS(&buf, &lastChar, char, html, &i, &isBufInTag, &scriptTagState, &isScriptTagSrc)
			continue
		}

		// remove line feed, tab and carriage return
		if styleTagState != StyleTagInCSS &&
			scriptTagState != ScriptTagInJS &&
			!isScriptTagSrc &&
			(char == '\n' || char == '\t' || char == '\r') {
			continue
		}

		// start HTML tag
		// INFO: < and > are already handled by the script switch case
		if char == '<' && !isBufInTag {
			isBufInTag = true
			writeByteToBuf(&buf, &lastChar, char)
			continue
		}

		if isBufInTag {
			// remove double space in tag
			if (lastChar == ' ' && char == ' ') || (i+1 < len(html) && char == ' ' && html[i+1] == ' ') {
				continue
			}

			if isBufInAttr {
				if bufAttrSeparator == 0 {
					// attribute separator not defined
					// INFO: attribute value's can be separate by spaces
					if char == ' ' {
						continue
					}

					// only " and ' are allowed
					if char == '\'' || char == '"' {
						bufAttrSeparator = char
						char = '"'
						writeByteToBuf(&buf, &lastChar, '=')
						writeByteToBuf(&buf, &lastChar, char)
						continue

					}

					// > are not allowed without proper separators
					if char == '>' {
						isBufInAttr = false
						writeByteToBuf(&buf, &lastChar, char)
						continue
					}

					bufAttrSeparator = ' '
					writeByteToBuf(&buf, &lastChar, '=')
					writeStrToBuf(&buf, &lastChar, "\""+string(char))
					continue
				}

				if bufAttrSeparator != 0 {
					// remove start trailing space from attribute value
					if lastChar == '"' && char == ' ' {
						continue
					}

					// remove end trailing space from attribute value
					if i+1 < len(html) && char == ' ' && html[i+1] == bufAttrSeparator {
						continue
					}

					// replace all " in attribute value with entity encoded value
					if bufAttrSeparator == '\'' && char == '"' && !isBufInURLAttr {
						writeStrToBuf(&buf, &lastChar, entityEncodedDoubleQuote)
						continue
					}

					// attribute end value
					if bufAttrSeparator == char {
						bufAttrSeparator = 0
						char = '"'
						isBufInAttr = false
						isBufInURLAttr = false
						writeByteToBuf(&buf, &lastChar, char)
						continue
					}

					// handle case where > is in attribute without separator defined
					if bufAttrSeparator == ' ' {
						goNowToNextIteration := handleHTMLTagClosing(&buf, &lastChar, char, &isBufInTag, &isBufInAttr, &isBufInURLAttr, &bufAttrSeparator)
						if goNowToNextIteration {
							continue
						}
					}

					// add for href attritube value the repeated spaces
					if isBufInURLAttr && repeatedSpaces[0] > 1 && bufLen > 1 && bufBytes[bufLen-1] != '"' &&
						char != ' ' && char != bufAttrSeparator {

						spacesToAdd := ""
						for i := 0; i < repeatedSpaces[0]-1; i++ {
							spacesToAdd += " "
						}

						// double quotes in URLs needs to be URL encoded to work and not close the attribute value
						if bufAttrSeparator == '\'' && char == '"' {
							spacesToAdd += urlEncodedDoubleQuote
						} else {
							spacesToAdd += string(char)
						}

						writeStrToBuf(&buf, &lastChar, spacesToAdd)
						continue
					}

					// double quotes in URLs needs to be URL encoded to work and not close the attribute value
					if bufAttrSeparator == '\'' && char == '"' && isBufInURLAttr {
						writeStrToBuf(&buf, &lastChar, urlEncodedDoubleQuote)
						continue
					}

				}

				// attribute value declaration
			} else if char == '=' {
				// > can't be an attribute value's without proper separators like ' and "
				if i+1 < len(html) && html[i+1] == '>' {
					continue
				}

				// URL in href, src, action, data attributes
				if bytes.HasSuffix(bufBytes, hrefAttrSuffix) || bytes.HasSuffix(bufBytes, srcAttrSuffix) ||
					bytes.HasSuffix(bufBytes, actionAttrSuffix) || bytes.HasSuffix(bufBytes, dataAttrSuffix) {

					isBufInURLAttr = true
				}

				isBufInAttr = true
				// TEST:
				// writeByteToBuf(&buf, &lastChar, char)
				continue

				// !isBufInAttr
			} else {
				// remove space at HTML tag start
				if (lastChar == '<' && char == ' ') || (bufLen > 2 && bufBytes[bufLen-2] == '<' && lastChar == '/' && char == ' ') {
					continue
				}

				// remove space at HTML tag end
				if i+1 < len(html) && char == ' ' && html[i+1] == '>' {
					continue
				}

				// handle > character
				goNowToNextIteration := handleHTMLTagClosing(&buf, &lastChar, char, &isBufInTag, &isBufInAttr, &isBufInURLAttr, &bufAttrSeparator)
				if goNowToNextIteration {
					continue
				}
			}
		}

		writeByteToBuf(&buf, &lastChar, char)
	}

	return buf.String()
}

func writeByteToBuf(buf *bytes.Buffer, lastChar *byte, value byte) {
	buf.WriteByte(value)
	*lastChar = value
}

func writeStrToBuf(buf *bytes.Buffer, lastChar *byte, value string) {
	buf.WriteString(value)
	*lastChar = value[len(value)-1]
}

// TODO: typotestcolor: mettre en place une feature pour afficher que les tests du fichier et/ou que de la ligne X à Y

func handleHTMLTagClosing(
	buf *bytes.Buffer,
	lastChar *byte,
	char byte,
	isBufInTag *bool,
	isBufInAttr *bool,
	isBufInURLAttr *bool,
	bufAttrSeparator *byte,

	// INFO: return true for executing goNowToNextIteration and trigger the continue keyword
) bool {
	if char == '>' {
		if *bufAttrSeparator != 0 {
			*bufAttrSeparator = 0
			*isBufInAttr = false
			*isBufInURLAttr = false
			writeByteToBuf(buf, lastChar, '"')
		}

		*isBufInTag = false
		writeByteToBuf(buf, lastChar, char)

		return true
	}

	return false
}

func handleStyleInCSS(
	buf *bytes.Buffer,
	lastChar *byte,
	char byte,
	html string,
	i *int,
	isBufInTag *bool,
	styleTagState *int,
) {
	// can't be </style>
	if char != '<' {
		handleWriteCSS(buf, lastChar, char)
		return
	}

	n := 0
	matching := 0
	styleTagClosingSuffix := []byte("</style>")
	isWritingTagName := false

	for {
		// manage array overflow
		if *i+n >= len(html) {
			break
		}

		// manage spaces
		if html[*i+n] == ' ' {
			// spaces are forbidden between tag name
			if !isWritingTagName {
				n++
				continue
			}

			break
		}

		// not a legal </style> character
		if !slices.Contains(styleTagClosingSuffix[matching:], html[*i+n]) {
			break
		}

		if html[*i+n] == styleTagClosingSuffix[matching] {
			matching++

			if matching < len(styleTagClosingSuffix) {
				switch styleTagClosingSuffix[matching] {
				// tag name's second letter
				case 't':
					isWritingTagName = true

					// tag name's closer
				case '>':
					isWritingTagName = false
				}
			}
		}

		// valid </style> found
		if matching == len(styleTagClosingSuffix) {
			writeStrToBuf(buf, lastChar, "</style>")
			*isBufInTag = true
			*styleTagState = StyleTagOutside

			// for not adding a double > at the end
			*i += n
			return
		}

		n++
	}

	// add CSS character
	handleWriteCSS(buf, lastChar, char)
}

// INFO: add all CSS code here
func handleWriteCSS(buf *bytes.Buffer, lastChar *byte, char byte) {
	writeByteToBuf(buf, lastChar, char)
}

func handleScriptInJS(
	buf *bytes.Buffer,
	lastChar *byte,
	char byte,
	html string,
	i *int,
	isBufInTag *bool,
	scriptTagState *int,
	isScriptTagSrc *bool,
) {
	// can't be </script>
	if char != '<' {
		if !*isScriptTagSrc {
			// add JS character
			handleWriteJS(buf, lastChar, char)
		}
		return
	}

	n := 0
	matching := 0
	scriptTagClosingSuffix := []byte("</script>")
	isWritingTagName := false

	for {
		// manage array overflow
		if *i+n >= len(html) {
			break
		}

		// manage spaces
		if html[*i+n] == ' ' {
			// spaces are forbidden between tag name
			if !isWritingTagName {
				n++
				continue
			}

			break
		}

		// not a legal </script> character
		if !slices.Contains(scriptTagClosingSuffix[matching:], html[*i+n]) {
			break
		}

		if html[*i+n] == scriptTagClosingSuffix[matching] {
			matching++

			if matching < len(scriptTagClosingSuffix) {
				switch scriptTagClosingSuffix[matching] {
				// tag name's second letter
				case 'c':
					isWritingTagName = true

					// tag name's closer
				case '>':
					isWritingTagName = false
				}
			}
		}

		// valid </script> found
		if matching == len(scriptTagClosingSuffix) {
			writeStrToBuf(buf, lastChar, "</script>")
			*isBufInTag = true
			*scriptTagState = StyleTagOutside
			*isScriptTagSrc = false

			// for not adding a double > at the end
			*i += n
			return
		}

		n++
	}

	if !*isScriptTagSrc {
		// add JS character
		handleWriteJS(buf, lastChar, char)
	}
}

// INFO: add all JS code here
func handleWriteJS(buf *bytes.Buffer, lastChar *byte, char byte) {
	writeByteToBuf(buf, lastChar, char)
}
