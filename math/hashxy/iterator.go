package hashxy

import "container/list"

type Iterator struct {
	index   int
	lists   []list.List
	element *list.Element
}

func NewIterator(m *Matrix) *Iterator {
	return newIterator(m.lists)
}

func newIterator(ls []list.List) *Iterator {

	it := &Iterator{lists: ls}

	for i, list := range ls {
		if e := list.Front(); e != nil {
			it.index = i
			it.element = e
			break
		}
	}

	return it
}

func (it *Iterator) Current() (p Point, v interface{}) {

	if it.element == nil {
		return
	}

	pv, ok := it.element.Value.(*pointValue)
	if !ok {
		return
	}

	p = pv.Point
	v = pv.Value

	return
}

func (it *Iterator) SetCurrent(v interface{}) {

	if it.element == nil {
		return
	}

	pv, ok := it.element.Value.(*pointValue)
	if !ok {
		return
	}

	if v != nil {
		pv.Value = v
	}
}

func (it *Iterator) Next() {

	if it.element == nil {
		return
	}

	it.element = it.element.Next()
	if it.element != nil {
		return
	}

	ls := it.lists
	for i := it.index + 1; i < len(ls); i++ {
		if e := ls[i].Front(); e != nil {
			it.element = e
			it.index = i
			break
		}
	}
}

func (it *Iterator) Done() bool {
	return (it.element == nil)
}
