package tests

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"testing"
)

func addLineFeedBetweenErrorThrown(w *os.File, errorBefore *bool, isError bool) {
	if (!isError && *errorBefore) || (isError && !*errorBefore) {
		w.Write([]byte("\n"))
	}

	*errorBefore = isError
}

func handleLineType(
	line *[]byte,
	lineType LineType,
	defaultTitleType []byte,
	color *[]byte,
	w *os.File,
	errorBefore *bool,
	isError bool,
) {
	*color = ColorANSI(lineType.Colors)
	addLineFeedBetweenErrorThrown(w, errorBefore, isError)
	*line = bytes.Replace(*line, defaultTitleType, []byte(lineType.Title), 1)
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
	Run         LineType
	Fail        LineType
	Pass        LineType
	Skip        LineType
	Failed      LineType
	Ok          LineType
	ErrorThrown LineType
}

type LineType struct {
	Colors ANSIConfig
	Title  string
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
		Run: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundCyan],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "=== RUN  ",
		},
		Fail: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundRed],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "--- FAIL:",
		},
		Pass: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundGreen],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "--- PASS:",
		},
		Skip: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundYellow],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "--- SKIP:",
		},
		Failed: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundBlack],
				Background: ColorANSIBackground[ANSIBackgroundRed],
			},
			Title: "FAIL",
		},
		Ok: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleBold],
				Foreground: ColorANSIForeground[ANSIForegroundBlack],
				Background: ColorANSIBackground[ANSIBackgroundGreen],
			},
			Title: "PASS",
		},

		ErrorThrown: LineType{
			Colors: ANSIConfig{
				Style:      ColorANSISTyle[ANSIStyleNormal],
				Foreground: ColorANSIForeground[ANSIForegroundWhite],
				Background: ColorANSIBackground[ANSIBackgroundNone],
			},
			Title: "",
		},
	}
}

func RunColorizeTests(m *testing.M, opts Opts) int {
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

	defaultTitle := struct {
		Run         []byte
		Fail        []byte
		Pass        []byte
		Skip        []byte
		Failed      []byte
		Ok          []byte
		ErrorThrown []byte
	}{
		Run:         []byte("=== RUN  "),
		Fail:        []byte("--- FAIL:"),
		Pass:        []byte("--- PASS:"),
		Skip:        []byte("--- SKIP:"),
		Failed:      []byte("FAIL"),
		Ok:          []byte("PASS"),
		ErrorThrown: []byte(""),
	}

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
			if bytes.Contains(line, defaultTitle.Run) {
				handleLineType(&line, opts.Run, defaultTitle.Run, &color, stdout, &errorBefore, false)
				tabs = true

				// --- FAIL:
			} else if bytes.Contains(line, defaultTitle.Fail) {
				handleLineType(&line, opts.Fail, defaultTitle.Fail, &color, stdout, &errorBefore, false)
				tabs = true

				// --- PASS:
			} else if bytes.Contains(line, defaultTitle.Pass) {
				handleLineType(&line, opts.Pass, defaultTitle.Pass, &color, stdout, &errorBefore, false)
				tabs = true

				// --- SKIP:
			} else if bytes.Contains(line, defaultTitle.Skip) {
				handleLineType(&line, opts.Skip, defaultTitle.Skip, &color, stdout, &errorBefore, false)
				tabs = true

				// FAIL
			} else if bytes.Equal(line, defaultTitle.Failed) {
				handleLineType(&line, opts.Failed, defaultTitle.Failed, &color, stdout, &errorBefore, false)
				stdout.Write([]byte("\n"))

				// ok
			} else if bytes.Equal(line, defaultTitle.Ok) {
				handleLineType(&line, opts.Ok, defaultTitle.Ok, &color, stdout, &errorBefore, false)
				stdout.Write([]byte("\n"))

				// error thrown
			} else {
				handleLineType(&line, opts.ErrorThrown, defaultTitle.ErrorThrown, &color, stdout, &errorBefore, true)
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

	// restore outputs
	os.Stdout = stdout
	os.Stderr = stderr

	return exitCode
}

// WARN: all tests must be in this folder, no subfolder authorized
func TestMain(m *testing.M) {
	os.Setenv("APP_GO_TEST", "true")

	opts := NewDefaultOpts()
	exitCode := RunColorizeTests(m, opts)

	os.Exit(exitCode)
}
