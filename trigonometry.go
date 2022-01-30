package main

import (
	"fmt"
	"math"
)

type trigonometryTest struct {
	n      uint
	p      uint
	isDone bool
}

func newTrigonometryTest(n uint) *trigonometryTest {
	return &trigonometryTest{n: n}
}

func (t *trigonometryTest) start() {
	a := 1.0
	for t.p = 0; t.p < t.n; t.p++ {
		a = math.Atan(math.Tan(a + math.E))
	}

	t.isDone = true
}

func (t *trigonometryTest) done() bool {
	return t.isDone
}

func (t *trigonometryTest) progress() float64 {
	return float64(t.p) / float64(t.n)
}

func (t *trigonometryTest) name() string {
	return fmt.Sprintf("trigonometry-%d", t.n)
}
