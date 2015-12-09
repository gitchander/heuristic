package main

import (
	"fmt"

	"github.com/gitchander/heuristic/math/ternary"
)

func main() {
	ExampleSet()
	ExampleMul()
}

func ExampleSet() {
	x := ternary.NewTryte()
	x.SetInt(-13)
	fmt.Println(x)
	fmt.Println()
}

func ExampleMul() {

	a := ternary.NewTryteInt(250)
	b := ternary.NewTryteInt(-7)
	c := ternary.NewTryte()

	c.Mul(a, b)

	fmt.Println(c)
	fmt.Println(c.Int())
	fmt.Println()
}
