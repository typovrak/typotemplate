package tests

import (
	"os"
	"testing"

	"github.com/typovrak/typotestcolor"
)

var Opts = typotestcolor.NewDefaultOpts()

// WARN: all tests must be in this folder, no subfolder authorized
func TestMain(m *testing.M) {
	os.Setenv("APP_GO_TEST", "true")

	Opts.Run.Section.Header.Hide = true
	Opts.Run.Section.Footer.Hide = true

	Opts.Fail.Section.Header.Hide = true
	Opts.Fail.Section.Footer.Hide = true

	Opts.Pass.Section.Header.Hide = true
	Opts.Pass.Section.Footer.Hide = true

	Opts.Skip.Section.Header.Hide = true
	Opts.Skip.Section.Footer.Hide = true

	exitCode := typotestcolor.RunTestColor(m, Opts)

	os.Exit(exitCode)
}
