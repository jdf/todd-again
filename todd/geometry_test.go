package todd

import "testing"

func TestIdentity(t *testing.T) {
	m := Identity()
	for _, p := range []*Point{{0, 0}, {1, 2}, {2, 1}, {0, 0}, {4.1, 1.4}} {
		if *m.TransformPoint(p) != *p {
			t.Errorf("Got %v, expected %v", m.TransformPoint(p), p)
		}
	}
}
