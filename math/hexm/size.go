package hexm

type Size interface {
	GetSize() (sizeX, sizeY, sizeZ int)
	Сontained(c Coord) bool
	IsEmpty() bool
}

type privSize struct {
	sizeX, sizeY, sizeZ int
}

func NewSize(sizeX, sizeY, sizeZ int) (Size, error) {

	if (sizeX <= 0) || (sizeY <= 0) || (sizeZ <= 0) {
		return nil, ErrorSizeZeroParameter
	}

	return &privSize{sizeX, sizeY, sizeZ}, nil
}

func (this *privSize) GetSize() (sizeX, sizeY, sizeZ int) {

	sizeX = this.sizeX
	sizeY = this.sizeY
	sizeZ = this.sizeZ

	return
}

func (this *privSize) IsEmpty() bool {
	return (this.sizeX == 0) || (this.sizeY == 0) || (this.sizeZ == 0)
}

func (this *privSize) Сontained(c Coord) bool {

	x, y, z := c.GetCoord()

	if (x < 0) || (x >= this.sizeX) {
		return false
	}
	if (y < 0) || (y >= this.sizeY) {
		return false
	}
	if (z < 0) || (z >= this.sizeZ) {
		return false
	}

	return true
}
