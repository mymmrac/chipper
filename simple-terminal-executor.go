package main

import (
	"fmt"
	"strconv"
	"time"
)

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
