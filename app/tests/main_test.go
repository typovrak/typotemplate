package tests

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

// TODO: add emoji?

func addLineFeedBetweenErrorThrown(w *os.File, runBefore *bool, errorBefore *bool, isRun bool, isError bool) {
	if *runBefore || *errorBefore {
		w.Write([]byte("\n"))
	}

	*runBefore = isRun
	*errorBefore = isError
}

// WARN: all tests must be in this folder, no subfolder authorized
func TestMain(m *testing.M) {
	os.Setenv("APP_GO_TEST", "true")

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

	// Close the writer end of the pipe so the reader stops at EOF
	w.Close()

	// Setup the reader
	reader := bufio.NewReader(r)

	runMatch := []byte("=== RUN")
	failMatch := []byte("--- FAIL:")
	passMatch := []byte("--- PASS:")
	skipMatch := []byte("--- SKIP")
	passedMatch := []byte("PASS")
	failedMatch := []byte("FAIL")

	colorBlue := []byte("\033[0;36m")
	colorRed := []byte("\033[0;31m")
	colorGreen := []byte("\033[0;32m")
	colorPurple := []byte("\033[0;35m")
	colorReset := []byte("\033[0m")

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
				color = colorBlue
				tabs = true
				addLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, true, false)

				// --- FAIL:
			} else if bytes.Contains(line, failMatch) {
				color = colorRed
				tabs = true
				addLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, false)

				// --- PASS:
			} else if bytes.Contains(line, passMatch) {
				color = colorGreen
				tabs = true
				addLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, false)

				// --- SKIP:
			} else if bytes.Contains(line, skipMatch) {
				color = colorPurple
				tabs = true
				addLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, false)

				// FAIL
			} else if bytes.Equal(line, failedMatch) {
				color = colorRed
				addLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, false)

				// ok
			} else if bytes.Equal(line, passedMatch) {
				color = colorGreen
				addLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, false)

				// error thrown
			} else {
				addLineFeedBetweenErrorThrown(stdout, &runBefore, &errorBefore, false, true)
			}

			if color != nil {
				stdout.Write(color)
			}

			if tabs {
				stdout.Write([]byte("\t"))
			}

			stdout.Write(line)

			if color != nil {
				stdout.Write(colorReset)
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
