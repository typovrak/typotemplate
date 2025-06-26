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

func IsCharAlphanumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}

func addToken(tokens *[]Token, buffer *bytes.Buffer, tokenType TokenType, tokenValue string) {
	*tokens = append(*tokens, Token{Type: tokenType, Value: tokenValue})
	buffer.Reset()
}

func Tokenizer(html string) []Token {
	var tokens []Token
	var buffer bytes.Buffer

	isBufferInTag := false
	isBufferAlphanumeric := false

	for i := 0; i < len(html); i++ {
		if len(buffer.String()) == 0 {
			isBufferAlphanumeric = true
		}

		char := html[i]

		// *buffer == < -> TokenTagOpenChar
		// *buffer == > -> TokenTagCloseChar

		// *buffer == _  && isBufferInTag == false -> continue

		// buffer == string && isBufferInTag && inAttribute == false -> TokenAttributeKey
		// buffer == = && isBufferInTag == true -> TokenAttributeOperator
		// buffer == "' && isBufferInTag == true && pas de \ avant -> TokenAttributeOpen TokenAttributeClose

		// *else -> TokenText

		if char == ' ' && !isBufferInTag {
			// no need to manage alphanumeric because not in tag
			buffer.WriteByte(char)
			continue
		}

		if char == '<' {
			if len(buffer.String()) > 0 {
				if isBufferAlphanumeric && !isBufferInTag {
					addToken(&tokens, &buffer, TokenTagName, buffer.String())
				}
			}

			addToken(&tokens, &buffer, TokenTagOpenChar, string(char))
			isBufferInTag = true
			continue
		}

		if char == '>' {
			if len(buffer.String()) > 0 {
				if isBufferAlphanumeric && isBufferInTag {
					addToken(&tokens, &buffer, TokenTagName, buffer.String())
				}
			}

			addToken(&tokens, &buffer, TokenTagCloseChar, string(char))
			isBufferInTag = false
			continue
		}

		if char == '/' && len(buffer.String()) == 0 &&
			len(tokens)-1 >= 0 && tokens[len(tokens)-1].Type == TokenTagOpenChar && tokens[len(tokens)-1].Value == "<" {
			// TODO: need to do something with the previous non empty buffer

			addToken(&tokens, &buffer, TokenTagOpenChar, string(char))
			continue
		}

		if char != ' ' && isBufferInTag {
			if isBufferInTag && isBufferAlphanumeric {
				isBufferAlphanumeric = IsCharAlphanumeric(char)
			}

			buffer.WriteByte(char)
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
