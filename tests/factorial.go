package tests

import (
	"fmt"
	"math/big"

	"github.com/mymmrac/chipper/core"
)

type factorialTest struct {
	n      uint
	p      uint
	isDone bool
}

// NewFactorialTest creates new factorial test
func NewFactorialTest(n uint) core.Test {
	return &factorialTest{n: n}
}

func (f *factorialTest) Name() string {
	return fmt.Sprintf("factorial-%d", f.n)
}

func (f *factorialTest) Start() {
	a := big.NewInt(1)
	for f.p = 1; f.p <= f.n; f.p++ {
		a.Mul(a, big.NewInt(int64(f.p)))
	}

	f.isDone = true
}

func (f *factorialTest) Done() bool {
	return f.isDone
}

func (f *factorialTest) Progress() float64 {
	return float64(f.p) / float64(f.n)
}
