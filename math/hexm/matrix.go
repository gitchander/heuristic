package hexm

type Matrix struct {
	s       Size
	vs      []interface{}
	indexXY int
	indexYZ int
	indexZX int
}

func NewMatrix(s Size) (*Matrix, error) {

	if err := s.getError(); err != nil {
		return nil, err
	}

	var (
		sizeXY = s.Dx * (s.Dy - 1)
		sizeYZ = s.Dy * (s.Dz - 1)
		sizeZX = s.Dz * (s.Dx - 1)

		sizeXYZ = 1 + sizeXY + sizeYZ + sizeZX

		indexXY = 1
		indexYZ = indexXY + sizeXY
		indexZX = indexYZ + sizeYZ
	)

	return &Matrix{
		s:       s,
		vs:      make([]interface{}, sizeXYZ),
		indexXY: indexXY,
		indexYZ: indexYZ,
		indexZX: indexZX,
	}, nil
}

func (m *Matrix) Size() Size {
	return m.s
}

func (m *Matrix) coordToIndex(c Coord) (index int, err error) {

	if !m.s.Сontained(c) {
		err = ErrorSizeOutOfRange
		return
	}

	switch {

	case (c.X == 0 && c.Z > 0): // matrix yz
		index = m.indexYZ + ((c.Z-1)*m.s.Dy + c.Y)

	case (c.Y == 0 && c.X > 0): // matrix zx
		index = m.indexZX + ((c.X-1)*m.s.Dz + c.Z)

	case (c.Z == 0 && c.Y > 0): // matrix xy
		index = m.indexXY + ((c.Y-1)*m.s.Dx + c.X)

	default:
		index = 0
	}

	return
}

// quo = x / y
// rem = x % y
func quoRem(x, y int) (quo, rem int) {

	quo = x / y
	rem = x - quo*y

	return
}

func (m *Matrix) indexToCoord(index int) (c Coord, err error) {

	if !m.indexIsValid(index) {
		err = ErrorIndexIsNotValit
		return
	}

	var x, y, z int

	switch {

	case (index >= m.indexZX): // matrix zx
		{
			y = 0
			x, z = quoRem(index-m.indexZX, m.s.Dz)
			x++
		}

	case (index >= m.indexYZ): // matrix yz
		{
			x = 0
			z, y = quoRem(index-m.indexYZ, m.s.Dy)
			z++
		}

	case (index >= m.indexXY): // matrix xy
		{
			z = 0
			y, x = quoRem(index-m.indexXY, m.s.Dx)
			y++
		}

	default:
		x, y, z = 0, 0, 0
	}

	c.set_XYZ(x, y, z)
	err = c.getError()

	return
}

func (m *Matrix) indexIsValid(index int) bool {
	return (index >= 0) && (index < len(m.vs))
}

func (m *Matrix) Set(c Coord, v interface{}) error {

	err := c.getError()
	if err != nil {
		return err
	}

	index, err := m.coordToIndex(c)
	if err != nil {
		return err
	}

	m.vs[index] = v

	return nil
}

func (m *Matrix) Get(c Coord) (interface{}, error) {

	err := c.getError()
	if err != nil {
		return nil, err
	}

	index, err := m.coordToIndex(c)
	if err != nil {
		return nil, err
	}

	v := m.vs[index]

	return v, nil
}
