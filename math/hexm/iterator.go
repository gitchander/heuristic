package hexm

type Iterator struct {
	index int
	m     *Matrix
}

func NewIterator(m *Matrix) *Iterator {
	return &Iterator{0, m}
}

func (p *Iterator) HasValue() bool {
	return (p.index < len(p.m.vs))
}

func (p *Iterator) GetValue() interface{} {
	return p.m.vs[p.index]
}

func (p *Iterator) SetValue(v interface{}) {
	p.m.vs[p.index] = v
}

func (p *Iterator) Coord() Coord {
	return p.m.indexToCoord(p.index)
}

func (p *Iterator) Next() {
	p.index++
}
