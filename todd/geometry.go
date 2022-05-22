package todd

import "math"

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
	return NewRectFromEdges(corner1.X, corner1.Y, corner2.X, corner2.Y)
}

// NewRectFromEdges creates a new Rect from the given axis-aligned lines, enforcing the
// ordering of the corners.
func NewRectFromEdges(x1, y1, x2, y2 float64) *Rect {
	return &Rect{
		Min: Point{math.Min(x1, x2), math.Min(y1, y2)},
		Max: Point{math.Max(x1, x2), math.Max(y1, y2)},
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

// MakeScale creates a scaling transform.
func MakeScale(x, y float64) *Affine {
	return &Affine{
		x, 0, 0,
		0, y, 0,
	}
}

// MakeTranslate creates a translation transform.
func MakeTranslate(x, y float64) *Affine {
	return &Affine{
		1, 0, x,
		0, 1, y,
	}
}

// MakeRotate creates a rotation transform.
func MakeRotate(angle float64) *Affine {
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	return &Affine{
		cos, -sin, 0,
		sin, cos, 0,
	}
}

// Compose returns the composition of two affine transforms, such that "first"
// and "second" are applied in that order.
func Compose(first, second *Affine) *Affine {
	a := second
	b := first

	return &Affine{
		a[0]*b[0] + a[1]*b[3],
		a[0]*b[1] + a[1]*b[4],
		a[0]*b[2] + a[1]*b[5] + a[2],
		a[3]*b[0] + a[4]*b[3],
		a[3]*b[1] + a[4]*b[4],
		a[3]*b[2] + a[4]*b[5] + a[5],
	}
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
	return &Rect{
		*m.TransformPoint(&r.Min),
		*m.TransformPoint(&r.Max),
	}
}
