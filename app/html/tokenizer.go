package html

import (
	"bytes"
)

type Token struct {
	Type  TokenType
	Value string
}

type TokenType int

const (
	TokenTagOpen TokenType = iota
	TokenTagClose

	TokenAttributeKey
	TokenAttributeOperator

	TokenAttributeOpen
	TokenAttributeValue
	TokenAttributeClose

	TokenText
)

func Tokenizer(html string) []Token {
	var tokens []Token
	var buffer bytes.Buffer
	inTag := false

	for i := 0; i < len(html); i++ {
		char := html[i]

		// *buffer == < -> TokenTagOpen
		// *buffer == > -> TokenTagClose

		// *buffer == _  && inTag == false -> continue

		// buffer == string && inTag && inAttribute == false -> TokenAttributeKey
		// buffer == = && inTag == true -> TokenAttributeOperator
		// buffer == "' && inTag == true && pas de \ avant -> TokenAttributeOpen TokenAttributeClose

		// *else -> TokenText

		if char == ' ' && inTag == false {
			buffer.WriteByte(char)
			continue
		}

		if char == '<' {
			tokens = append(tokens, Token{Type: TokenTagOpen, Value: string(char)})
			buffer.Reset()
			inTag = true
			continue
		}

		if char == '>' {
			tokens = append(tokens, Token{Type: TokenTagClose, Value: string(char)})
			buffer.Reset()
			inTag = false
			continue
		}

		buffer.WriteByte(char)
	}

	return tokens
}
