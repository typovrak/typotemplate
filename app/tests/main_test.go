package tests

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

// WARN: all tests must be in this folder, no subfolder authorized
//func TestMain(m *testing.M) {
//	// before tests
//	os.Setenv("APP_GO_TEST", "true")
//
//	// create a pipe between reader and writer
//	r, w, _ := os.Pipe()
//
//	// save os.Stdout
//	stdout := os.Stdout
//
//	// redirect all os.Stdout into os.Pipe()
//	os.Stdout = w
//
//	failMatch := []byte("--- FAIL:")
//	passMatch := []byte("--- PASS:")
//
//	//	colorDefault := []byte("\033[36m")
//	colorReset := []byte("\033[0m")
//
//	//	colorWhite := []byte("\033[0;37m")
//	colorRed := []byte("\033[0;31m")
//	colorGreen := []byte("\033[0;32m")
//
//	// start a goroutine
//	go func() {
//		// 1Ko, standart value for this type of buffer
//		buf := make([]byte, 1024*2)
//
//		// infinite loop
//		for {
//			n, err := r.Read(buf)
//			bufToAdd := buf[:n]
//
//			bufSplited := bytes.Split(bufToAdd, []byte("\n"))
//
//			var colored []byte
//
//			// TODO: j'ai 0 dans la boucle pour la valeur bufSplitedLen, pourquoi?
//			// TODO: manage inconsistency in test runner print
//
//			// TODO: manage === RUN and errors print color and trimLeft value
//
//			bufSplitedLen := len(bufSplited)
//
//			for i := 0; i < bufSplitedLen; i++ {
//				bufSplited[i] = bytes.TrimLeft(bufSplited[i], " ")
//
//				// TODO: add tab before every result
//				bufSplited[i] = append(bufSplited[i], '\t')
//
//				if bytes.Contains(bufSplited[i], failMatch) {
//					bufSplited[i] = append(colorRed, bufSplited[i]...)
//				} else if bytes.Contains(bufSplited[i], passMatch) {
//					bufSplited[i] = append(colorGreen, bufSplited[i]...)
//				}
//
//				bufSplited[i] = append(bufSplited[i], colorReset...)
//				bufSplited[i] = append(bufSplited[i], []byte("bufSplited len:")...)
//				bufSplited[i] = append(bufSplited[i], []byte(strconv.Itoa(bufSplitedLen))...)
//				bufSplited[i] = append(bufSplited[i], '\n')
//
//				colored = append(colored, bufSplited[i]...)
//			}
//
//			// stdout.Write([]byte(strconv.QuoteToASCII(string(colored))))
//			stdout.Write(colored)
//
//			// exit at w.Close() or os.Exit()
//			if n == 0 && err != nil {
//				break
//			}
//		}
//	}()
//
//	// run tests
//	exit := m.Run()
//
//	// after tests
//
//	// restore stdout
//	w.Close()
//	os.Stdout = stdout
//
//	// exit value from tests
//	os.Exit(exit)
//}

//func TestMain(m *testing.M) {
//	os.Setenv("APP_GO_TEST", "true")
//
//	// create a pipe
//	r, w, _ := os.Pipe()
//
//	// backup original outputs
//	stdout := os.Stdout
//	stderr := os.Stderr
//
//	// redirect stdout and stderr to the pipe
//	os.Stdout = w
//	os.Stderr = w
//
//	// Run tests
//	exitCode := m.Run()
//
//	// Cleanup
//	w.Close()
//	buf := make([]byte, 1024*10)
//	var output bytes.Buffer
//
//	runMatch := []byte("=== RUN")
//	failMatch := []byte("--- FAIL:")
//	passMatch := []byte("--- PASS:")
//
//	passedMatch := []byte("PASS")
//	failedMatch := []byte("FAIL")
//
//	colorBlue := []byte("\033[0;36m")
//	colorRed := []byte("\033[0;31m")
//	colorGreen := []byte("\033[0;32m")
//	colorReset := []byte("\033[0m")
//
//	for {
//		n, err := r.Read(buf)
//
//		if n > 0 {
//			output.Write(buf[:n])
//
//			// Process lines
//			lines := bytes.Split(output.Bytes(), []byte("\n"))
//
//			// Keep incomplete line in buffer
//			if !bytes.HasSuffix(output.Bytes(), []byte("\n")) {
//				output.Reset()
//				output.Write(lines[len(lines)-1])
//				lines = lines[:len(lines)-1]
//			} else {
//				output.Reset()
//			}
//
//			for _, line := range lines {
//				line = bytes.TrimLeft(line, " ")
//
//				var color []byte
//				tabs := false
//
//				switch {
//				// === RUN
//				case bytes.Contains(line, runMatch):
//					color = colorBlue
//					tabs = true
//
//				// --- FAIL:
//				case bytes.Contains(line, failMatch):
//					color = colorRed
//					tabs = true
//
//				// --- PASS:
//				case bytes.Contains(line, passMatch):
//					color = colorGreen
//					tabs = true
//
//				// FAIL
//				case bytes.Equal(line, failedMatch):
//					color = colorRed
//
//				// PASS
//				case bytes.Equal(line, passedMatch):
//					color = colorGreen
//
//				// error thrown
//				default:
//					color = nil
//				}
//
//				if color != nil {
//					stdout.Write(color)
//				}
//
//				if tabs {
//					stdout.Write([]byte("\t"))
//				}
//
//				stdout.Write(line)
//
//				if color != nil {
//					stdout.Write(colorReset)
//				}
//
//				stdout.Write([]byte("\n"))
//			}
//		}
//
//		if err != nil {
//			break
//		}
//	}
//
//	os.Stdout = stdout
//	os.Stderr = stderr
//
//	os.Exit(exitCode)
//}

//func TestMain(m *testing.M) {
//	os.Setenv("APP_GO_TEST", "true")
//
//	// create a pipe
//	r, w, _ := os.Pipe()
//
//	// backup original outputs
//	stdout := os.Stdout
//	stderr := os.Stderr
//
//	// redirect stdout and stderr to the pipe
//	os.Stdout = w
//	os.Stderr = w
//
//	// Start goroutine to read and process the output
//	done := make(chan struct{})
//
//	go func() {
//		buf := make([]byte, 1024*2)
//		var output bytes.Buffer
//
//		runMatch := []byte("=== RUN")
//		failMatch := []byte("--- FAIL:")
//		passMatch := []byte("--- PASS:")
//
//		passedMatch := []byte("PASS")
//		failedMatch := []byte("FAIL")
//
//		colorBlue := []byte("\033[0;36m")
//		colorRed := []byte("\033[0;31m")
//		colorGreen := []byte("\033[0;32m")
//		colorReset := []byte("\033[0m")
//
//		for {
//			n, err := r.Read(buf)
//
//			if n > 0 {
//				output.Write(buf[:n])
//
//				// Process lines
//				lines := bytes.Split(output.Bytes(), []byte("\n"))
//
//				// Keep incomplete line in buffer
//				if !bytes.HasSuffix(output.Bytes(), []byte("\n")) {
//					output.Reset()
//					output.Write(lines[len(lines)-1])
//					lines = lines[:len(lines)-1]
//				} else {
//					output.Reset()
//				}
//
//				for _, line := range lines {
//					line = bytes.TrimLeft(line, " ")
//
//					var color []byte
//					tabs := false
//
//					switch {
//					// === RUN
//					case bytes.Contains(line, runMatch):
//						color = colorBlue
//						tabs = true
//
//					// --- FAIL:
//					case bytes.Contains(line, failMatch):
//						color = colorRed
//						tabs = true
//
//					// --- PASS:
//					case bytes.Contains(line, passMatch):
//						color = colorGreen
//						tabs = true
//
//					// FAIL
//					case bytes.Equal(line, failedMatch):
//						color = colorRed
//
//					// PASS
//					case bytes.Equal(line, passedMatch):
//						color = colorGreen
//
//					// error thrown
//					default:
//						color = nil
//					}
//
//					if color != nil {
//						stdout.Write(color)
//					}
//
//					if tabs {
//						stdout.Write([]byte("\t"))
//					}
//
//					stdout.Write(line)
//
//					if color != nil {
//						stdout.Write(colorReset)
//					}
//
//					stdout.Write([]byte("\n"))
//				}
//			}
//
//			if err != nil {
//				break
//			}
//		}
//
//		close(done)
//	}()
//
//	// Run tests
//	exitCode := m.Run()
//
//	// Cleanup
//	w.Close()
//	os.Stdout = stdout
//	os.Stderr = stderr
//
//	// wait for the goroutine to finish
//	<-done
//
//	os.Exit(exitCode)
//}

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

	// read line by line
	for {
		line, err := reader.ReadBytes('\n')

		if len(line) > 0 {
			line = bytes.TrimRight(line, "\n")
			line = bytes.TrimLeft(line, " ")

			var color []byte
			tabs := false

			switch {
			case bytes.Contains(line, runMatch):
				color = colorBlue
				tabs = true
			case bytes.Contains(line, failMatch):
				color = colorRed
				tabs = true
			case bytes.Contains(line, passMatch):
				color = colorGreen
				tabs = true
			case bytes.Contains(line, skipMatch):
				color = colorPurple
				tabs = true
			case bytes.Equal(line, failedMatch):
				color = colorRed
			case bytes.Equal(line, passedMatch):
				color = colorGreen
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
