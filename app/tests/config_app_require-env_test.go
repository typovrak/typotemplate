package tests

import (
	"os"
	"testing"
	"typotemplate/config/app"
)

func TestRequireEnv(t *testing.T) {
	t.Run("APP_GO_TEST env var not defined", func(t *testing.T) {
		os.Unsetenv("APP_GO_TEST")

		err := app.RequireEnv()
		if err == nil {
			t.Error("expected an error, got nil")
		}

		os.Setenv("APP_GO_TEST", "true")
	})

	t.Run("all vars must be defined", func(t *testing.T) {
		err := app.RequireEnv()
		if err != nil {
			t.Errorf("expected an nil, got %v", err)
		}
	})
}
