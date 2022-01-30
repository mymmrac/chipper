package tests

import (
	"fmt"
	"math/big"

	"github.com/mymmrac/chipper/core"
)

type fibonacciTest struct {
	n      uint
	p      uint
	isDone bool
}

// NewFibonacciTest creates new fibonacci test
func NewFibonacciTest(n uint) core.Test {
	return &fibonacciTest{n: n}
}

func (f *fibonacciTest) Name() string {
	return fmt.Sprintf("fibonnacci-%d", f.n)
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

func (f *fibonacciTest) Done() bool {
	return f.isDone
}

func (f *fibonacciTest) Progress() float64 {
	return float64(f.p) / float64(f.n)
}
