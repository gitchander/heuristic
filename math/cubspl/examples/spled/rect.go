package main

import (
	"github.com/gitchander/heuristic/math/cubspl"
)

type Rectangle struct {
	Min cubspl.Point
	Max cubspl.Point
}

func (r *Rectangle) IsValid() bool {

	if r.Min.X > r.Max.X {
		return false
	}

	if r.Min.Y > r.Max.Y {
		return false
	}

	return true
}

func (r *Rectangle) Width() float64 {
	return r.Max.X - r.Min.X
}

func (r *Rectangle) Height() float64 {
	return r.Max.Y - r.Min.Y
}

func (r *Rectangle) NormX(x float64) float64 {

	if x < r.Min.X {
		x = r.Min.X
	}

	if x > r.Max.X {
		x = r.Max.X
	}

	return x
}

func (r *Rectangle) NormY(y float64) float64 {

	if y < r.Min.Y {
		y = r.Min.Y
	}

	if y > r.Max.Y {
		y = r.Max.Y
	}

	return y
}

func (r *Rectangle) NormPoint(p cubspl.Point) cubspl.Point {

	p.X = r.NormX(p.X)
	p.Y = r.NormY(p.Y)

	return p
}
