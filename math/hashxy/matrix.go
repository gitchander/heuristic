package hashxy

import (
	"container/list"
	"errors"
)

type Matrix struct {
	width  int
	height int
	lists  []list.List
}

func NewMatrix(Width, Height int) (*Matrix, error) {

	if (Width <= 0) || (Height <= 0) {
		return nil, errors.New("value Width or Height <= 0")
	}

	return &Matrix{
		width:  Width,
		height: Height,
		lists:  make([]list.List, Width*Height),
	}, nil
}

func (m *Matrix) getList(p Point) *list.List {

	x := mod(p.X, m.width)
	y := mod(p.Y, m.height)

	return &(m.lists[y*m.width+x])
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

func (m *Matrix) Set(p Point, newValue interface{}) (oldValue interface{}) {

	list := m.getList(p)
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

func (m *Matrix) Get(p Point) interface{} {

	list := m.getList(p)
	if e := findElement(list, p); e != nil {
		if pv := e.Value.(*pointValue); pv != nil {
			return pv.Value
		}
	}

	return nil
}

func (m *Matrix) Remove(p Point) interface{} {

	list := m.getList(p)
	if e := findElement(list, p); e != nil {
		pv := e.Value.(*pointValue)
		list.Remove(e)
		if pv != nil {
			return pv.Value
		}
	}

	return nil
}
