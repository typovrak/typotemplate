package html

import (
	"bytes"
)

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
	var buf bytes.Buffer
	buf.Grow(len(html))

	var (
		char             byte
		lastChar         byte
		bufAttrSeparator byte
		repeatedSpaces   [2]int

		isBufInTag      bool
		isBufInAttr     bool
		isBufInHrefAttr bool
		bufBytes        []byte
		bufLen          int
	)

	entityEncoderDoubleQuote := "&quot;"

	// isBufferInComment := true

	// TODO: gérer les commentaires
	// TODO: gérer le doctype
	// TODO: gérer les <script> <style>

	// TODO: vu que je retire tous les \n etc, est-ce que je garde les trailing space du contenu ? les doubles espaces ?
	// TODO: tout paramétrer en bool

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
			if (lastChar == ' ' && char == ' ') || (i+1 <= len(html)-1 && char == ' ' && html[i+1] == ' ') {
				continue
			}

			// remove space at HTML tag start
			if (lastChar == '<' && char == ' ') ||
				(2 < bufLen && bufBytes[bufLen-2] == '<' && lastChar == '/' && char == ' ') {
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
					if i+1 <= len(html)-1 && char == ' ' && html[i+1] == bufAttrSeparator {
						continue
					}

					// replace all " in attribute value with entity encoded value
					if bufAttrSeparator == '\'' && char == '"' {
						writeStrToBuf(&buf, &lastChar, entityEncoderDoubleQuote)
						continue
					}
					// add for href attritube value the repeated spaces
					if isBufInHrefAttr && repeatedSpaces[0] > 1 && char != ' ' && char != '\'' && char != '"' {
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
				if i >= 0 && html[i-4] == 'h' && html[i-3] == 'r' && html[i-2] == 'e' && html[i-1] == 'f' {
					isBufInHrefAttr = true
				}

				isBufInAttr = true
				writeByteToBuf(&buf, &lastChar, char)
				continue
			}

			// remove space at HTML tag end
			if i+1 <= len(html)-1 && char == ' ' && html[i+1] == '>' {
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
