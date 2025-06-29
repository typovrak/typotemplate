package tests

import (
	"testing"
	"typotemplate/html"
)

func validateHTMLMinifier(t *testing.T, raw string, expected string) {
	res := html.Minifier(raw)

	if res != expected {
		t.Errorf("expected %s, got %s", expected, res)
	}
}

func TestHTMLMinifier(t *testing.T) {
	t.Run("minifier_0", func(t *testing.T) {
		raw := "<a></a>"
		expected := "<a></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_1", func(t *testing.T) {
		raw := "< a></a>"
		expected := "<a></a>"

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_2", func(t *testing.T) {
		raw := "< a  ></a>"
		expected := "<a></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_3", func(t *testing.T) {
		raw := "< a>   </a>"
		expected := "<a>   </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_4", func(t *testing.T) {
		raw := "< a>Test a</a>"
		expected := "<a>Test a</a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_5", func(t *testing.T) {
		raw := "< a></  a>"
		expected := "<a></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_6", func(t *testing.T) {
		raw := "< a></a  >"
		expected := "<a></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_7", func(t *testing.T) {
		raw := "< a>  Test  </a>"
		expected := "<a>  Test  </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_8", func(t *testing.T) {
		raw := "<a  href=\" https://mscholz.dev\"></a>"
		expected := "<a href=\"https://mscholz.dev\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_9", func(t *testing.T) {
		raw := "<a  href=\" https://mscholz.dev/blog test coucou \"></a>"
		expected := "<a href=\"https://mscholz.dev/blog test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_10", func(t *testing.T) {
		raw := "<a  title=\"test\"></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_11", func(t *testing.T) {
		raw := "<a  title='test'></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_12", func(t *testing.T) {
		raw := "<a  title='  test '></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_13", func(t *testing.T) {
		raw := "<a  title=test></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_14", func(t *testing.T) {
		raw := "<a  title='test \" '></a>"
		expected := "<a title=\"test &quot;\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_15", func(t *testing.T) {
		raw := "<a  title=test   ></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_16", func(t *testing.T) {
		raw := "<a  title=test ></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_17", func(t *testing.T) {
		raw := "<a  title= test ></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})
}

//t.Run("title", func(t *testing.T) {
//	t.Errorf("error")
//})
