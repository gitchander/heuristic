package hexm

type Coord interface {
	GetCoord() (x, y, z int)	
	Equal(other Coord) bool // return true if other equal this
	IsZero() bool
}

type privCoord struct {
	x, y, z int
}

func NewCoord(x, y, z int) (Coord, error) {

	if (x < 0) || (y < 0) || (z < 0) {
		return nil, ErrorCoordNegativeParameter
	}

	if (x != 0) && (y != 0) && (z != 0) {
		return nil, ErrorCoordOneZeroParameter
	}

	return &privCoord{x, y, z}, nil
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
