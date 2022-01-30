package tests

import (
	"math/big"

	"github.com/mymmrac/chipper/core"
)

type fibonacciTest struct {
	baseStepTest
}

// NewFibonacciTest creates new fibonacci test
func NewFibonacciTest(n uint) core.Test {
	return &fibonacciTest{baseStepTest{n: n}}
}

func (f *fibonacciTest) Name() string {
	return f.nameBase(Fibonacci)
}

func (f *fibonacciTest) Start() {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for f.p = 0; f.p < f.n; f.p++ {
		tmp := &big.Int{}
		a, b = b, tmp.Add(a, b)
	}

	f.isDone = true
}
