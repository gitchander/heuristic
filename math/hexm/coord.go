package hexm

type NeighborDir int

const (
	_ NeighborDir = iota
	ND_POSITIVE
	ND_NEGATIVE
)

func (this NeighborDir) IsValid() bool {
	return (this == ND_POSITIVE) || (this == ND_NEGATIVE)
}

//---------------------------------------------------------------------------------
type Coord interface {
	GetCoord() (x, y, z int)
	IsNeighbor(other Coord) bool // return true if other is neighbor off this
	Equal(other Coord) bool      // return true if other equal this
	IsZero() bool
}

type privCoord struct {
	x, y, z int
}

func NewCoord(x, y, z int) (c Coord, err error) {

	if (x < 0) || (y < 0) || (z < 0) {
		err = ErrorCoordNegativeParameter
		return
	}

	if (x != 0) && (y != 0) && (z != 0) {
		err = ErrorCoordOneZeroParameter
		return
	}

	c = &privCoord{x, y, z}

	return
}

func (this *privCoord) GetCoord() (x, y, z int) {

	x = this.x
	y = this.y
	z = this.z

	return
}

func (this *privCoord) Equal(other Coord) bool {

	x, y, z := other.GetCoord()

	if this.x != x {
		return false
	}

	if this.y != y {
		return false
	}

	if this.z != z {
		return false
	}

	return true
}

func (this *privCoord) IsZero() bool {
	return (this.x == 0) && (this.y == 0) && (this.z == 0)
}

func (this *privCoord) NeighborX(nd NeighborDir) (Coord, error) {

	x, y, z := this.GetCoord()

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
		return nil, ErrorNeighborDirInvalid
	}

	return NewCoord(x, y, z)
}

func (this *privCoord) NeighborY(nd NeighborDir) (Coord, error) {

	x, y, z := this.GetCoord()

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
		return nil, ErrorNeighborDirInvalid
	}

	return NewCoord(x, y, z)
}

func (this *privCoord) NeighborZ(nd NeighborDir) (Coord, error) {

	x, y, z := this.GetCoord()

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
		return nil, ErrorNeighborDirInvalid
	}

	return NewCoord(x, y, z)
}

func (this *privCoord) IsNeighbor(other Coord) bool {

	var n Coord

	// X neighbors
	if n, _ = this.NeighborX(ND_POSITIVE); n.Equal(other) {
		return true
	}
	if n, _ = this.NeighborX(ND_NEGATIVE); n.Equal(other) {
		return true
	}

	// Y neighbors
	if n, _ = this.NeighborY(ND_POSITIVE); n.Equal(other) {
		return true
	}
	if n, _ = this.NeighborY(ND_NEGATIVE); n.Equal(other) {
		return true
	}

	// Z neighbors
	if n, _ = this.NeighborZ(ND_POSITIVE); n.Equal(other) {
		return true
	}
	if n, _ = this.NeighborZ(ND_NEGATIVE); n.Equal(other) {
		return true
	}

	return false
}
