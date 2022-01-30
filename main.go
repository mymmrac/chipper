package main

import (
	"fmt"
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

func main() {
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
