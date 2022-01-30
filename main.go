package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"

	"github.com/mymmrac/chipper/core"
	"github.com/mymmrac/chipper/tests"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read config: %v\n", err)
		os.Exit(1)
	}

	var tcs tests.TestCases
	if err := viper.UnmarshalKey("test-case-list", &tcs); err != nil {
		fmt.Printf("Failed to get test cases: %v\n", err)
		os.Exit(1)
	}

	testList, err := tests.ParseTestCases(tcs)
	if err != nil {
		fmt.Printf("Failed to parce test cases: %v\n", err)
		os.Exit(1)
	}

	if len(testList) == 0 {
		fmt.Println("No test cases found")
		os.Exit(1)
	}

	progressReadInterval := viper.GetDuration("progress-read-interval")
	if progressReadInterval == 0 {
		fmt.Println("Progress read interval can't be 0")
		os.Exit(1)
	}

	core.ExecuteTests(testList, progressReadInterval, &simpleTerminalExecutor{})
}

type simpleTerminalExecutor struct {
	testCount   int
	currentTest string
}

func (s *simpleTerminalExecutor) OnExecutionStart(count int) {
	fmt.Printf("Tests to be executed: %d\n\n", count)
	s.testCount = count
}

func (s *simpleTerminalExecutor) OnExecutionEnd(duration time.Duration) {
	fmt.Printf("All tests done in %s\n", duration)
}

func (s *simpleTerminalExecutor) OnTestStart(name string, index int) {
	fmt.Printf("[%d/%d] Starting test %s\n", index+1, s.testCount, name)
	s.currentTest = name
}

func (s *simpleTerminalExecutor) OnTestProgress(progress float64) {
	if progress != 0 {
		fmt.Print("\033[1A\r\033[K")
	}
	fmt.Printf("Progress: %s%%\n", strconv.FormatFloat(progress*100, 'f', 2, 64))
}

func (s *simpleTerminalExecutor) OnTestEnd(duration time.Duration) {
	fmt.Printf("Test %s done in %s\n", s.currentTest, duration)
	fmt.Println()
}
