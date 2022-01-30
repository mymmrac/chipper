/*
Package core contains fundamental things related to tests
*/
package core

import (
	"time"
)

// Test represents test that can be executed
type Test interface {
	// Name returns uniq name of the test
	Name() string

	// Start starts execution of the test
	Start()

	// Done returns true only if test execution ended
	Done() bool

	// Progress returns current progress of test execution in range [0,1]
	Progress() float64
}

// Tests represents slice of tests
type Tests []Test

// TestExecutor represents handler of test execution
type TestExecutor interface {
	// OnExecutionStart called before execution is started
	OnExecutionStart(count int)

	// OnExecutionEnd called after execution is ended
	OnExecutionEnd(duration time.Duration)

	// OnTestStart called before test is started
	OnTestStart(name string, index int)

	// OnTestProgress called every time progress updates
	OnTestProgress(progress float64)

	// OnTestEnd called after test is ended
	OnTestEnd(duration time.Duration)
}

// ExecuteTests executes all test and calls callbacks
func ExecuteTests(tests Tests, progressReadInterval time.Duration, executor TestExecutor) {
	executor.OnExecutionStart(len(tests))
	executionStartTime := time.Now()

	for i, t := range tests {
		executor.OnTestStart(t.Name(), i)
		testStartTime := time.Now()
		testProgressTime := time.Now()

		go t.Start()

		executor.OnTestProgress(0)

		for !t.Done() {
			if time.Since(testProgressTime) > progressReadInterval {
				testProgressTime = time.Now()
				progress := t.Progress()

				executor.OnTestProgress(progress)
			}
		}

		executor.OnTestProgress(1)
		executor.OnTestEnd(time.Since(testStartTime))
	}

	executor.OnExecutionEnd(time.Since(executionStartTime))
}
