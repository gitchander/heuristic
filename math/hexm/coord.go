package hexm

type Coord struct {
	X, Y, Z int
}

func (c Coord) Norm() Coord {
	min := min3(c.X, c.Y, c.Z)
	if min == 0 {
		return c
	}
	return Coord{
		X: c.X - min,
		Y: c.Y - min,
		Z: c.Z - min,
	}
}

func (a Coord) Equal(b Coord) bool {

	if a.X != b.X {
		return false
	}

	if a.Y != b.Y {
		return false
	}

	if a.Z != b.Z {
		return false
	}

	return true
}

func contained(size, c Coord) bool {

	if (c.X < 0) || (c.X >= size.X) {
		return false
	}
	if (c.Y < 0) || (c.Y >= size.Y) {
		return false
	}
	if (c.Z < 0) || (c.Z >= size.Z) {
		return false
	}

	return true
}
