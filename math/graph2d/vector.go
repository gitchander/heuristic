package graph2d

import "fmt"

type Vector struct {
	X, Y float32
}

func (v Vector) String() string {
	return fmt.Sprintf("(%g,%g)", v.X, v.Y)
}

func (v Vector) Add(w Vector) Vector {
	return Vector{v.X + w.X, v.Y + w.Y}
}

func (v Vector) Sub(w Vector) Vector {
	return Vector{v.X - w.X, v.Y - w.Y}
}

func (v Vector) MulScalar(k float32) Vector {
	return Vector{v.X * k, v.Y * k}
}

func (v Vector) Magnitude() float32 {
	x := v.X
	y := v.Y
	return Sqrt(x*x + y*y)
}

func (v Vector) Identity() Vector {

	m := v.Magnitude()
	if !Equal(m, 1) {
		inv_m := 1 / m
		return Vector{
			X: (v.X * inv_m),
			Y: (v.Y * inv_m),
		}
	}

	return v
}

func (v Vector) Distance(w Vector) float32 {
	x := v.X - w.X
	y := v.Y - w.Y
	return Sqrt(x*x + y*y)
}

func VectorInPolygon(v Vector, poly []Vector) bool {

	ok := false
	if n := len(poly); n > 0 {
		b := poly[n-1]
		for _, a := range poly {
			if (a.X <= v.X) && (b.X > v.X) || (b.X <= v.X) && (a.X > v.X) {
				if v.Y < (v.X-a.X)*(b.Y-a.Y)/(b.X-a.X)+a.Y {
					ok = !ok
				}
			}
			b = a
		}
	}
	return ok
}
