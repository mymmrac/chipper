/*
Package tests contains all tests that can be executed
*/
package tests

import (
	"errors"

	"github.com/mymmrac/chipper/core"
)

var (
	errUnknownTestName  = errors.New("unknown test name")
	errInvalidArgsCount = errors.New("invalid test args count")
	errInvalidArgType   = errors.New("invalid test arg type")
)

type testName string

// Name of all available tests
const (
	Fibonacci    testName = "fibonacci"
	Factorial    testName = "factorial"
	Trigonometry testName = "trigonometry"
)

// TestCaseArgs represents arguments of test
type TestCaseArgs []interface{}

// TestCase represents test with it's args
type TestCase struct {
	Name testName
	Args TestCaseArgs
}

// TestCases represents slice of test cases
type TestCases []TestCase

func expectOneUint(args TestCaseArgs) (uint, error) {
	if len(args) != 1 {
		return 0, errInvalidArgsCount
	}

	a, ok := args[0].(uint)
	if !ok {
		return 0, errInvalidArgType
	}

	return a, nil
}

// ParseTestCases parses test cases into tests
func ParseTestCases(testList TestCases) (core.Tests, error) {
	ts := make(core.Tests, len(testList))
	for i, tc := range testList {
		switch tc.Name {
		case Fibonacci:
			n, err := expectOneUint(tc.Args)
			if err != nil {
				return nil, err
			}

			ts[i] = NewFibonacciTest(n)
		case Factorial:
			n, err := expectOneUint(tc.Args)
			if err != nil {
				return nil, err
			}

			ts[i] = NewFactorialTest(n)
		case Trigonometry:
			n, err := expectOneUint(tc.Args)
			if err != nil {
				return nil, err
			}

			ts[i] = NewTrigonometryTest(n)
		default:
			return nil, errUnknownTestName
		}
	}

	return ts, nil
}
