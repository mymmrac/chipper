/*
Package core contains fundamental things related to tests
*/
package core

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
