package hexm

type Iterator interface {
	Done() bool
	Current() (c Coord, v interface{}, err error)
	SetCurrent(v interface{}) error
	Next()
}

type privIterator struct {
	index        int
	cells        []interface{}
	indexToCoord func(index int) (Coord, error)
}

func (this *privIterator) Done() bool {
	return this.index >= len(this.cells)
}

func (this *privIterator) Current() (c Coord, v interface{}, err error) {

	if this.index >= len(this.cells) {
		err = ErrorIteratorIndexOut
		return
	}

	if c, err = this.indexToCoord(this.index); err != nil {
		return
	}

	v = this.cells[this.index]

	return
}

func (this *privIterator) SetCurrent(v interface{}) error {

	if this.index >= len(this.cells) {
		return ErrorIteratorIndexOut
	}

	this.cells[this.index] = v

	return nil
}

func (this *privIterator) Next() {
	this.index++
}
