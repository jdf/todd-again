package todd

import (
	"math"
)

// All geometry takes place on a Cartesian plane where X increases to the right
// and Y increases upward.

// Point is a location in 2D space.
type Point struct {
	X, Y float64
}

// Distance returns the Euclidean distance between two points.
func Distance(p1, p2 *Point) float64 {
	return math.Hypot(p1.X-p2.X, p1.Y-p2.Y)
}

// RectDistance returns the max Euclidean distance between the corners of two rects.
func RectDistance(r1, r2 *Rect) float64 {
	return math.Max(Distance(&r1.Min, &r2.Min), Distance(&r1.Max, &r2.Max))
}

// Rect is an axis-aligned rectangle specified by its bottom left corner and top
// right corner. Please use NewRectFrom.* functions to create new Rects, as they
// enforce ordering of the corners.
type Rect struct {
	Min, Max Point
}

// NewRectFromCorners creates a new Rect from the given corners, enforcing the
// ordering of the corners.
func NewRectFromCorners(corner1, corner2 *Point) *Rect {
	return NewRect(corner1.X, corner1.Y, corner2.X, corner2.Y)
}

// NewRect creates a new Rect from the given axis-aligned lines, enforcing the
// ordering of the corners.
func NewRect(left, bottom, right, top float64) *Rect {
	return &Rect{
		Min: Point{math.Min(left, right), math.Min(bottom, top)},
		Max: Point{math.Max(left, right), math.Max(bottom, top)},
	}
}

// Affine is a transformation matrix.
//
// The matrix is stored in row-major order, with an implicit bottom
// row of [0 0 1], so for Affine m:
//
// ⎡m[0]   m[1]   m[2]⎤
// ⎢m[3]   m[4]   m[5]⎥
// ⎣ 0      0       1 ⎦
type Affine [6]float64

// Identity is the identity transform.
func Identity() *Affine {
	return &Affine{
		1, 0, 0,
		0, 1, 0,
	}
}

// Scale creates a scaling transform.
func Scale(x, y float64) *Affine {
	return &Affine{
		x, 0, 0,
		0, y, 0,
	}
}

// Translation creates a translation transform.
func Translation(x, y float64) *Affine {
	return &Affine{
		1, 0, x,
		0, 1, y,
	}
}

// Rotation creates a rotation transform.
func Rotation(angle float64) *Affine {
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	return &Affine{
		cos, -sin, 0,
		sin, cos, 0,
	}
}

// RotationAround returns a transform that rotates around the given point.
func RotationAround(angle float64, p *Point) *Affine {
	return Compose(Translation(-p.X, -p.Y), Rotation(angle), Translation(p.X, p.Y))
}

// Copy returns a copy of this affine transform.
func (m *Affine) Copy() *Affine {
	return &Affine{
		m[0], m[1], m[2],
		m[3], m[4], m[5],
	}
}

// composeInto returns the composition of two affine transforms, such that "first"
// and "second" are applied in that order. The result Affine may be equal to either
// of the inputs, or nil, in which case a new Affine is allocated.
func composeInto(first, second, target *Affine) *Affine {
	if target == nil {
		target = &Affine{}
	}
	a := second
	b := first
	t0 := a[0]*b[0] + a[1]*b[3]
	t1 := a[0]*b[1] + a[1]*b[4]
	t2 := a[0]*b[2] + a[1]*b[5] + a[2]
	t3 := a[3]*b[0] + a[4]*b[3]
	t4 := a[3]*b[1] + a[4]*b[4]
	t5 := a[3]*b[2] + a[4]*b[5] + a[5]
	target[0] = t0
	target[1] = t1
	target[2] = t2
	target[3] = t3
	target[4] = t4
	target[5] = t5
	return target
}

// Compose returns the composition of any number of affine transforms, such that
// the first Affine is applied first, then the second, then the third, and so on.
func Compose(ts ...*Affine) *Affine {
	result := Identity()
	for _, t := range ts {
		result.Append(t)
	}
	return result
}

// Prepend mutates this affine transform by prepending the other affine transform;
// the resulting Affine has the effect of first applying the other Affine, then
// applying this Affine.
func (m *Affine) Prepend(other *Affine) *Affine {
	return composeInto(other, m, m)
}

// Append mutates this affine transform by appending the other affine transform;
// the resulting Affine has the effect of first applying this Affine, then
// applying the other Affine.
func (m *Affine) Append(other *Affine) *Affine {
	return composeInto(m, other, m)
}

// TransformPoint applies the affine transform to a Point,
// returning a new Point. It does not mutate the original point.
func (m *Affine) TransformPoint(p *Point) *Point {
	return &Point{
		m[0]*p.X + m[1]*p.Y + m[2],
		m[3]*p.X + m[4]*p.Y + m[5],
	}
}

// TransformRect applies the affine transform to a Rect,
// returning a new Rect. It does not mutate the original rect.
func (m *Affine) TransformRect(r *Rect) *Rect {
	return NewRectFromCorners(
		m.TransformPoint(&r.Min),
		m.TransformPoint(&r.Max),
	)
}
