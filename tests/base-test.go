package tests

import (
	"fmt"
)

type baseStepTest struct {
	n      uint
	p      uint
	isDone bool
}

func (b *baseStepTest) nameBase(name testName) string {
	return fmt.Sprintf("%s-%d", name, b.n)
}

func (b *baseStepTest) Done() bool {
	return b.isDone
}

func (b *baseStepTest) Progress() float64 {
	return float64(b.p) / float64(b.n)
}
