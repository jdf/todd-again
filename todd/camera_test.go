package todd

import "testing"

func TestFullViewport(t *testing.T) {
	// We imagine a world from -100, -100 to 100, 100.
	// We map it to a display from 0, 0 to 1, 1.
	v := NewView(NewRect(-100, -100, 100, 100))

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
