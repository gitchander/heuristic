package cubspl

import (
	"testing"
)

func TestSpline(t *testing.T) {

	ps := []Point{
		Point{0, 0},
		Point{1, 1},
	}

	spline, err := NewCubicSpline(ps)
	if err != nil {
		t.Error(err)
		return
	}

	y := spline.Interpolate(0.5)
	if !equal(y, 0.5) {
		t.Error("wrong Y")
		return
	}
}
