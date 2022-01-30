package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mymmrac/chipper/tests"
)

const progressReadTimeout = time.Second

var testCaseList = tests.TestCases{
	{
		Name: tests.Fibonacci,
		Args: tests.TestCaseArgs{uint(10e4)},
	},
	{
		Name: tests.Factorial,
		Args: tests.TestCaseArgs{uint(10e3)},
	},
	{
		Name: tests.Trigonometry,
		Args: tests.TestCaseArgs{uint(10e5)},
	},
	{
		Name: tests.Fibonacci,
		Args: tests.TestCaseArgs{uint(10e5)},
	},
	{
		Name: tests.Factorial,
		Args: tests.TestCaseArgs{uint(10e4 * 2)},
	},
	{
		Name: tests.Trigonometry,
		Args: tests.TestCaseArgs{uint(10e7)},
	},
}

func main() {
	testList, err := tests.ParseTestCases(testCaseList)
	if err != nil {
		fmt.Printf("Test cases: %v\n", err)
		os.Exit(1)
	}

	startTime := time.Now()

	for _, t := range testList {
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
