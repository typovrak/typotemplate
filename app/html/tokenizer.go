package html

import (
	"bytes"
)

func Minifier(html string) string {
	var buffer bytes.Buffer
	buffer.Grow(len(html))

	var (
		char     byte
		lastChar byte

		isBufferInTag bool
		bufBytes      []byte
		bufLen        int
	)

	// isBufferInComment := true

	// TODO: gérer les commentaires
	// TODO: gérer le doctype

	// WARN: previous characters = use buffer
	// WARN: next characters = use html
	for i := 0; i < len(html); i++ {
		char = html[i]
		bufBytes = buffer.Bytes()
		bufLen = len(bufBytes)

		if char == '<' {
			isBufferInTag = true
		}

		if isBufferInTag {
			// remove double space in tag
			if (lastChar == ' ' && char == ' ') || (i+1 <= len(html)-1 && char == ' ' && html[i+1] == ' ') {
				continue
			}

			// remove space at HTML tag start
			// TODO: gérer </
			if (lastChar == '<' && char == ' ') ||
				(2 < bufLen && bufBytes[bufLen-2] == '<' && lastChar == '/' && char == ' ') {
				continue
			}

			// remove space at HTML tag end
			if i+1 <= len(html)-1 && char == ' ' && html[i+1] == '>' {
				continue
			}

		} else {
		}

		if char == '>' {
			isBufferInTag = false
		}

		buffer.WriteByte(char)
		lastChar = char
	}

	return buffer.String()
}
