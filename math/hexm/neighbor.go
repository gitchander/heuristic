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

func NeighborX(c Coord, nd NeighborDir) (n Coord, err error) {

	if err = c.getError(); err != nil {
		return
	}

	x, y, z := c.get_XYZ()

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

	default:
		err = ErrorNeighborDirInvalid
		return
	}

	n.set_XYZ(x, y, z)

	return
}

func NeighborY(c Coord, nd NeighborDir) (n Coord, err error) {

	if err = c.getError(); err != nil {
		return
	}

	x, y, z := c.get_XYZ()

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

	default:
		err = ErrorNeighborDirInvalid
		return
	}

	n.set_XYZ(x, y, z)

	return
}

func NeighborZ(c Coord, nd NeighborDir) (n Coord, err error) {

	if err = c.getError(); err != nil {
		return
	}

	x, y, z := c.get_XYZ()

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

	default:
		err = ErrorNeighborDirInvalid
		return
	}

	n.set_XYZ(x, y, z)

	return
}

func IsNeighbors(a, b Coord) bool {

	var n Coord

	// X neighbors
	if n, _ = NeighborX(a, ND_POSITIVE); n.Equal(b) {
		return true
	}
	if n, _ = NeighborX(a, ND_NEGATIVE); n.Equal(b) {
		return true
	}

	// Y neighbors
	if n, _ = NeighborY(a, ND_POSITIVE); n.Equal(b) {
		return true
	}
	if n, _ = NeighborY(a, ND_NEGATIVE); n.Equal(b) {
		return true
	}

	// Z neighbors
	if n, _ = NeighborZ(a, ND_POSITIVE); n.Equal(b) {
		return true
	}
	if n, _ = NeighborZ(a, ND_NEGATIVE); n.Equal(b) {
		return true
	}

	return false
}
