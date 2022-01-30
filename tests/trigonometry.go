package tests

import (
	"fmt"
	"math"

	"github.com/mymmrac/chipper/core"
)

type trigonometryTest struct {
	n      uint
	p      uint
	isDone bool
}

// NewTrigonometryTest creates new trigonometry test
func NewTrigonometryTest(n uint) core.Test {
	return &trigonometryTest{n: n}
}

func (t *trigonometryTest) Name() string {
	return fmt.Sprintf("trigonometry-%d", t.n)
}

func (t *trigonometryTest) Start() {
	a := 1.0
	for t.p = 0; t.p < t.n; t.p++ {
		a = math.Atan(math.Tan(a + math.E))
	}

	t.isDone = true
}

func (t *trigonometryTest) Done() bool {
	return t.isDone
}

func (t *trigonometryTest) Progress() float64 {
	return float64(t.p) / float64(t.n)
}
