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

// TestStartCallback called before test is started
type TestStartCallback func(name string)

// TestProgressCallback called every time progress updates
type TestProgressCallback func(progress float64)

// TestEndCallback called after test is ended
type TestEndCallback func(name string, duration time.Duration)

// ExecuteTests executes all test and calls callbacks
func ExecuteTests(tests Tests, progressReadTimeout time.Duration, startCallback TestStartCallback,
	progressCallback TestProgressCallback, endCallback TestEndCallback) time.Duration {
	executionStartTime := time.Now()

	for _, t := range tests {
		startCallback(t.Name())
		testStartTime := time.Now()
		testProgressTime := time.Now()

		go t.Start()

		progressCallback(0)

		for !t.Done() {
			if time.Since(testProgressTime) > progressReadTimeout {
				testProgressTime = time.Now()
				progress := t.Progress()
				progressCallback(progress)
			}
		}

		progressCallback(1)
		endCallback(t.Name(), time.Since(testStartTime))
	}

	return time.Since(executionStartTime)
}
