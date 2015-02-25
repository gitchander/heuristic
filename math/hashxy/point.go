package hashxy

type Point struct {
	X, Y int
}

func (p Point) Equal(q Point) bool {

	if p.X != q.X {
		return false
	}

	if p.Y != q.Y {
		return false
	}

	return true
}

type pointValue struct {
	Point Point
	Value interface{}
}
