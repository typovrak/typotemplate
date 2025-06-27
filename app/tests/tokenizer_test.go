package tests

import (
	"testing"
	"typotemplate/html"
)

func validateHTMLTokenizer(t *testing.T, raw string, tokenized []html.Token) {
	tokens := html.Tokenizer(raw)

	if len(tokens) != len(tokenized) {
		t.Errorf("expected length %d, got %d", len(tokenized), len(tokens))
	}

	for i := 0; i < len(tokens); i++ {
		if i < len(tokenized) && tokens[i].Type != tokenized[i].Type {
			t.Errorf("expected Tokens[%d].Type %d, got %d", i, tokenized[i].Type, tokens[i].Type)
		}

		if i < len(tokenized) && tokens[i].Value != tokenized[i].Value {
			t.Errorf("expected Tokens[%d].Value %s (length: %d), got %s (length: %d)",
				i, tokenized[i].Value, len(tokenized[i].Value), tokens[i].Value, len(tokens[i].Value))
		}
	}
}

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

		validateHTMLTokenizer(t, raw, tokenized)
	})

	t.Run("tokenize an anchor", func(t *testing.T) {
		raw := "<a> </a>"
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
				Type:  html.TokenText,
				Value: " ",
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

		validateHTMLTokenizer(t, raw, tokenized)
	})

	t.Run("tokenize an anchor", func(t *testing.T) {
		raw := "< a></a>"
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

		validateHTMLTokenizer(t, raw, tokenized)
	})

	t.Run("tokenize an anchor", func(t *testing.T) {
		raw := "<a ></a>"
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

		validateHTMLTokenizer(t, raw, tokenized)
	})

	t.Run("tokenize an anchor", func(t *testing.T) {
		raw := "<a>< /a>"
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

		validateHTMLTokenizer(t, raw, tokenized)
	})

	t.Run("tokenize an anchor", func(t *testing.T) {
		raw := "<a></ a>"
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

		validateHTMLTokenizer(t, raw, tokenized)
	})

	t.Run("tokenize an anchor", func(t *testing.T) {
		raw := "<a></a >"
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

		validateHTMLTokenizer(t, raw, tokenized)
	})

	t.Run("tokenize an anchor", func(t *testing.T) {
		raw := "<a></a >"
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

		validateHTMLTokenizer(t, raw, tokenized)
	})

	t.Run("tokenize an anchor", func(t *testing.T) {
		raw := "< a    >     Test   </   a>"
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
				Type:  html.TokenText,
				Value: "     Test   ",
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

		validateHTMLTokenizer(t, raw, tokenized)
	})
}

//t.Run("title", func(t *testing.T) {
//	t.Errorf("error")
//})
