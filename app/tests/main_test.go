package tests

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

// TODO: add emoji?

func AddLineFeedBetweenErrorThrown(w *os.File, runBefore *bool, errorBefore *bool, isRun bool, isError bool) {
	if (isRun && *errorBefore) || (isError && *runBefore) {
		w.Write([]byte("\n"))
	}

	*runBefore = isRun
	*errorBefore = isError
}

type Color int

const (
	ColorNone Color = iota
	ColorReset
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorPurple
	ColorCyan
	ColorWhite

	ColorBgRed
)

type ANSIStyle int

const (
	ANSIStyleReset ANSIStyle = iota
	ANSIStyleBold
	ANSIStyleDim
	ANSIStyleUnderline
	ANSIStyleInverse
	ANSIStyleHidden
	ANSIStyleNormal
)

var ColorANSISTyle = map[ANSIStyle]int{
	ANSIStyleReset:     0,
	ANSIStyleBold:      1,
	ANSIStyleDim:       2,
	ANSIStyleUnderline: 4,
	ANSIStyleInverse:   7,
	ANSIStyleHidden:    8,
	ANSIStyleNormal:    22,
}

type ANSIForeground int

const (
	ANSIForegroundBlack ANSIForeground = iota
	ANSIForegroundRed
	ANSIForegroundGreen
	ANSIForegroundYellow
	ANSIForegroundBlue
	ANSIForegroundPurple
	ANSIForegroundCyan
	ANSIForegroundWhite
)

var ColorANSIForeground = map[ANSIForeground]int{
	ANSIForegroundBlack:  30,
	ANSIForegroundRed:    31,
	ANSIForegroundGreen:  32,
	ANSIForegroundYellow: 33,
	ANSIForegroundBlue:   34,
	ANSIForegroundPurple: 35,
	ANSIForegroundCyan:   36,
	ANSIForegroundWhite:  37,
}

type ANSIBackground int

const (
	ANSIBackgroundBlack ANSIBackground = iota
	ANSIBackgroundRed
	ANSIBackgroundGreen
	ANSIBackgroundYellow
	ANSIBackgroundBlue
	ANSIBackgroundPurple
	ANSIBackgroundCyan
	ANSIBackgroundWhite
)

var ColorANSIBackground = map[ANSIBackground]int{
	ANSIBackgroundBlack:  40,
	ANSIBackgroundRed:    41,
	ANSIBackgroundGreen:  42,
	ANSIBackgroundYellow: 43,
	ANSIBackgroundBlue:   44,
	ANSIBackgroundPurple: 45,
	ANSIBackgroundCyan:   46,
	ANSIBackgroundWhite:  47,
}

type ANSIConfig struct {
	Style      int
	Foreground int
	Background int
}

type Opts struct {
	Color ColorOpts
}

type ColorOpts struct {
	Run         ANSIConfig
	Fail        ANSIConfig
	Pass        ANSIConfig
	Skip        ANSIConfig
	Failed      ANSIConfig
	Ok          ANSIConfig
	ErrorThrown ANSIConfig
}

func ColorANSI(config ANSIConfig) []byte {
	return
}

var ColorANSI = map[Color][]byte{
	ColorNone:   []byte(""),
	ColorReset:  []byte("\033[0m"),
	ColorBlack:  []byte("\033[0;30m"),
	ColorRed:    []byte("\033[0;31m"),
	ColorGreen:  []byte("\033[0;32m"),
	ColorYellow: []byte("\033[0;33m"),
	ColorBlue:   []byte("\033[0;34m"),
	ColorPurple: []byte("\033[0;35m"),
	ColorCyan:   []byte("\033[0;36m"),
	ColorWhite:  []byte("\033[0;37m"),

	ColorBgRed: []byte("\033[41m"),
}

func ColorizeTests(m *testing.M, opts Opts) {
	// default values
	if opts.Color.Run == 0 {
		opts.Color.Run = ColorCyan
	}

	if opts.Color.Fail == 0 {
		opts.Color.Fail = ColorRed
	}

	if opts.Color.Pass == 0 {
		opts.Color.Pass = ColorGreen
	}

	if opts.Color.Skip == 0 {
		opts.Color.Skip = ColorYellow
	}

	if opts.Color.Failed == 0 {
		opts.Color.Failed = ColorRed
	}

	if opts.Color.Ok == 0 {
		opts.Color.Ok = ColorGreen
	}

	if opts.Color.ErrorThrown == 0 {
		opts.Color.ErrorThrown = ColorWhite
	}

	// create a pipe
	r, w, _ := os.Pipe()

	// backup original outputs
	stdout := os.Stdout
	stderr := os.Stderr

	// redirect stdout and stderr to the pipe
	os.Stdout = w
	os.Stderr = w

	// Run tests
	exitCode := m.Run()

	// close the writer end of the pipe so the reader stops at EOF
	w.Close()

	// setup the reader
	reader := bufio.NewReader(r)

	runMatch := []byte("=== RUN")
	failMatch := []byte("--- FAIL:")
	passMatch := []byte("--- PASS:")
	skipMatch := []byte("--- SKIP")
	passedMatch := []byte("PASS")
	failedMatch := []byte("FAIL")

	runBefore := false
	errorBefore := false

	// read line by line
	for {
		line, err := reader.ReadBytes('\n')

		if len(line) > 0 {
			line = bytes.TrimRight(line, "\n")
			line = bytes.TrimLeft(line, " ")

			var color []byte
			tabs := false

			// manage styling depending on bytes match
			// === RUN
			if bytes.Contains(line, runMatch) {
				color = ColorANSI[opts.Color.Run]
				tabs = true
				AddLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, true, false)

				// --- FAIL:
			} else if bytes.Contains(line, failMatch) {
				color = ColorANSI[opts.Color.Fail]
				tabs = true
				AddLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, false)

				// --- PASS:
			} else if bytes.Contains(line, passMatch) {
				color = ColorANSI[opts.Color.Pass]
				tabs = true
				AddLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, false)

				// --- SKIP:
			} else if bytes.Contains(line, skipMatch) {
				color = ColorANSI[opts.Color.Skip]
				tabs = true
				AddLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, false)

				// FAIL
			} else if bytes.Equal(line, failedMatch) {
				color = ColorANSI[opts.Color.Failed]
				AddLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, false)
				stdout.Write([]byte("\n"))

				// ok
			} else if bytes.Equal(line, passedMatch) {
				color = ColorANSI[opts.Color.Ok]
				AddLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, false)
				stdout.Write([]byte("\n"))

				// error thrown
			} else {
				AddLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, true)
			}

			if color != nil {
				stdout.Write(color)
			}

			if tabs {
				stdout.Write([]byte("\t"))
			}

			stdout.Write(line)

			if color != nil {
				stdout.Write(ColorANSI[ColorReset])
			}

			stdout.Write([]byte("\n"))
		}

		if err != nil {
			break
		}
	}

	// Restore outputs
	os.Stdout = stdout
	os.Stderr = stderr

	os.Exit(exitCode)
}

// WARN: all tests must be in this folder, no subfolder authorized
func TestMain(m *testing.M) {
	os.Setenv("APP_GO_TEST", "true")

	ColorizeTests(m, Opts{
		ColorOpts{
			Failed: ColorBgRed,
		},
	})
}
