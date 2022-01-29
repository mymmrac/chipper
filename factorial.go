package main

import (
	"fmt"
	"math/big"
)

type factorialTest struct {
	n uint
	p uint

	w      testResultWriter
	isDone bool
}

func newFactorialTest(n uint) *factorialTest {
	f := &factorialTest{n: n}
	f.w = newTestResultWriter(f.name())
	return f
}

func (f *factorialTest) start() {
	a := big.NewInt(1)
	for f.p = 1; f.p <= f.n; f.p++ {
		a.Mul(a, big.NewInt(int64(f.p)))
	}

	f.w.save(a.String())
	f.w.done()

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
