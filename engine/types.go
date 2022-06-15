package engine

import (
	"golang.org/x/exp/constraints"
)

// Numeric types can be used as coordinate components.
type Numeric interface {
	constraints.Float | constraints.Integer
}

type Stack[T any] []T

func (t *Stack[T]) Push(v T) {
	*t = append(*t, v)
	// log.Printf("after push stack: %v", t)
}

func (t *Stack[T]) Pop() {
	*t = (*t)[:len(*t)-1]
	// log.Printf("after pop stack: %v", t)
}

func (t *Stack[T]) Top() T {
	return (*t)[len(*t)-1]
}

func (t *Stack[T]) Empty() bool {
	return len(*t) == 0
}
