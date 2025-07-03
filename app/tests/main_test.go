package tests

import (
	"bytes"
	"os"
	"strconv"
	"testing"
)

// WARN: all tests must be in this folder, no subfolder authorized
func TestMain(m *testing.M) {
	// before tests
	os.Setenv("APP_GO_TEST", "true")

	// create a pipe between reader and writer
	r, w, _ := os.Pipe()

	// save os.Stdout
	stdout := os.Stdout

	// redirect all os.Stdout into os.Pipe()
	os.Stdout = w

	failMatch := []byte("--- FAIL:")
	passMatch := []byte("--- PASS:")

	//	colorDefault := []byte("\033[36m")
	colorReset := []byte("\033[0m")

	//	colorWhite := []byte("\033[0;37m")
	colorRed := []byte("\033[0;31m")
	colorGreen := []byte("\033[0;32m")

	// start a goroutine
	go func() {
		// 1Ko, standart value for this type of buffer
		buf := make([]byte, 1024*2)

		// infinite loop
		for {
			n, err := r.Read(buf)
			bufToAdd := buf[:n]

			bufSplited := bytes.Split(bufToAdd, []byte("\n"))

			var colored []byte

			// TODO: j'ai 0 dans la boucle pour la valeur bufSplitedLen, pourquoi?
			// TODO: manage inconsistency in test runner print

			// TODO: manage === RUN and errors print color and trimLeft value

			bufSplitedLen := len(bufSplited)

			for i := 0; i < bufSplitedLen; i++ {
				bufSplited[i] = bytes.TrimLeft(bufSplited[i], " ")

				// TODO: add tab before every result
				bufSplited[i] = append(bufSplited[i], '\t')

				if bytes.Contains(bufSplited[i], failMatch) {
					bufSplited[i] = append(colorRed, bufSplited[i]...)
				} else if bytes.Contains(bufSplited[i], passMatch) {
					bufSplited[i] = append(colorGreen, bufSplited[i]...)
				}

				bufSplited[i] = append(bufSplited[i], colorReset...)
				bufSplited[i] = append(bufSplited[i], []byte("bufSplited len:")...)
				bufSplited[i] = append(bufSplited[i], []byte(strconv.Itoa(bufSplitedLen))...)
				bufSplited[i] = append(bufSplited[i], '\n')

				colored = append(colored, bufSplited[i]...)
			}

			// stdout.Write([]byte(strconv.QuoteToASCII(string(colored))))
			stdout.Write(colored)

			// exit at w.Close() or os.Exit()
			if err != nil {
				break
			}
		}
	}()

	// run tests
	exit := m.Run()

	// after tests

	// restore stdout
	w.Close()
	os.Stdout = stdout

	// exit value from tests
	os.Exit(exit)
}
