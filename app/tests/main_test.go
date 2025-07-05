package tests

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"testing"
)

// TODO: add emoji?

func AddLineFeedBetweenErrorThrown(w *os.File, errorBefore *bool, isError bool) {
	if (!isError && *errorBefore) || (isError && !*errorBefore) {
		w.Write([]byte("\n"))
	}

	*errorBefore = isError
}

// ansi
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
	ANSIForegroundNone ANSIForeground = iota
	ANSIForegroundBlack
	ANSIForegroundRed
	ANSIForegroundGreen
	ANSIForegroundYellow
	ANSIForegroundBlue
	ANSIForegroundPurple
	ANSIForegroundCyan
	ANSIForegroundWhite
)

var ColorANSIForeground = map[ANSIForeground]int{
	ANSIForegroundNone:   0,
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
	ANSIBackgroundNone ANSIBackground = iota
	ANSIBackgroundBlack
	ANSIBackgroundRed
	ANSIBackgroundGreen
	ANSIBackgroundYellow
	ANSIBackgroundBlue
	ANSIBackgroundPurple
	ANSIBackgroundCyan
	ANSIBackgroundWhite
)

var ColorANSIBackground = map[ANSIBackground]int{
	ANSIBackgroundNone:   0,
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

// global
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
	color := []byte("\033[")

	color = append(color, []byte(strconv.Itoa(config.Style))...)
	color = append(color, ';')
	color = append(color, []byte(strconv.Itoa(config.Foreground))...)

	if config.Background != ColorANSIBackground[ANSIBackgroundNone] {
		color = append(color, ';')
		color = append(color, []byte(strconv.Itoa(config.Background))...)
	}

	color = append(color, 'm')

	return color
}

func NewDefaultOpts() Opts {
	return Opts{
		Color: ColorOpts{
			Run: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundCyan],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Fail: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundRed],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Pass: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundGreen],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Skip: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundYellow],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Failed: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundBlack],
				Background: ColorANSIBackground[ANSIBackgroundRed],
			},
			Ok: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundBlack],
				Background: ColorANSIBackground[ANSIBackgroundGreen],
			},
			ErrorThrown: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundWhite],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
		},
	}
}

func ColorizeTests(m *testing.M, opts Opts) {
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

	errorBefore := false

	// read line by line
	for {
		line, err := reader.ReadBytes('\n')

		if len(line) > 0 {
			line = bytes.TrimRight(line, "\n")
			line = bytes.TrimLeft(line, " ")

			var color []byte
			tabs := false

			// manage color and style line depending on the content
			// === RUN
			if bytes.Contains(line, []byte("=== RUN")) {
				color = ColorANSI(opts.Color.Run)
				tabs = true
				AddLineFeedBetweenErrorThrown(stdout, &errorBefore, false)

				// --- FAIL:
			} else if bytes.Contains(line, []byte("--- FAIL:")) {
				color = ColorANSI(opts.Color.Fail)
				tabs = true
				AddLineFeedBetweenErrorThrown(stdout, &errorBefore, false)

				// --- PASS:
			} else if bytes.Contains(line, []byte("--- PASS:")) {
				color = ColorANSI(opts.Color.Pass)
				tabs = true
				AddLineFeedBetweenErrorThrown(stdout, &errorBefore, false)

				// --- SKIP:
			} else if bytes.Contains(line, []byte("--- SKIP:")) {
				color = ColorANSI(opts.Color.Skip)
				tabs = true
				AddLineFeedBetweenErrorThrown(stdout, &errorBefore, false)

				// FAIL
			} else if bytes.Equal(line, []byte("FAIL")) {
				color = ColorANSI(opts.Color.Failed)
				AddLineFeedBetweenErrorThrown(stdout, &errorBefore, false)
				stdout.Write([]byte("\n"))

				// ok
			} else if bytes.Equal(line, []byte("PASS")) {
				color = ColorANSI(opts.Color.Ok)
				AddLineFeedBetweenErrorThrown(stdout, &errorBefore, false)
				stdout.Write([]byte("\n"))

				// error thrown
			} else {
				color = ColorANSI(opts.Color.ErrorThrown)
				AddLineFeedBetweenErrorThrown(stdout, &errorBefore, true)
			}

			if tabs {
				stdout.Write([]byte("\t"))
			}

			stdout.Write(color)
			stdout.Write(line)
			stdout.Write([]byte("\033[0m"))
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

	opts := NewDefaultOpts()
	ColorizeTests(m, opts)
}
