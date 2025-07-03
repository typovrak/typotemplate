package tests

import (
	"os"
	"testing"
)

//func colorizeTestResult() string {
//}

// WARN: all tests must be in this folder, no subfolder authorized
func TestMain(m *testing.M) {
	// before tests
	os.Setenv("APP_GO_TEST", "true")

	// run tests
	exitVal := m.Run()

	// after tests

	// add colors

	// exit value from tests
	os.Exit(exitVal)
}
