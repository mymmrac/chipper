package main

import (
	"fmt"
	"math"
	"strconv"
)

type trigonometryTest struct {
	n uint
	p uint

	w      testResultWriter
	isDone bool
}

func newTrigonometryTest(n uint) *trigonometryTest {
	t := &trigonometryTest{n: n}
	t.w = newTestResultWriter(t.name())
	return t
}

func (t *trigonometryTest) start() {
	a := 1.0
	for t.p = 0; t.p < t.n; t.p++ {
		a = math.Atan(math.Tan(a + math.E))
	}

	t.w.save(strconv.FormatFloat(a, 'f', -1, 64))
	t.w.done()

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
