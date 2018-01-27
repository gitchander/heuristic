package hexm

type NeighborDir int

const (
	_ NeighborDir = iota
	ND_POSITIVE
	ND_NEGATIVE
)

func (nd NeighborDir) IsValid() bool {
	return (nd == ND_POSITIVE) || (nd == ND_NEGATIVE)
}

func NeighborX(c Coord, nd NeighborDir) Coord {

	x, y, z := c.X, c.Y, c.Z

	switch nd {

	case ND_POSITIVE:
		{
			if y == 0 {
				x++
			} else {
				if z == 0 {
					x++
				} else {
					y--
					z--
				}
			}
		}

	case ND_NEGATIVE:
		{
			if x == 0 {
				y++
				z++
			} else {
				x--
			}
		}
	}

	return Coord{X: x, Y: y, Z: z}
}

func NeighborY(c Coord, nd NeighborDir) Coord {

	x, y, z := c.X, c.Y, c.Z

	switch nd {

	case ND_POSITIVE:
		{
			if z == 0 {
				y++
			} else {
				if x == 0 {
					y++
				} else {
					z--
					x--
				}
			}
		}

	case ND_NEGATIVE:
		{
			if y == 0 {
				z++
				x++
			} else {
				y--
			}
		}
	}

	return Coord{X: x, Y: y, Z: z}
}

func NeighborZ(c Coord, nd NeighborDir) Coord {

	x, y, z := c.X, c.Y, c.Z

	switch nd {

	case ND_POSITIVE:
		{
			if x == 0 {
				z++
			} else {
				if y == 0 {
					z++
				} else {
					x--
					y--
				}
			}
		}

	case ND_NEGATIVE:
		{
			if z == 0 {
				x++
				y++
			} else {
				z--
			}
		}
	}

	return Coord{X: x, Y: y, Z: z}
}

func IsNeighbors(a, b Coord) bool {

	var n Coord

	// X neighbors
	if n = NeighborX(a, ND_POSITIVE); n.Equal(b) {
		return true
	}
	if n = NeighborX(a, ND_NEGATIVE); n.Equal(b) {
		return true
	}

	// Y neighbors
	if n = NeighborY(a, ND_POSITIVE); n.Equal(b) {
		return true
	}
	if n = NeighborY(a, ND_NEGATIVE); n.Equal(b) {
		return true
	}

	// Z neighbors
	if n = NeighborZ(a, ND_POSITIVE); n.Equal(b) {
		return true
	}
	if n = NeighborZ(a, ND_NEGATIVE); n.Equal(b) {
		return true
	}

	return false
}
