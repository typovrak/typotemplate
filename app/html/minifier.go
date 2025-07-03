package html

import (
	"bytes"
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
	StyleTagClosing
)

// all state for managing <script></script>
const (
	ScriptTagOutside int = iota
	ScriptTagOpening
	ScriptTagInJS
	ScriptTagClosing
)

// INFO: HTML comments in HTML comments are forbidden.
func Minifier(html string) string {
	var (
		buf      bytes.Buffer
		char     byte
		lastChar byte

		bufAttrSeparator byte
		repeatedSpaces   [2]int

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
		styleTagState = StyleTagOutside

		// all state for managing <script></script>
		scriptTagState           = ScriptTagOutside
		isScriptTagSrc           = false
		scriptTagClosingMatchLen = 0

		styleTagOpeningSuffix  = []byte("<style")
		styleTagClosingSuffix  = []byte("</style")
		scriptTagOpeningSuffix = []byte("<script")
		scriptTagClosingSuffix = []byte("</script")
		hrefAttrSuffix         = []byte("href")
		srcAttrSuffix          = []byte("src")
		actionAttrSuffix       = []byte("action")
		dataAttrSuffix         = []byte("data")
	)

	buf.Grow(len(html))

	// TODO: gérer les <script> <style>
	// src= action= data=

	// TODO: gérer les balises auto fermante
	// TODO: gérer les chevrons dans le contenu?

	// TODO: vu que je retire tous les \n etc, est-ce que je garde les trailing space du contenu ? les doubles espaces ?
	// TODO: retirer les espaces entre les balises html ou pas si aucun texte?

	// TODO: tout paramétrer en bool
	// TODO: mettre des couleurs dans les tests (rouge et vert)

	// TODO: benchmark

	// WARN: previous characters = use buf
	// WARN: next characters = use html
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
			// </style
			if bytes.HasSuffix(bufBytes, styleTagClosingSuffix) {
				styleTagState = StyleTagClosing
			}

		case StyleTagClosing:
			if !isBufInTag {
				styleTagState = StyleTagOutside
			}
		}

		// manage <script></script>
		if scriptTagState == ScriptTagOutside {
			// <script
			if bytes.HasSuffix(bufBytes, scriptTagOpeningSuffix) {
				scriptTagState = ScriptTagOpening
			}
		} else if scriptTagState == ScriptTagOpening {
			// src="x"
			if bufLen > 6 && bufBytes[bufLen-6] == 's' && bufBytes[bufLen-5] == 'r' && bufBytes[bufLen-4] == 'c' &&
				bufBytes[bufLen-3] == '=' && bufBytes[bufLen-2] == '"' && bufBytes[bufLen-1] != ' ' && bufBytes[bufLen-1] != '"' {
				isScriptTagSrc = true
			}

			if !isBufInTag {
				scriptTagState = ScriptTagInJS

				if isScriptTagSrc {
					goNowToNextIteration := handleScriptTagInJS(&buf, &lastChar, char, &isBufInTag, &scriptTagState, &isScriptTagSrc, &scriptTagClosingMatchLen, scriptTagClosingSuffix)
					if goNowToNextIteration {
						continue
					}
				}
			}

		} else if scriptTagState == ScriptTagInJS {
			goNowToNextIteration := handleScriptTagInJS(&buf, &lastChar, char, &isBufInTag, &scriptTagState, &isScriptTagSrc, &scriptTagClosingMatchLen, scriptTagClosingSuffix)
			if goNowToNextIteration {
				continue
			}

		} else if scriptTagState == ScriptTagClosing {
			if !isBufInTag {
				scriptTagState = ScriptTagOutside
			}
		}

		// remove line feed, tab and carriage return
		if styleTagState != StyleTagInCSS && scriptTagState != ScriptTagInJS && !isScriptTagSrc && (char == '\n' || char == '\t' || char == '\r') {
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
					if char == ' ' {
						continue
					}

					// only " and ' are allowed
					if char == '\'' || char == '"' {
						bufAttrSeparator = char
						char = '"'
						writeByteToBuf(&buf, &lastChar, char)
						continue

					}

					bufAttrSeparator = ' '
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

					// TODO: problème de " ' avec les espaces, voir aussi pour URL encoder les '

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
				// URL in href, src, action, data attributes
				if bytes.HasSuffix(bufBytes, hrefAttrSuffix) || bytes.HasSuffix(bufBytes, srcAttrSuffix) ||
					bytes.HasSuffix(bufBytes, actionAttrSuffix) || bytes.HasSuffix(bufBytes, dataAttrSuffix) {

					isBufInURLAttr = true
				}

				isBufInAttr = true
				writeByteToBuf(&buf, &lastChar, char)
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

func handleScriptTagInJS(
	buf *bytes.Buffer,
	lastChar *byte,
	char byte,
	isBufInTag *bool,
	scriptTagState *int,
	isScriptTagSrc *bool,
	scriptTagClosingMatchLen *int,
	scriptTagClosingSuffix []byte,

	// INFO: return goNowToNextIteration for executing continue keyword
) bool {
	if *isScriptTagSrc {

		// TODO: manage \t ... here?
		// TODO: need to reset if not correct

		if char != ' ' && scriptTagClosingSuffix[*scriptTagClosingMatchLen] == char {
			*scriptTagClosingMatchLen++

			// </script
			if *scriptTagClosingMatchLen == len(scriptTagClosingSuffix) {
				*isScriptTagSrc = false
				*scriptTagState = ScriptTagClosing

				*isBufInTag = true
				// INFO: buf and lastChar are already a pointer
				writeStrToBuf(buf, lastChar, "</script")
			}
		}

		return true
	}

	// </script
	if bytes.HasSuffix(buf.Bytes(), scriptTagClosingSuffix) {
		*scriptTagState = ScriptTagClosing
	}

	return false
}

func handleHTMLTagClosing(
	buf *bytes.Buffer,
	lastChar *byte,
	char byte,
	isBufInTag *bool,
	isBufInAttr *bool,
	isBufInURLAttr *bool,
	bufAttrSeparator *byte,

	// INFO: return goNowToNextIteration for executing continue keyword
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
