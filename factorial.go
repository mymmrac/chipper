package main

import (
	"fmt"
	"math/big"
)

type factorialTest struct {
	n      uint
	p      uint
	isDone bool
}

func newFactorialTest(n uint) *factorialTest {
	return &factorialTest{n: n}
}

func (f *factorialTest) start() {
	a := big.NewInt(1)
	for f.p = 1; f.p <= f.n; f.p++ {
		a.Mul(a, big.NewInt(int64(f.p)))
	}

	f.isDone = true
}

func (f *factorialTest) done() bool {
	return f.isDone
}

func (f *factorialTest) progress() float64 {
	return float64(f.p) / float64(f.n)
}

func (f *factorialTest) name() string {
	return fmt.Sprintf("factorial-%d", f.n)
}
