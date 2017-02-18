package hexm

type Vector struct {
	X, Y float64
}

func (a Vector) Add(b Vector) (c Vector) {
	c.X = a.X + b.X
	c.Y = a.Y + b.Y
	return
}

func (a Vector) Sub(b Vector) (c Vector) {
	c.X = a.X - b.X
	c.Y = a.Y - b.Y
	return
}

func (a Vector) MulScalar(t float64) (b Vector) {
	b.X = a.X * t
	b.Y = a.Y * t
	return
}

func (a Vector) DivScalar(t float64) (b Vector) {
	b.X = a.X / t
	b.Y = a.Y / t
	return
}

func VectorInPolygon(v Vector, poly []Vector) (ok bool) {
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
	return
}
