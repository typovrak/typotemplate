package tests

import (
	"testing"
	"typotemplate/html"
)

func TestHTMLTokenizer(t *testing.T) {
	t.Run("tokenize an anchor", func(t *testing.T) {
		raw := "<a></a>"
		tokenized := []html.Token{
			{
				Type:  html.TokenTagOpen,
				Value: "v",
			},
		}

		// t.Errorf("error")
	})
}

//t.Run("title", func(t *testing.T) {
//	t.Errorf("error")
//})
