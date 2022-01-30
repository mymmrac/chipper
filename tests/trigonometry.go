package tests

import (
	"math"

	"github.com/mymmrac/chipper/core"
)

type trigonometryTest struct {
	baseStepTest
}

// NewTrigonometryTest creates new trigonometry test
func NewTrigonometryTest(n uint) core.Test {
	return &trigonometryTest{baseStepTest{n: n}}
}

func (t *trigonometryTest) Name() string {
	return t.nameBase(Trigonometry)
}

func (t *trigonometryTest) Start() {
	a := 1.0
	for t.p = 0; t.p < t.n; t.p++ {
		a = math.Atan(math.Tan(a + math.E))
	}

	t.isDone = true
}
