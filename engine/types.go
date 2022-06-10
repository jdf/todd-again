package engine

import (
	"golang.org/x/exp/constraints"
)

// Numeric types can be used as coordinate components.
type Numeric interface {
	constraints.Float | constraints.Integer
}
