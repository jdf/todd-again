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
			m := MakeScale(s[0], s[1])
			want := Point{s[0] * p.X, s[1] * p.Y}
			got := *m.TransformPoint(p)
			if got != want {
				t.Errorf("Got %v, want %v", got, want)
			}
		}
	}
}

func TestTranslate(t *testing.T) {
	for _, p := range []*Point{{0, 0}, {1, 2}, {2, 1}, {0, 0}, {4.1, 1.4}} {
		for _, s := range [][2]float64{{0, 0}, {1, 2}, {2, 1}, {-1, -1.5}} {
			m := MakeTranslate(s[0], s[1])
			want := Point{s[0] + p.X, s[1] + p.Y}
			got := *m.TransformPoint(p)
			if got != want {
				t.Errorf("Got %v, want %v", got, want)
			}
		}
	}
}

func TestRotate(t *testing.T) {
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
		m := MakeRotate(tt.angle)
		got := *m.TransformPoint(&tt.p)
		if Distance(&got, &tt.want) > epsilon {
			t.Errorf("Got %v, want %v when rotating %v by %v", got, tt.want, tt.p, tt.angle)
		}
	}
}

func snapTenth(f float64) float64 {
	return math.Round(f*10) / 10
}

func rp(p *Point) *Point {
	return &Point{snapTenth(p.X), snapTenth(p.Y)}
}

type labeledAffine struct {
	label string
	a     *Affine
}

var transforms = []interface{}{
	labeledAffine{"identity", Identity()},
	labeledAffine{"translate(1, 2)", MakeTranslate(1, 2)},
	labeledAffine{"scale(3, 4)", MakeScale(3, 4)},
	labeledAffine{"rotate(-π/2)", MakeRotate(-math.Pi / 2)},
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
					*rp(got), *rp(want),
					first.label, *rp(first.a.TransformPoint(p)),
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
					*rp(got), *rp(want),
					first.label, *rp(first.a.TransformPoint(p)),
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
					*rp(got), *rp(want),
					first.label, *rp(first.a.TransformPoint(p)),
					second.label)
			}
		}
	}
}
func TestNewRectFromCorners(t *testing.T) {
	type tc struct {
		a, b *Point
		want *Rect
	}
	for _, tt := range []tc{
		{&Point{0, 0}, &Point{1, 1}, &Rect{Point{0, 0}, Point{1, 1}}},
		{&Point{0, 1}, &Point{1, 0}, &Rect{Point{0, 0}, Point{1, 1}}},
		{&Point{1, 1}, &Point{0, 0}, &Rect{Point{0, 0}, Point{1, 1}}},
		{&Point{1, 0}, &Point{0, 1}, &Rect{Point{0, 0}, Point{1, 1}}},
	} {
		got := NewRectFromCorners(tt.a, tt.b)
		if *got != *tt.want {
			t.Errorf("Got %v, want %v", got, tt.want)
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
		{NewRectFromEdges(0, 0, 1, 1), Identity(), NewRectFromEdges(0, 0, 1, 1)},
	} {
		got := tt.m.TransformRect(tt.r)
		if *got != *tt.want {
			t.Errorf("Got %v, want %v", got, tt.want)
		}
	}
}
