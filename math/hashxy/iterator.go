package hashxy

import "container/list"

type Iterator interface {
	Done() bool
	Next()
	Current() (p Point, v interface{})
	SetCurrent(v interface{})
}

type iterator struct {
	index   int
	lists   []list.List
	element *list.Element
}

func newIterator(ls []list.List) (I *iterator) {

	I = &iterator{lists: ls}

	for i, list := range ls {
		if e := list.Front(); e != nil {
			I.index = i
			I.element = e
			break
		}
	}

	return I
}

func (this *iterator) Current() (p Point, v interface{}) {

	if this.element == nil {
		return
	}

	pv, ok := this.element.Value.(*pointValue)
	if !ok {
		return
	}

	p = pv.Point
	v = pv.Value

	return
}

func (this *iterator) SetCurrent(v interface{}) {

	if this.element == nil {
		return
	}

	pv, ok := this.element.Value.(*pointValue)
	if !ok {
		return
	}

	if v != nil {
		pv.Value = v
	}
}

func (this *iterator) Next() {

	if this.element == nil {
		return
	}

	this.element = this.element.Next()
	if this.element != nil {
		return
	}

	ls := this.lists
	for i := this.index + 1; i < len(ls); i++ {
		if e := ls[i].Front(); e != nil {
			this.element = e
			this.index = i
			break
		}
	}
}

func (this *iterator) Done() bool {
	return (this.element == nil)
}
