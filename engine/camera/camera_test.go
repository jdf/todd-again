package camera

import (
	"testing"

	"github.com/jdf/todd-again/engine/geometry"
)

type Vec2 = geometry.Vec2

var NewRect = geometry.NewRect

func TestNonInverted(t *testing.T) {
	// We imagine a world from -100, -100 to 100, 100.
	// We map it to a display from 0, 0 to 1, 1.
	v := New(NewRect(-100, -100, 100, 100), NewRect(0, 0, 1, 1), DoNotFlip)

	got := v.ToScreenVec2(&Vec2{0, 0})
	want := &Vec2{0.5, 0.5}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got = v.ToScreenVec2(&Vec2{-100, -100})
	want = &Vec2{0, 0}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got = v.ToScreenVec2(&Vec2{100, 100})
	want = &Vec2{1, 1}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInverted(t *testing.T) {
	// We imagine a world from -100, -100 to 100, 100.
	// We map it to a display from 0, 0 to 1, 1 with an inverted y-axis.
	v := New(NewRect(-100, -100, 100, 100), NewRect(0, 0, 1, 1), Flip)

	got := v.ToScreenVec2(&Vec2{0, 0})
	want := &Vec2{0.5, 0.5}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got = v.ToScreenVec2(&Vec2{-100, -100})
	want = &Vec2{0, 1}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got = v.ToScreenVec2(&Vec2{100, 100})
	want = &Vec2{1, 0}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
