package hexm

type Coord struct {
	X, Y, Z int
}

func (c Coord) getError() error {

	if (c.X < 0) || (c.Y < 0) || (c.Z < 0) {
		return ErrorCoordNegativeParameter
	}

	if (c.X != 0) && (c.Y != 0) && (c.Z != 0) {
		return ErrorCoordOneZeroParameter
	}

	return nil
}

func (c *Coord) get_XYZ() (x, y, z int) {

	x = c.X
	y = c.Y
	z = c.Z

	return
}

func (c *Coord) set_XYZ(x, y, z int) {

	c.X = x
	c.Y = y
	c.Z = z

	return
}

func (c Coord) IsValid() bool {
	return (c.getError() == nil)
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

func (c Coord) IsZero() bool {
	return (c.X == 0) && (c.Y == 0) && (c.Z == 0)
}
