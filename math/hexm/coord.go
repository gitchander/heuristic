package hexm

type Coord interface {
	GetCoord() (x, y, z int)
}

type coordXYZ struct {
	x, y, z int
}

func NewCoord(x, y, z int) (Coord, error) {

	if (x < 0) || (y < 0) || (z < 0) {
		return nil, ErrorCoordNegativeParameter
	}

	if (x != 0) && (y != 0) && (z != 0) {
		return nil, ErrorCoordOneZeroParameter
	}

	return &coordXYZ{x, y, z}, nil
}

func (c *coordXYZ) GetCoord() (x, y, z int) {

	x = c.x
	y = c.y
	z = c.z

	return
}

func CoordEqual(a, b Coord) bool {

	aX, aY, aZ := a.GetCoord()
	bX, bY, bZ := b.GetCoord()

	if aX != bX {
		return false
	}

	if aY != bY {
		return false
	}

	if aZ != bZ {
		return false
	}

	return true
}

func CoordIsZero(c Coord) bool {

	x, y, z := c.GetCoord()

	return (x == 0) && (y == 0) && (z == 0)
}
