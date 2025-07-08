package tests

import (
	"os"
	"testing"

	"github.com/typovrak/typotestcolor"
)

// WARN: all tests must be in this folder, no subfolder authorized
func TestMain(m *testing.M) {
	os.Setenv("APP_GO_TEST", "true")

	exitCode := typotestcolor.Default(m)

	os.Exit(exitCode)
}
