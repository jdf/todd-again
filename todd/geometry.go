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

// Rect is an axis-aligned rectangle specified by its bottom left corner and top
// right corner.
type Rect struct {
	Min, Max Point
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
		cos, sin, 0,
		-sin, cos, 0,
	}
}

// TransformPoint applies the affine transform to a Point,
// returning a new Point. It does not mutate the original point.
func (m *Affine) TransformPoint(p *Point) *Point {
	return &Point{
		m[0]*p.X + m[3]*p.Y + m[2],
		m[1]*p.X + m[4]*p.Y + m[5],
	}
}
