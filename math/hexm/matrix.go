package hexm

type Matrix struct {
	size    Coord
	vs      []interface{}
	indexXY int
	indexYZ int
	indexZX int
}

func NewMatrix(size Coord) *Matrix {

	var (
		sizeXY = size.X * (size.Y - 1)
		sizeYZ = size.Y * (size.Z - 1)
		sizeZX = size.Z * (size.X - 1)

		sizeXYZ = 1 + sizeXY + sizeYZ + sizeZX

		indexXY = 1
		indexYZ = indexXY + sizeXY
		indexZX = indexYZ + sizeYZ
	)

	return &Matrix{
		size:    size,
		vs:      make([]interface{}, sizeXYZ),
		indexXY: indexXY,
		indexYZ: indexYZ,
		indexZX: indexZX,
	}
}

func (m *Matrix) Size() Coord {
	return m.size
}

func (m *Matrix) coordToIndex(c Coord) (index int) {

	if !contained(m.size, c) {
		return -1
	}

	switch {

	case (c.X == 0 && c.Z > 0): // matrix yz
		index = m.indexYZ + ((c.Z-1)*m.size.Y + c.Y)

	case (c.Y == 0 && c.X > 0): // matrix zx
		index = m.indexZX + ((c.X-1)*m.size.Z + c.Z)

	case (c.Z == 0 && c.Y > 0): // matrix xy
		index = m.indexXY + ((c.Y-1)*m.size.X + c.X)

	default:
		index = 0
	}

	return
}

func (m *Matrix) indexToCoord(index int) Coord {

	var x, y, z int

	switch {

	case (index >= m.indexZX): // matrix zx
		{
			y = 0
			x, z = quoRem(index-m.indexZX, m.size.Z)
			x++
		}

	case (index >= m.indexYZ): // matrix yz
		{
			x = 0
			z, y = quoRem(index-m.indexYZ, m.size.Y)
			z++
		}

	case (index >= m.indexXY): // matrix xy
		{
			z = 0
			y, x = quoRem(index-m.indexXY, m.size.X)
			y++
		}

	default:
		x, y, z = 0, 0, 0
	}

	return Coord{x, y, z}
}

func (m *Matrix) Set(c Coord, v interface{}) (ok bool) {
	index := m.coordToIndex(c)
	if index != -1 {
		m.vs[index] = v
		ok = true
	}
	return
}

func (m *Matrix) Get(c Coord) (v interface{}, ok bool) {
	index := m.coordToIndex(c)
	if index != -1 {
		v = m.vs[index]
		ok = true
	}
	return
}
