package tests

import (
	"testing"
	"typotemplate/html"
)

func TestIsCharAlphanumeric(t *testing.T) {
	t.Run("isCharAlphanumeric", func(t *testing.T) {
		res := html.IsCharAlphanumeric('a')

		if !res {
			t.Error("expected true, got false")
		}
	})

	t.Run("isCharAlphanumeric", func(t *testing.T) {
		res := html.IsCharAlphanumeric('1')

		if !res {
			t.Error("expected true, got false")
		}
	})

	t.Run("isCharAlphanumeric", func(t *testing.T) {
		res := html.IsCharAlphanumeric('<')

		if res {
			t.Error("expected false, got true")
		}
	})

	t.Run("isCharAlphanumeric", func(t *testing.T) {
		res := html.IsCharAlphanumeric(' ')

		if res {
			t.Error("expected false, got true")
		}
	})
}
