package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mymmrac/chipper/core"
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

	totalDuration := core.ExecuteTests(testList, progressReadTimeout,
		func(testName string) {
			fmt.Printf("Starting test %s\n", testName)
		},
		func(progress float64) {
			if progress != 0 {
				fmt.Print("\033[1A\r\033[K")
			}
			fmt.Printf("Progress: %s%%\n", strconv.FormatFloat(progress*100, 'f', 2, 64))
		},
		func(testName string, testDuration time.Duration) {
			fmt.Printf("Test %s done in %s\n", testName, testDuration)
			fmt.Println()
		})

	fmt.Printf("All tests done in %s\n", totalDuration)
}
