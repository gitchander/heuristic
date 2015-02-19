package graph2d

type Vector struct {
	X, Y float32
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

func (v Vector) MulMatrix(m Matrix) Vector {

	return Vector{}
}

func (v Vector) rotate(angle float32) Vector {
	sin, cos := SinCos(angle)
	return Vector{
		X: (v.X*cos - v.Y*sin),
		Y: (v.Y*cos + v.X*sin),
	}
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
