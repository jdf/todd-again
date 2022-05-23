package todd

import (
	"math"
	"testing"

	"github.com/schwarmco/go-cartesian-product"
)

const epsilon = 1e-6

func TestIdentity(t *testing.T) {
	m := Identity()
	for _, p := range []*Point{{0, 0}, {1, 2}, {2, 1}, {0, 0}, {4.1, 1.4}} {
		if *m.TransformPoint(p) != *p {
			t.Errorf("Got %v, expected %v", m.TransformPoint(p), p)
		}
	}
}

func TestScale(t *testing.T) {
	for _, p := range []*Point{{0, 0}, {1, 2}, {2, 1}, {0, 0}, {4.1, 1.4}} {
		for _, s := range [][2]float64{{0, 0}, {1, 2}, {2, 1}, {-1, -1.5}} {
			m := Scale(s[0], s[1])
			want := Point{s[0] * p.X, s[1] * p.Y}
			got := *m.TransformPoint(p)
			if got != want {
				t.Errorf("Got %v, want %v", got, want)
			}
		}
	}
}

func TestTranslation(t *testing.T) {
	for _, p := range []*Point{{0, 0}, {1, 2}, {2, 1}, {0, 0}, {4.1, 1.4}} {
		for _, s := range [][2]float64{{0, 0}, {1, 2}, {2, 1}, {-1, -1.5}} {
			m := Translation(s[0], s[1])
			want := Point{s[0] + p.X, s[1] + p.Y}
			got := *m.TransformPoint(p)
			if got != want {
				t.Errorf("Got %v, want %v", got, want)
			}
		}
	}
}

func TestRotation(t *testing.T) {
	type tc struct {
		p     Point
		angle float64
		want  Point
	}
	for _, tt := range []tc{
		{Point{0, 0}, 0, Point{0, 0}},
		{Point{1, 2}, math.Pi / 2, Point{-2, 1}},
		{Point{2, 1}, math.Pi, Point{-2, -1}},
		{Point{2, 1}, -math.Pi / 2, Point{1, -2}},
		{Point{0, 0}, math.Pi / 2, Point{0, 0}},
		{Point{4.1, 1.4}, math.Pi / 2, Point{-1.4, 4.1}},
	} {
		m := Rotation(tt.angle)
		got := *m.TransformPoint(&tt.p)
		if Distance(&got, &tt.want) > epsilon {
			t.Errorf("Got %v, want %v when rotating %v by %v", got, tt.want, tt.p, tt.angle)
		}
	}
}

func snapTenth(f float64) float64 {
	return math.Round(f*10) / 10
}

func roundPoint(p *Point) *Point {
	return &Point{snapTenth(p.X), snapTenth(p.Y)}
}

func roundRect(r *Rect) *Rect {
	return &Rect{*roundPoint(&r.Min), *roundPoint(&r.Max)}
}

type labeledAffine struct {
	label string
	a     *Affine
}

var transforms = []interface{}{
	labeledAffine{"identity", Identity()},
	labeledAffine{"translate(1, 2)", Translation(1, 2)},
	labeledAffine{"scale(3, 4)", Scale(3, 4)},
	labeledAffine{"rotate(-π/2)", Rotation(-math.Pi / 2)},
}

func TestCompose(t *testing.T) {
	for _, p := range []*Point{{1, 2}, {2, 1}, {0, 0}} {
		for affinePair := range cartesian.Iter(transforms, transforms) {
			first := affinePair[0].(labeledAffine)
			second := affinePair[1].(labeledAffine)
			want := second.a.TransformPoint(first.a.TransformPoint(p))
			got := Compose(first.a, second.a).TransformPoint(p)
			if Distance(got, want) > epsilon {
				t.Errorf(
					"Got %v, want %v with first %s (= %v) then %s",
					*roundPoint(got), *roundPoint(want),
					first.label, *roundPoint(first.a.TransformPoint(p)),
					second.label)
			}
		}
	}
}

func TestPrepend(t *testing.T) {
	for _, p := range []*Point{{1, 2}, {2, 1}, {0, 0}} {
		for affinePair := range cartesian.Iter(transforms, transforms) {
			first := affinePair[0].(labeledAffine)
			second := affinePair[1].(labeledAffine)
			firstCopy := first.a.Copy()
			secondCopy := second.a.Copy()
			want := secondCopy.TransformPoint(firstCopy.TransformPoint(p))
			secondCopy.Prepend(firstCopy)
			got := secondCopy.TransformPoint(p)
			if Distance(got, want) > epsilon {
				t.Errorf(
					"Got %v, want %v with first %s (= %v) then %s",
					*roundPoint(got), *roundPoint(want),
					first.label, *roundPoint(first.a.TransformPoint(p)),
					second.label)
			}
		}
	}
}

func TestAppend(t *testing.T) {
	for _, p := range []*Point{{1, 2}, {2, 1}, {0, 0}} {
		for affinePair := range cartesian.Iter(transforms, transforms) {
			first := affinePair[0].(labeledAffine)
			second := affinePair[1].(labeledAffine)
			firstCopy := first.a.Copy()
			secondCopy := second.a.Copy()
			want := secondCopy.TransformPoint(firstCopy.TransformPoint(p))
			firstCopy.Append(secondCopy)
			got := firstCopy.TransformPoint(p)
			if Distance(got, want) > epsilon {
				t.Errorf(
					"Got %v, want %v with first %s (= %v) then %s",
					*roundPoint(got), *roundPoint(want),
					first.label, *roundPoint(first.a.TransformPoint(p)),
					second.label)
			}
		}
	}
}

func TestRotationAround(t *testing.T) {
	type tc struct {
		p    Point
		want *Rect
	}

	r := NewRect(0, 0, 2, 1)
	angle := -math.Pi / 2
	for _, tt := range []tc{
		{Point{0, 0}, NewRect(0, 0, 1, -2)},
		{Point{2, 1}, NewRect(1, 1, 2, 3)},
		{Point{1, .5}, NewRect(.5, -.5, 1.5, 1.5)},
	} {
		got := roundRect(RotationAround(angle, &tt.p).TransformRect(r))
		if *got != *tt.want {
			t.Errorf("Got %v, want %v when rotating around %v", got, tt.want, tt.p)
		}
	}
}

func TestNewRectFromCorners(t *testing.T) {
	type tc struct {
		a, b *Point
	}
	want := &Rect{Point{0, 0}, Point{1, 1}}
	for _, tt := range []tc{
		{&Point{0, 0}, &Point{1, 1}},
		{&Point{0, 1}, &Point{1, 0}},
		{&Point{1, 1}, &Point{0, 0}},
		{&Point{1, 0}, &Point{0, 1}},
	} {
		if got := NewRectFromCorners(tt.a, tt.b); *got != *want {
			t.Errorf("Got %v, want %v", got, want)
		}
	}
}

func TestNewRectFromEdges(t *testing.T) {
	want := &Rect{Point{0, 0}, Point{1, 1}}
	for _, tt := range [][4]float64{
		{0, 0, 1, 1},
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{1, 1, 0, 0},
	} {
		if got := NewRect(tt[0], tt[1], tt[2], tt[3]); *got != *want {
			t.Errorf("Got %v, want %v from %v", got, want, tt)
		}
	}
}

func TestTransformRect(t *testing.T) {
	type tc struct {
		r    *Rect
		m    *Affine
		want *Rect
	}
	for _, tt := range []tc{
		{NewRect(0, 0, 1, 1), Identity(), NewRect(0, 0, 1, 1)},
	} {
		got := tt.m.TransformRect(tt.r)
		if *got != *tt.want {
			t.Errorf("Got %v, want %v", got, tt.want)
		}
	}
}
