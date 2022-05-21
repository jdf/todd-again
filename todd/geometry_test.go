package todd

import (
	"math"
	"testing"
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
