package main

import (
	"fmt"
	"github.com/gitchander/heuristic/math/graph2d"
)

func main() {

	v := graph2d.Vector{1, 0}

	t := graph2d.NewTransform()
	t.Rotate(graph2d.Pi / 4)
	t.Scale(2, 2)
	t.Move(5, 3)

	w := t.Apply(v)

	fmt.Println(w)
}
