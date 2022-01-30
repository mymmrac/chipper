package main

import (
	"fmt"
	"math/big"
)

type fibonacciTest struct {
	n      uint
	p      uint
	isDone bool
}

func newFibonacciTest(n uint) *fibonacciTest {
	return &fibonacciTest{n: n}
}

func (f *fibonacciTest) start() {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for f.p = 0; f.p < f.n; f.p++ {
		tmp := &big.Int{}
		a, b = b, tmp.Add(a, b)
	}

	f.isDone = true
}

func (f *fibonacciTest) done() bool {
	return f.isDone
}

func (f *fibonacciTest) progress() float64 {
	return float64(f.p) / float64(f.n)
}

func (f *fibonacciTest) name() string {
	return fmt.Sprintf("fibonnacci-%d", f.n)
}
