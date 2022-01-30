package tests

import (
	"math/big"

	"github.com/mymmrac/chipper/core"
)

type factorialTest struct {
	baseStepTest
}

// NewFactorialTest creates new factorial test
func NewFactorialTest(n uint) core.Test {
	return &factorialTest{baseStepTest{n: n}}
}

func (f *factorialTest) Name() string {
	return f.nameBase(Factorial)
}

func (f *factorialTest) Start() {
	a := big.NewInt(1)
	for f.p = 1; f.p <= f.n; f.p++ {
		a.Mul(a, big.NewInt(int64(f.p)))
	}

	f.isDone = true
}
