package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mymmrac/chipper/core"
	"github.com/mymmrac/chipper/tests"
)

const progressReadTimeout = time.Second

func main() {
	testsFast := []core.Test{
		tests.NewFibonacciTest(10e4),
		tests.NewFactorialTest(10e3),
		tests.NewTrigonometryTest(10e5),
	}

	testsSlow := []core.Test{
		tests.NewFibonacciTest(10e5),
		tests.NewFactorialTest(10e4 * 2),
		tests.NewTrigonometryTest(10e7),
	}

	testsAll := append(testsFast, testsSlow...)

	startTime := time.Now()

	for _, t := range testsAll {
		fmt.Printf("Starting test %s\n", t.Name())
		testProgressTime := time.Now()
		testStartTime := time.Now()

		go t.Start()

		firstTime := true
		for !t.Done() {
			if time.Since(testProgressTime) > progressReadTimeout {
				testProgressTime = time.Now()
				progress := t.Progress()

				if firstTime {
					firstTime = false
				} else {
					fmt.Print("\033[1A\r\033[K")
				}

				fmt.Printf("Progress: %s%%\n", strconv.FormatFloat(progress*100, 'f', 2, 64))
			}
		}

		if !firstTime {
			fmt.Print("\033[1A\r\033[K")
		}

		fmt.Printf("Progress: %s%%\n", strconv.FormatFloat(100, 'f', 2, 64))
		fmt.Printf("Test %s done in %s\n", t.Name(), time.Since(testStartTime))
		fmt.Println()
	}

	fmt.Printf("All tests done in %s\n", time.Since(startTime))
}
