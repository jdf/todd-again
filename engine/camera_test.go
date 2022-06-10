package engine

import (
	"testing"
)

type Vec2 = Vec2

var NewRect = NewRect

func TestNonInverted(t *testing.T) {
	// We imagine a world from -100, -100 to 100, 100.
	// We map it to a display from 0, 0 to 1, 1.
	cam := NewCamera(NewRect(-100, -100, 100, 100), NewRect(0, 0, 1, 1), DoNotFlipYAxis)

	got := cam.ToScreenVec2(&Vec2{0, 0})
	want := &Vec2{0.5, 0.5}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got = cam.ToScreenVec2(&Vec2{-100, -100})
	want = &Vec2{0, 0}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got = cam.ToScreenVec2(&Vec2{100, 100})
	want = &Vec2{1, 1}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInverted(t *testing.T) {
	// We imagine a world from -100, -100 to 100, 100.
	// We map it to a display from 0, 0 to 1, 1 with an inverted y-axis.
	cam := NewCamera(NewRect(-100, -100, 100, 100), NewRect(0, 0, 1, 1), FlipYAxis)

	got := cam.ToScreenVec2(&Vec2{0, 0})
	want := &Vec2{0.5, 0.5}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got = cam.ToScreenVec2(&Vec2{-100, -100})
	want = &Vec2{0, 1}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got = cam.ToScreenVec2(&Vec2{100, 100})
	want = &Vec2{1, 0}
	if !got.Equals(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
