package html

import (
	"bytes"
)

// WARN: je pars du principe que un nom d'élément HTML est un alpha numérique.

type Token struct {
	Type  TokenType
	Value string
}

type TokenType int

const (
	TokenTagOpenChar TokenType = iota
	TokenTagName
	TokenTagCloseChar

	TokenAttributeKey
	TokenAttributeOperatorChar

	TokenAttributeOpenChar
	TokenAttributeValue
	TokenAttributeCloseChar

	TokenText
)

func isCharAlphanumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}

func Tokenizer(html string) []Token {
	var tokens []Token
	var buffer bytes.Buffer

	isBufferInTag := false
	isBufferAlphanumeric := false

	for i := 0; i < len(html); i++ {
		char := html[i]

		// *buffer == < -> TokenTagOpenChar
		// *buffer == > -> TokenTagCloseChar

		// *buffer == _  && isBufferInTag == false -> continue

		// buffer == string && isBufferInTag && inAttribute == false -> TokenAttributeKey
		// buffer == = && isBufferInTag == true -> TokenAttributeOperator
		// buffer == "' && isBufferInTag == true && pas de \ avant -> TokenAttributeOpen TokenAttributeClose

		// *else -> TokenText

		if char == ' ' && !isBufferInTag {
			buffer.WriteByte(char)
			continue
		}

		if char == '<' {
			tokens = append(tokens, Token{Type: TokenTagOpenChar, Value: string(char)})
			buffer.Reset()
			isBufferInTag = true
			continue
		}

		if char == '>' {
			tokens = append(tokens, Token{Type: TokenTagCloseChar, Value: string(char)})
			buffer.Reset()
			isBufferInTag = false
			continue
		}

		if char == '/' && len(buffer.String()) == 0 &&
			len(tokens)-1 >= 0 && tokens[len(tokens)-1].Type == TokenTagOpenChar && tokens[len(tokens)-1].Value == "<" {

			tokens = append(tokens, Token{Type: TokenTagOpenChar, Value: string(char)})
			buffer.Reset()
			continue
		}

		if char != ' ' && isBufferInTag {
			buffer.WriteByte(char)

			if isBufferAlphanumeric {
				isBufferAlphanumeric = isCharAlphanumeric(char)
			}

			continue
		}

		//if char == ' ' && len(buffer.String()) > 0 {
		//	tokens = append(tokens, Token{Type: TokenTagName, Value: buffer.String()})
		//	buffer.Reset()
		//	continue
		//}
	}

	return tokens
}
