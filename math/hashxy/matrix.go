package hashxy

import "container/list"

type Matrix interface {
	Set(p Point, newValue interface{}) (oldValue interface{})
	Get(p Point) interface{}
	Remove(p Point) interface{}
	NewIterator() Iterator
}

type matrix struct {
	width  int
	height int
	lists  []list.List
}

func NewMatrix(Width, Height int) Matrix {

	if (Width <= 0) || (Height <= 0) {
		return nil
	}

	return &matrix{
		width:  Width,
		height: Height,
		lists:  make([]list.List, Width*Height),
	}
}

func (this *matrix) getList(p Point) *list.List {

	x := mod(p.X, this.width)
	y := mod(p.Y, this.height)

	return &(this.lists[y*this.width+x])
}

func findElement(list *list.List, p Point) *list.Element {

	for e := list.Front(); e != nil; e = e.Next() {
		if pv, ok := e.Value.(*pointValue); ok {
			if pv.Point.Equal(p) {
				return e
			}
		}
	}

	return nil
}

func (this *matrix) Set(p Point, newValue interface{}) (oldValue interface{}) {

	list := this.getList(p)
	if e := findElement(list, p); e != nil {
		if pv := e.Value.(*pointValue); pv != nil {
			oldValue = pv.Value
			pv.Value = newValue
		}
	} else {
		pv := &pointValue{
			Point: p,
			Value: newValue,
		}
		list.PushBack(pv)
	}

	return
}

func (this *matrix) Get(p Point) interface{} {

	list := this.getList(p)
	if e := findElement(list, p); e != nil {
		if pv := e.Value.(*pointValue); pv != nil {
			return pv.Value
		}
	}

	return nil
}

func (this *matrix) Remove(p Point) interface{} {

	list := this.getList(p)
	if e := findElement(list, p); e != nil {
		pv := e.Value.(*pointValue)
		list.Remove(e)
		if pv != nil {
			return pv.Value
		}
	}

	return nil
}

func (this *matrix) NewIterator() Iterator {
	return newIterator(this.lists)
}
