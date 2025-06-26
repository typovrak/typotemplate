package main

import (
	"bytes"
	"fmt"
	"strings"
)

// TODO: je veux que le résultat de < useoathueoastn > -> </>
// < uasthnueo />
// <img /> <img >
// node-html-parser
// regarder le parser html chrome

// WARN: no regex challenge and no depedencies.

// étapes
// 1. créer un composant
// 		* qui est une fonction golang
// 		* qui prend des props
// 		* qui prend des childs composants
// 		* return html pure static à 100%
// 		-
// 2. créer un objet de configuration pour page
// 		* définir les composants à importer
// 		* définir des propos variables pour les composants importés
//		-
// 3. créer une fonction qui va générer la page static
// 		* ajouter le contenu variable des children composant
// 		- minifier le html en une ligne et supprimer les whitespaces inutiles
//		-
// 4. mise en cache de la géneration dans un .tmp.html ???
// 		- en chunck
// 		-

type HeaderProps struct {
	backgroundColor  string
	titleBrandSuffix string
}

func Header(props HeaderProps) string {
	return "<header  >" +
		TitleBrand(TitleBrandProps{
			Suffix: props.titleBrandSuffix,
		}) +
		"</  header>"
}

type TitleBrandProps struct {
	Suffix string
}

func TitleBrand(props TitleBrandProps) string {
	return "<a   href=\"/\" class=\"\" >   " +
		"typovrakéé`eèè˙çç  ^💜" +
		props.Suffix +
		"  </a>" +
		"<div class=\"test \"></div>" +
		"<div    data-content=\"   \" title=\" test \"></div>" +
		"<p class=\"coucou bonjour  test aurevoir\"></p>"
}

func PageIndex() string {
	return Header(HeaderProps{
		backgroundColor:  "purple",
		titleBrandSuffix: "!",
	})
}

func MinifyHTML(html string) string {
	fmt.Println(html)
	tokens := Tokenize(html)

	fmt.Println(tokens)
	minified := minify(tokens)

	return minified
}

type Token struct {
	Type  TokenType
	Value string
}

type TokenType int

const (
	TokenTagOpen TokenType = iota
	TokenTagClose

	TokenAttributeKey
	TokenAttributeValue

	TokenText
)

// TODO: tester les caractères qui sont sur plusieurs bytes.
func Tokenize(html string) []Token {
	var tokens []Token
	var buffer bytes.Buffer

	for i := 0; i < len(html); i++ {
		char := html[i]

		switch char {
		case '<':
			if buffer.Len() > 0 {
				tokens = append(tokens, Token{Type: TokenText, Value: buffer.String()})
				buffer.Reset()
			}

		case '>':
			tokens = parseTag(tokens, buffer.String())
			buffer.Reset()

		default:
			buffer.WriteByte(char)
		}
	}

	return tokens
}

func parseTag(tokens []Token, buffer string) []Token {
	if strings.HasPrefix(buffer, "/") {
		tokens = append(tokens, Token{Type: TokenTagClose, Value: buffer[1:]})
		return tokens
	}

	// WARN: string.Fields() this will split attributes and remove whitespace from TokenText also so no terminal cli website...
	fmt.Println("PARTS:", buffer)

	var bufferParts []string

	// manage \" as non delimiter
	delimiter := ""
	currentString := ""

	// DOC: RFC attributes delimiter can only be " or ' (or > for ending)
	// [Page 14] https://www.ietf.org/rfc/rfc1866.txt
	for i := 0; i < len(buffer); i++ {
		char := string(buffer[i])

		// whitespace between tags/attributes
		if char == " " && delimiter == "" {
			if len(currentString) > 0 {
				bufferParts = append(bufferParts, currentString)
				fmt.Println("currentString:", currentString)
				currentString = ""
			}

			continue
		}

		// start delimiter
		// based on RFC HTML attributes delimiters
		if char == "\"" || char == "'" {
			delimiter = char
			currentString += "\""
			continue
		}

		// end delimiter
		// based on RFC HTML attributes delimiters
		if char == "\"" || char == "'" || char == ">" {
			delimiter = ""
			currentString += "\""
			continue
		}

		currentString += string(buffer[i])
	}

	tokens = append(tokens, Token{Type: TokenTagOpen, Value: bufferParts[0]})

	// TEST: key without value
	for i := 1; i < len(bufferParts); i++ {
		tokens = append(tokens, Token{Type: TokenAttribute, Value: bufferParts[i]})
	}

	fmt.Println(tokens)
	fmt.Println("parts", bufferParts)
	fmt.Println()

	return tokens
}

func minify(tokens []Token) string {
	var builder strings.Builder
	// lastAttribute := ""

	// WARN: do everything to manage the spacing for the next element.
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Type {
		// TODO: add intag
		case TokenTagOpen:
			s := "<" + tokens[i].Value

			if tokens[i+1].Type != TokenAttribute {
				s += ">"
			} else {
				s += " "
			}

			builder.WriteString(s)

		case TokenTagClose:
			builder.WriteString("</" + tokens[i].Value + ">")

		case TokenAttribute:
			s := tokens[i].Value

			if tokens[i+1].Type != TokenAttribute {
				s += ">"
			} else {
				s += " "
			}

			// store attribute name in lastAttribute
			// split := strings.Split(tokens[i].Value, "=")

			// if len(split) > 0 {
			//	lastAttribute = split[0]

			//	content := split[:1]
			//}

			//// check if the other one is an attribute
			//// remove all whitespaces
			//// close the attribute

			//if len(tokens[i].Value) >= 5 && tokens[i].Value[:5] == "class" {
			//	// remove trailing whitespace in class
			//	tokens[i].Value = "class=\"" + strings.Trim(tokens[i].Value, " \"'") + "\""
			//}

			builder.WriteString(s)

		case TokenText:
			builder.WriteString(strings.TrimSpace(tokens[i].Value))
		}
	}
	return builder.String()
}

func main() {
	fmt.Println(MinifyHTML(PageIndex()))
}
