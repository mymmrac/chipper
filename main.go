package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type test interface {
	start()
	done() bool
	progress() float64
	name() string
}

const progressReadTimeout = time.Second

type testResultWriter struct {
	file *os.File
}

const testResultsDir = "./results"

func newTestResultWriter(name string) testResultWriter {
	file, err := os.OpenFile(filepath.Join(testResultsDir, name+".txt"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	return testResultWriter{file: file}
}

func (w testResultWriter) save(result string) {
	_, err := w.file.WriteString(result + "\n")
	if err != nil {
		panic(err)
	}
}

func (w testResultWriter) done() {
	_ = w.file.Close()
}

func main() {
	if err := os.Mkdir(testResultsDir, 0777); err != nil && !errors.Is(err, os.ErrExist) {
		panic(err)
	}

	tests := []test{
		newFibonacciTest(10e4),
		newFactorialTest(10e3),
		newTrigonometryTest(10e5),

		newFibonacciTest(10e5),
		newFactorialTest(10e4 * 2),
		newTrigonometryTest(10e7),
	}

	startTime := time.Now()

	for _, t := range tests {
		fmt.Printf("Starting test %s\n", t.name())
		testProgressTime := time.Now()
		testStartTime := time.Now()

		go t.start()

		firstTime := true
		for !t.done() {
			if time.Since(testProgressTime) > progressReadTimeout {
				testProgressTime = time.Now()
				progress := t.progress()

				if firstTime {
					firstTime = false
				} else {
					fmt.Print("\033[1A\r\033[K")
				}

				fmt.Printf("Progress: %s%%\n", strconv.FormatFloat(progress*100, 'f', 2, 64))
			}
		}

		if firstTime {
			firstTime = false
		} else {
			fmt.Print("\033[1A\r\033[K")
		}

		fmt.Printf("Progress: %s%%\n", strconv.FormatFloat(100, 'f', 2, 64))
		fmt.Printf("Test %s done in %s\n", t.name(), time.Since(testStartTime))
		fmt.Println()
	}

	fmt.Printf("All tests done in %s\n", time.Since(startTime))
}
