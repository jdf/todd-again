package todd

import (
	"math"
	"testing"

	"github.com/schwarmco/go-cartesian-product"
)

const epsilon = 1e-6

func TestIdentity(t *testing.T) {
	m := Identity()
	for _, p := range []*Vec2{{0, 0}, {1, 2}, {2, 1}, {0, 0}, {4.1, 1.4}} {
		if *m.TransformVec2(p) != *p {
			t.Errorf("Got %v, expected %v", m.TransformVec2(p), p)
		}
	}
}

func TestScale(t *testing.T) {
	for _, p := range []*Vec2{{0, 0}, {1, 2}, {2, 1}, {0, 0}, {4.1, 1.4}} {
		for _, s := range []*Vec2{{0, 0}, {1, 2}, {2, 1}, {-1, -1.5}} {
			m := Scale(s)
			want := Vec2{s.X * p.X, s.Y * p.Y}
			got := *m.TransformVec2(p)
			if got != want {
				t.Errorf("Got %v, want %v", got, want)
			}
		}
	}
}

func TestTranslation(t *testing.T) {
	for _, p := range []*Vec2{{0, 0}, {1, 2}, {2, 1}, {0, 0}, {4.1, 1.4}} {
		for _, v := range []*Vec2{{0, 0}, {1, 2}, {2, 1}, {-1, -1.5}} {
			m := Translation(v)
			want := Vec2{v.X + p.X, v.Y + p.Y}
			got := *m.TransformVec2(p)
			if got != want {
				t.Errorf("Got %v, want %v", got, want)
			}
		}
	}
}

func TestRotation(t *testing.T) {
	type tc struct {
		p     Vec2
		angle float64
		want  Vec2
	}
	for _, tt := range []tc{
		{Vec2{0, 0}, 0, Vec2{0, 0}},
		{Vec2{1, 2}, math.Pi / 2, Vec2{-2, 1}},
		{Vec2{2, 1}, math.Pi, Vec2{-2, -1}},
		{Vec2{2, 1}, -math.Pi / 2, Vec2{1, -2}},
		{Vec2{0, 0}, math.Pi / 2, Vec2{0, 0}},
		{Vec2{4.1, 1.4}, math.Pi / 2, Vec2{-1.4, 4.1}},
	} {
		m := Rotation(tt.angle)
		got := *m.TransformVec2(&tt.p)
		if Distance(&got, &tt.want) > epsilon {
			t.Errorf("Got %v, want %v when rotating %v by %v", got, tt.want, tt.p, tt.angle)
		}
	}
}

func snapTenth(f float64) float64 {
	return math.Round(f*10) / 10
}

func roundVec2(p *Vec2) *Vec2 {
	return &Vec2{snapTenth(p.X), snapTenth(p.Y)}
}

func roundRect(r *Rect) *Rect {
	return &Rect{*roundVec2(&r.Min), *roundVec2(&r.Max)}
}

type labeledAffine struct {
	label string
	a     *Affine
}

var transforms = []interface{}{
	labeledAffine{"identity", Identity()},
	labeledAffine{"translate(1, 2)", Translation(&Vec2{1, 2})},
	labeledAffine{"scale(3, 4)", Scale(&Vec2{3, 4})},
	labeledAffine{"rotate(-π/2)", Rotation(-math.Pi / 2)},
}

func TestCompose(t *testing.T) {
	for _, p := range []*Vec2{{1, 2}, {2, 1}, {0, 0}} {
		for affinePair := range cartesian.Iter(transforms, transforms) {
			first := affinePair[0].(labeledAffine)
			second := affinePair[1].(labeledAffine)
			want := second.a.TransformVec2(first.a.TransformVec2(p))
			got := Compose(first.a, second.a).TransformVec2(p)
			if Distance(got, want) > epsilon {
				t.Errorf(
					"Got %v, want %v with first %s (= %v) then %s",
					*roundVec2(got), *roundVec2(want),
					first.label, *roundVec2(first.a.TransformVec2(p)),
					second.label)
			}
		}
	}
}

func TestInverse(t *testing.T) {
	for _, p := range []*Vec2{{1, 2}, {2, 1}, {0, 0}} {
		for affinePair := range cartesian.Iter(transforms, transforms) {
			first := affinePair[0].(labeledAffine)
			second := affinePair[1].(labeledAffine)

			there := Compose(first.a, second.a)
			back := there.Inverse()
			q := there.TransformVec2(p)
			got := back.TransformVec2(q)
			if Distance(got, p) > epsilon {
				t.Errorf(
					"inverse of %s then %s gives %v, want %v",
					first.label, second.label, *roundVec2(got), *roundVec2(p))
			}
		}
	}
}

func TestPrepend(t *testing.T) {
	for _, p := range []*Vec2{{1, 2}, {2, 1}, {0, 0}} {
		for affinePair := range cartesian.Iter(transforms, transforms) {
			first := affinePair[0].(labeledAffine)
			second := affinePair[1].(labeledAffine)
			firstCopy := first.a.Copy()
			secondCopy := second.a.Copy()
			want := secondCopy.TransformVec2(firstCopy.TransformVec2(p))
			secondCopy.Prepend(firstCopy)
			got := secondCopy.TransformVec2(p)
			if Distance(got, want) > epsilon {
				t.Errorf(
					"Got %v, want %v with first %s (= %v) then %s",
					*roundVec2(got), *roundVec2(want),
					first.label, *roundVec2(first.a.TransformVec2(p)),
					second.label)
			}
		}
	}
}

func TestAppend(t *testing.T) {
	for _, p := range []*Vec2{{1, 2}, {2, 1}, {0, 0}} {
		for affinePair := range cartesian.Iter(transforms, transforms) {
			first := affinePair[0].(labeledAffine)
			second := affinePair[1].(labeledAffine)
			firstCopy := first.a.Copy()
			secondCopy := second.a.Copy()
			want := secondCopy.TransformVec2(firstCopy.TransformVec2(p))
			firstCopy.Append(secondCopy)
			got := firstCopy.TransformVec2(p)
			if Distance(got, want) > epsilon {
				t.Errorf(
					"Got %v, want %v with first %s (= %v) then %s",
					*roundVec2(got), *roundVec2(want),
					first.label, *roundVec2(first.a.TransformVec2(p)),
					second.label)
			}
		}
	}
}

func TestRotationAround(t *testing.T) {
	type tc struct {
		p    Vec2
		want *Rect
	}

	r := NewRect(0, 0, 2, 1)
	angle := -math.Pi / 2
	for _, tt := range []tc{
		{Vec2{0, 0}, NewRect(0, 0, 1, -2)},
		{Vec2{2, 1}, NewRect(1, 1, 2, 3)},
		{Vec2{1, .5}, NewRect(.5, -.5, 1.5, 1.5)},
	} {
		got := roundRect(RotationAround(angle, &tt.p).TransformRect(r))
		if *got != *tt.want {
			t.Errorf("Got %v, want %v when rotating around %v", got, tt.want, tt.p)
		}
	}
}

func TestNewRectFromCorners(t *testing.T) {
	type tc struct {
		a, b *Vec2
	}
	want := &Rect{Vec2{0, 0}, Vec2{1, 1}}
	for _, tt := range []tc{
		{&Vec2{0, 0}, &Vec2{1, 1}},
		{&Vec2{0, 1}, &Vec2{1, 0}},
		{&Vec2{1, 1}, &Vec2{0, 0}},
		{&Vec2{1, 0}, &Vec2{0, 1}},
	} {
		if got := NewRectFromCorners(tt.a, tt.b); *got != *want {
			t.Errorf("Got %v, want %v", got, want)
		}
	}
}

func TestNewRect(t *testing.T) {
	want := &Rect{Vec2{0, 0}, Vec2{1, 1}}
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

func TestRectIntersects(t *testing.T) {
	fred := NewRect(0, 0, 1, 1)
	type tc struct {
		george *Rect
		want   bool
	}
	for _, tt := range []tc{
		{NewRect(0, 0, 1, 1), true},
		{NewRect(-1, -1, 2, 2), true},
		{NewRect(.25, .25, .75, .75), true},
		{NewRect(0, 0, 1, 2), true},
		{NewRect(1, 1, 2, 2), true},
		{NewRect(1.000001, 1.00001, 2, 2), false},
		{NewRect(-1, -1, -.00001, -0.00001), false},
	} {
		if got := fred.Intersects(tt.george); got != tt.want {
			t.Errorf("got %v, want %v when intersecting %v and %v",
				got, tt.want, *fred, *tt.george)
		}
	}
}
