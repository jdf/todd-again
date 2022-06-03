package todd

import (
	"time"

	"golang.org/x/exp/constraints"
)

// Numeric types can be used as coordinate components.
type Numeric interface {
	constraints.Float | constraints.Integer
}

// Entity is a thing in the game that can be updated and drawn.
type Entity interface {
	Update(t *time.Time)
	Render(camera *Camera)
	Intersects(r *Rect) bool
}
