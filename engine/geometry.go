package engine

import (
	"fmt"
	"math"
)

// All geometry takes place on a Cartesian plane where X increases to the right
// and Y increases upward.

// Vec2 is a 2D vector. We treat this structure as vectors sometimes, and as
// points sometimes.
type Vec2 struct {
	X, Y float64
}

// Vec creates a Vec2 from any 2 numeric values.
func Vec[T, U Numeric](x T, y U) *Vec2 {
	return &Vec2{float64(x), float64(y)}
}

// Copy returns a deep copy of this Vec2.
func (v *Vec2) Copy() *Vec2 {
	return &Vec2{v.X, v.Y}
}

// AddToSelf adds the given Vec2 to this.
func (v *Vec2) AddToSelf(d *Vec2) {
	v.X += d.X
	v.Y += d.Y
}

// Plus adds the given Vec2 to this Vec2 and returns the result
// as a new Vec2.
func (v *Vec2) Plus(t *Vec2) *Vec2 {
	return &Vec2{v.X + t.X, v.Y + t.Y}
}

// Minus subtracts the given Vec2 from this Vec2 and returns the result
// as a new Vec2.
func (v *Vec2) Minus(t *Vec2) *Vec2 {
	return &Vec2{v.X - t.X, v.Y - t.Y}
}

// Distance returns the Euclidean distance between two points.
func Distance(p1, p2 *Vec2) float64 {
	return math.Hypot(p1.X-p2.X, p1.Y-p2.Y)
}

// Negate returns the component-wise negation of this Vec2.
func (v *Vec2) Negate() *Vec2 {
	return &Vec2{-v.X, -v.Y}
}

// Div component-wise divides this Vec2 by the given Vec2.
func (v *Vec2) Div(d *Vec2) *Vec2 {
	return &Vec2{v.X / d.X, v.Y / d.Y}
}

func (v *Vec2) Normalize() *Vec2 {
	r := v.Copy()
	r.NormSelf()
	return r
}

func (v *Vec2) NormSelf() {
	length := math.Hypot(v.X, v.Y)
	v.X, v.Y = v.X/length, v.Y/length
}

func Dot(v1, v2 *Vec2) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Equals returns true if both vecs have the same coordinates.
func (v *Vec2) Equals(other *Vec2) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v *Vec2) String() string {
	return fmt.Sprintf("(%0.2f, %0.2f)", v.X, v.Y)
}

// Rect is an axis-aligned rectangle specified by its bottom left corner and top
// right corner. Please use NewRect.* funcs to create new Rects, as they
// enforce ordering of the corners.
type Rect struct {
	Min, Max Vec2
}

// Copy returns a deep copy of this Rect.
func (r *Rect) Copy() *Rect {
	return &Rect{*r.Min.Copy(), *r.Max.Copy()}
}

// NewRectFromCorners creates a new Rect from the given corners, enforcing the
// ordering of the corners.
func NewRectFromCorners(corner1, corner2 *Vec2) *Rect {
	return NewRect(corner1.X, corner1.Y, corner2.X, corner2.Y)
}

// NewRect creates a new Rect from the given axis-aligned lines, enforcing the
// ordering of the corners.
func NewRect[T Numeric](left, bottom, right, top T) *Rect {
	fl := float64(left)
	fr := float64(right)
	ft := float64(top)
	fb := float64(bottom)
	return &Rect{
		Min: Vec2{math.Min(fl, fr), math.Min(fb, ft)},
		Max: Vec2{math.Max(fl, fr), math.Max(fb, ft)},
	}
}

// AddToSelf applies the given translation to this rectangle.
func (r *Rect) AddToSelf(t *Vec2) {
	r.Min.AddToSelf(t)
	r.Max.AddToSelf(t)
}

// Intersects returns true if the given rectangle intersects this one, where
// coincident edges are considered to intersect.
func (r *Rect) Intersects(other *Rect) bool {
	return r.Min.X <= other.Max.X && r.Max.X >= other.Min.X &&
		r.Min.Y <= other.Max.Y && r.Max.Y >= other.Min.Y
}

// Center returns the center of this Rect.
func (r *Rect) Center() *Vec2 {
	return &Vec2{
		(r.Min.X + r.Max.X) / 2,
		(r.Min.Y + r.Max.Y) / 2,
	}
}

// Left returns the leftmost point of this Rect.
func (r *Rect) Left() float64 {
	return r.Min.X
}

// Right returns the rightmost point of this Rect.
func (r *Rect) Right() float64 {
	return r.Max.X
}

// Bottom returns the bottommost point of this Rect.
func (r *Rect) Bottom() float64 {
	return r.Min.Y
}

// Top returns the topmost point of this Rect.
func (r *Rect) Top() float64 {
	return r.Max.Y
}

// Width returns the width of this rectangle.
func (r *Rect) Width() float64 {
	return r.Max.X - r.Min.X
}

// Height returns the height of this rectangle.
func (r *Rect) Height() float64 {
	return r.Max.Y - r.Min.Y
}

// Size returns the dimensions of this Rect.
func (r *Rect) Size() *Vec2 {
	return &Vec2{r.Max.X - r.Min.X, r.Max.Y - r.Min.Y}
}

func (r *Rect) Inset(delta *Vec2) *Rect {
	return &Rect{
		Min: *r.Min.Plus(delta),
		Max: *r.Max.Minus(delta),
	}
}

// Corners returns the four corners of this Rect from left-bottom counterclockwise.
func (r *Rect) Corners() [4]Vec2 {
	return [4]Vec2{
		Vec2{r.Min.X, r.Min.Y},
		Vec2{r.Min.X, r.Max.Y},
		Vec2{r.Max.X, r.Max.Y},
		Vec2{r.Max.X, r.Min.Y},
	}
}

func (r *Rect) String() string {
	return "Rect[" + r.Min.String() + "->" + r.Max.String() + "]"
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
func Scale(s *Vec2) *Affine {
	return &Affine{
		s.X, 0, 0,
		0, s.Y, 0,
	}
}

// UniformScale creates a uniform scaling transform.
func UniformScale(s float64) *Affine {
	return &Affine{
		s, 0, 0,
		0, s, 0,
	}
}

// Translation creates a translation transform.
func Translation(t *Vec2) *Affine {
	return &Affine{
		1, 0, t.X,
		0, 1, t.Y,
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
func RotationAround(angle float64, p *Vec2) *Affine {
	return Compose(Translation(&Vec2{-p.X, -p.Y}), Rotation(angle), Translation(&Vec2{p.X, p.Y}))
}

// Copy returns a copy of this affine transform.
func (m *Affine) Copy() *Affine {
	return &Affine{
		m[0], m[1], m[2],
		m[3], m[4], m[5],
	}
}

// Inverse returns the inverse of this affine transform.
func (m *Affine) Inverse() *Affine {
	oneOverDet := 1.0 / (m[0]*m[4] - m[1]*m[3])
	return &Affine{
		m[4] * oneOverDet, -m[1] * oneOverDet, (m[1]*m[5] - m[2]*m[4]) * oneOverDet,
		-m[3] * oneOverDet, m[0] * oneOverDet, (m[2]*m[3] - m[0]*m[5]) * oneOverDet,
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

// TransformVec2 applies the affine transform to a Vec2,
// returning a new Vec2. It does not mutate the original point.
func (m *Affine) TransformVec2(p *Vec2) *Vec2 {
	return &Vec2{
		m[0]*p.X + m[1]*p.Y + m[2],
		m[3]*p.X + m[4]*p.Y + m[5],
	}
}

// Transform applies the affine transform to the given coordinate.
func (m *Affine) Transform(x, y float64) (float64, float64) {
	return m[0]*x + m[1]*y + m[2], m[3]*x + m[4]*y + m[5]

}

// TransformRect applies the affine transform to a Rect,
// returning a new Rect. It does not mutate the original rect.
func (m *Affine) TransformRect(r *Rect) *Rect {
	return NewRectFromCorners(
		m.TransformVec2(&r.Min),
		m.TransformVec2(&r.Max),
	)
}
