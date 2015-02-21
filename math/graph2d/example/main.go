package main

import (
	"fmt"
	. "github.com/gitchander/heuristic/math/graph2d"
)

func main() {

	v := Vector{3, 5}

	t := Identity()
	t.Rotate(Pi / 4)
	t.Scale(2, 2)
	t.Move(5, 3)

	w := v.Transform(t)

	fmt.Println(w)
}
