package main

import (
	"fmt"

	"github.com/gitchander/heuristic/math/ternary"
)

func ExampleTryte() {

	a := ternary.NewTryte(-15)
	b := ternary.NewTryte(-7)
	c := new(ternary.Tryte)

	c.Mul(a, b)

	fmt.Println(c)
	fmt.Println(c.Int())
}

func main() {
	ExampleTryte()
}
