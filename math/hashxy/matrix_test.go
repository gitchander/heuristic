package hashxy

import (
	"testing"
)

func TestMatrixFill(t *testing.T) {

	nx, ny := 19, 17
	m := NewMatrix(nx, ny)

	for y := 0; y < ny; y++ {
		for x := 0; x < nx; x++ {
			m.Set(Point{x, y}, nil)
			m.Set(Point{x, y + ny}, nil)
			m.Set(Point{x + nx, y}, nil)
			m.Set(Point{x + nx, y + ny}, nil)
		}
	}

	M := m.(*matrix)
	for _, list := range M.lists {
		n := list.Len()
		if n != 4 {
			t.Error("list wrong len")
			return
		}
	}
}
