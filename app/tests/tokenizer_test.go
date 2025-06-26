package tests

import (
	"fmt"
	"testing"
	"typotemplate/html"
)

func TestHTMLTokenizer(t *testing.T) {
	t.Run("tokenize an anchor", func(t *testing.T) {
		raw := "<a></a>"
		tokenized := []html.Token{
			{
				Type:  html.TokenTagOpenChar,
				Value: "<",
			},
			{
				Type:  html.TokenTagName,
				Value: "a",
			},
			{
				Type:  html.TokenTagCloseChar,
				Value: ">",
			},
			{
				Type:  html.TokenTagOpenChar,
				Value: "<",
			},
			{
				Type:  html.TokenTagOpenChar,
				Value: "/",
			},
			{
				Type:  html.TokenTagName,
				Value: "a",
			},
			{
				Type:  html.TokenTagCloseChar,
				Value: ">",
			},
		}

		tokens := html.Tokenizer(raw)

		fmt.Println(tokenized)
		fmt.Println(tokens)

		if len(tokens) != len(tokenized) {
			t.Errorf("expected length %d, got %d", len(tokenized), len(tokens))
		}

		for i := 0; i < len(tokens); i++ {
			if i < len(tokenized) && tokens[i].Type != tokenized[i].Type {
				t.Errorf("expected Tokens[%d].Type %d, got %d", i, tokenized[i].Type, tokens[i].Type)
			}

			if i < len(tokenized) && tokens[i].Value != tokenized[i].Value {
				t.Errorf("expected Tokens[%d].Value %s, got %s", i, tokenized[i].Value, tokens[i].Value)
			}
		}
	})
}

//t.Run("title", func(t *testing.T) {
//	t.Errorf("error")
//})
