package main

import (
	"fmt"

	"github.com/gitchander/heuristic/math/ternary"
)

func main() {
	exampleCounter()
	exampleSet()
	exampleMul()
}

func exampleCounter() {
	for i := 0; i < 122; i++ {
		n := ternary.NewTryteInt(-i)
		p := ternary.NewTryteInt(i)
		fmt.Printf("%4d: %s\n", -i, n)
		fmt.Printf("%4d: %s\n", i, p)
		fmt.Println()
	}
}

func exampleSet() {
	x := ternary.NewTryte()
	x.SetInt(-13)
	fmt.Println(x)
	fmt.Println()
}

func exampleMul() {

	a := ternary.NewTryteInt(250)
	b := ternary.NewTryteInt(-7)
	c := ternary.NewTryte()

	c.Mul(a, b)

	fmt.Println(c)
	fmt.Println(c.Int())
	fmt.Println()
}
