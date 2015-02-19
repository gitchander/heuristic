package hexm

type Vector struct {
	X, Y float32
}

func (v Vector) Add(w Vector) Vector {
	return Vector{v.X + w.X, v.Y + w.Y}
}

func (v Vector) Sub(w Vector) Vector {
	return Vector{v.X - w.X, v.Y - w.Y}
}

func (v Vector) Mul(k float32) Vector {
	return Vector{v.X * k, v.Y * k}
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
