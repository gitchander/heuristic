package hexm

type Iterator struct {
	i int
	m *Matrix
}

func NewIterator(m *Matrix) *Iterator {
	return &Iterator{0, m}
}

func (p *Iterator) Done() bool {
	return (p.i >= len(p.m.vs))
}

func (p *Iterator) Current() (c Coord, v interface{}, err error) {

	vs := p.m.vs

	if p.i >= len(vs) {
		err = ErrorIteratorIndexOut
		return
	}

	if c, err = p.m.indexToCoord(p.i); err != nil {
		return
	}

	v = vs[p.i]

	return
}

func (p *Iterator) SetCurrent(v interface{}) error {

	vs := p.m.vs

	if p.i >= len(vs) {
		return ErrorIteratorIndexOut
	}

	vs[p.i] = v

	return nil
}

func (p *Iterator) Next() {
	p.i++
}
