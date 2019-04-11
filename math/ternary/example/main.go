package main

import (
	"fmt"

	ter "github.com/gitchander/heuristic/math/ternary"
)

func main() {
	exampleSet()
	exampleCounter()
	exampleMul()
}

func exampleSet() {
	x := ter.Ts9Int(-13)
	fmt.Println(x)
}

func exampleCounter() {
	for i := 0; i < 100; i++ {
		n := ter.Ts9Int(-i)
		p := ter.Ts9Int(i)
		fmt.Printf("%4d: %s\n", -i, n)
		fmt.Printf("%4d: %s\n", i, p)
		fmt.Println()
	}
}

func exampleMul() {

	a := ter.Ts9Int(250)
	b := ter.Ts9Int(-7)

	c := a.Mul(b)

	fmt.Println(c)
	fmt.Println(c.Int())
	fmt.Println()
}
