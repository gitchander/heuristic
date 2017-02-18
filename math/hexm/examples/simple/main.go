package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gitchander/heuristic/math/hexm"
)

func main() {

	m := hexm.NewMatrix(hexm.Coord{5, 5, 5})
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for I := hexm.NewIterator(m); I.HasValue(); I.Next() {
		I.SetValue(r.Intn(1000))
	}

	for I := hexm.NewIterator(m); I.HasValue(); I.Next() {
		var (
			v = I.GetValue()
			c = I.Coord()
		)
		fmt.Printf("coord: %v, val: %v\n", c, v)
	}

	var (
		cellIndex hexm.Coord
		c         interface{}
	)

	cellIndex = hexm.Coord{0, 2, 4}

	m.Set(cellIndex, "3254324326")
	cellIndex = hexm.Coord{0, 2, 4}

	c, ok := m.Get(cellIndex)
	fmt.Println(c, ok)
}
