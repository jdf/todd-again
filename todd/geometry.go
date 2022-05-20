package todd

// All geometry takes place on a Cartesian plane where X increases to the right
// and Y increases upward.

// Point is a location in 2D space.
type Point struct {
	X, Y float64
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

// Identity is the identity transformation.
func Identity() *Affine {
	return &Affine{1, 0, 0, 0, 1, 0}
}

// TransformPoint applies the affine transformation to a point,
// returning a new Point It does not mutate the original point.
func (m *Affine) TransformPoint(p *Point) *Point {
	return &Point{
		m[0]*p.X + m[3]*p.Y + m[2],
		m[1]*p.X + m[4]*p.Y + m[5],
	}
}
