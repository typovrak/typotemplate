package html

import (
	"bytes"
	"fmt"
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

func writeByteToBuf(buf *bytes.Buffer, lastChar *byte, value byte) {
	buf.WriteByte(value)
	*lastChar = value
}

func writeStrToBuf(buf *bytes.Buffer, lastChar *byte, value string) {
	buf.WriteString(value)
	*lastChar = value[len(value)-1]
}

func Minifier(html string) string {
	var (
		buf      bytes.Buffer
		char     byte
		lastChar byte

		bufAttrSeparator byte
		repeatedSpaces   [2]int

		isInComment bool

		isBufInTag      bool
		isBufInAttr     bool
		isBufInHrefAttr bool

		bufBytes []byte
		bufLen   int
	)

	buf.Grow(len(html))

	// entity encoded representation of "
	entityEncoderDoubleQuote := "&quot;"

	// TODO: gérer les commentaires, commentaire dans un commentaire?
	// TODO: gérer les <script> <style>
	// TODO: gérer les chevrons dans le contenu?

	// TODO: vu que je retire tous les \n etc, est-ce que je garde les trailing space du contenu ? les doubles espaces ?
	// TODO: tout paramétrer en bool
	// TODO: mettre des couleurs dans les tests (rouge et vert)

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

		fmt.Println(isInComment)
		// remove HTML comments
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

		// remove line feed, tab and carriage return
		if char == '\n' || char == '\t' || char == '\r' {
			continue
		}

		// start HTML tag
		if char == '<' {
			isBufInTag = true
			writeByteToBuf(&buf, &lastChar, char)
			continue
		}

		if isBufInTag {
			// remove double space in tag
			if (lastChar == ' ' && char == ' ') || (i+1 < len(html) && char == ' ' && html[i+1] == ' ') {
				continue
			}

			// remove space at HTML tag start
			if (lastChar == '<' && char == ' ') ||
				(bufLen > 2 && bufBytes[bufLen-2] == '<' && lastChar == '/' && char == ' ') {
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
					if bufAttrSeparator == '\'' && char == '"' {
						writeStrToBuf(&buf, &lastChar, entityEncoderDoubleQuote)
						continue
					}

					// add for href attritube value the repeated spaces
					if isBufInHrefAttr && repeatedSpaces[0] > 1 && char != ' ' && char != '\'' && char != '"' {
						// TODO: refacto with buf.Truncate() ?
						spacesToAdd := ""
						for i := 0; i < repeatedSpaces[0]-1; i++ {
							spacesToAdd += " "
						}

						spacesToAdd += string(char)
						writeStrToBuf(&buf, &lastChar, spacesToAdd)
						continue
					}

					// attribute end value
					if bufAttrSeparator == char {
						bufAttrSeparator = 0
						char = '"'
						isBufInAttr = false
						isBufInHrefAttr = false
						writeByteToBuf(&buf, &lastChar, char)
						continue
					}
				}

				// attribute value declaration
			} else if char == '=' {
				if i >= 4 && i < len(html) && html[i-4] == 'h' && html[i-3] == 'r' && html[i-2] == 'e' && html[i-1] == 'f' {
					isBufInHrefAttr = true
				}

				isBufInAttr = true
				writeByteToBuf(&buf, &lastChar, char)
				continue
			}

			// remove space at HTML tag end
			if i+1 < len(html) && char == ' ' && html[i+1] == '>' {
				continue
			}
		}

		if char == '>' {
			if bufAttrSeparator != 0 {
				bufAttrSeparator = 0
				isBufInAttr = false
				isBufInHrefAttr = false
				writeByteToBuf(&buf, &lastChar, '"')
			}

			isBufInTag = false
			writeByteToBuf(&buf, &lastChar, char)
			continue
		}

		// fmt.Println("add this char, no if triggered:", string(char), isBufInAttr, bufAttrSeparator, isBufInHrefAttr, repeatedSpaces)
		writeByteToBuf(&buf, &lastChar, char)
	}

	return buf.String()
}
