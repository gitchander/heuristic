package hexm

/*

//-----------------------
matrix size = 3 x 3 x 3
coord= (x, y, z)
//-----------------------

				(0,2,0)

		(0,2,1)			(1,2,0)

(0,2,2)			(0,1,0)			(2,2,0)

		(0,1,1)			(1,1,0)

(0,1,2)			(0,0,0)			(2,1,0)

		(0,0,1)			(1,0,0)

(0,0,2)			(1,0,1)			(2,0,0)

		(1,0,2)			(2,0,1)

				(2,0,2)

*/

type Matrix interface {
	SetCell(c Coord, cell interface{}) (err error)
	GetCell(c Coord) (cell interface{}, err error)
	NewIterator() Iterator
}

type privMatrix struct {
	size    Size
	cells   []interface{}
	indexXY int
	indexYZ int
	indexZX int
}

func NewMatrix(s Size) Matrix {

	if s == nil {
		return nil
	}

	sizeX, sizeY, sizeZ := s.GetSize()

	var size_xy = sizeX * (sizeY - 1)
	var size_yz = sizeY * (sizeZ - 1)
	var size_zx = sizeZ * (sizeX - 1)

	var size_xyz = 1 + size_xy + size_yz + size_zx

	var index_xy = 1
	var index_yz = index_xy + size_xy
	var index_zx = index_yz + size_yz

	return &privMatrix{
		size:    s,
		cells:   make([]interface{}, size_xyz),
		indexXY: index_xy,
		indexYZ: index_yz,
		indexZX: index_zx,
	}
}

func (this *privMatrix) GetSize() Size {

	if this == nil {
		return nil
	}

	return this.size
}

func (this *privMatrix) IsEmpty() bool {

	if this == nil {
		return true
	}

	return this.IsEmpty()
}

func (this *privMatrix) coordToIndex(c Coord) (index int, err error) {

	x, y, z := c.GetCoord()

	if !this.size.Сontained(c) {
		err = ErrorSizeOutOfRange
		return
	}

	sizeX, sizeY, sizeZ := this.size.GetSize()

	switch {

	case (x == 0 && z > 0): // matrix yz
		index = this.indexYZ + ((z-1)*sizeY + y)

	case (y == 0 && x > 0): // matrix zx
		index = this.indexZX + ((x-1)*sizeZ + z)

	case (z == 0 && y > 0): // matrix xy
		index = this.indexXY + ((y-1)*sizeX + x)

	default:
		index = 0
	}

	return
}

func (this *privMatrix) indexToCoord(index int) (Coord, error) {

	if !this.indexIsValid(index) {
		err := ErrorIndexIsNotValit
		return nil, err
	}

	var x, y, z int
	sizeX, sizeY, sizeZ := this.size.GetSize()

	switch {

	case (index >= this.indexZX): // matrix zx
		{
			y = 0
			x, z = divmod(index-this.indexZX, sizeZ)
			x++
		}

	case (index >= this.indexYZ): // matrix yz
		{
			x = 0
			z, y = divmod(index-this.indexYZ, sizeY)
			z++
		}

	case (index >= this.indexXY): // matrix xy
		{
			z = 0
			y, x = divmod(index-this.indexXY, sizeX)
			y++
		}

	default:
		x, y, z = 0, 0, 0
	}

	return NewCoord(x, y, z)
}

func (this *privMatrix) indexIsValid(index int) bool {
	return (index >= 0) && (index < len(this.cells))
}

func (this *privMatrix) SetCell(c Coord, cell interface{}) (err error) {

	var index int
	if index, err = this.coordToIndex(c); err != nil {
		return
	}

	this.cells[index] = cell

	return
}

func (this *privMatrix) GetCell(c Coord) (cell interface{}, err error) {

	var index int
	if index, err = this.coordToIndex(c); err != nil {
		return
	}

	cell = this.cells[index]

	return
}

func (this *privMatrix) NewIterator() Iterator {

	if this == nil {
		return nil
	}

	return &privIterator{
		index:        0,
		cells:        this.cells,
		indexToCoord: this.indexToCoord,
	}
}
