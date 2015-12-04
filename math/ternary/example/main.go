package main

import (
	"fmt"

	"github.com/gitchander/heuristic/math/ternary"
)

func TriBoolUse() {

	var a ternary.Bool

	if a == ternary.True {
		fmt.Println("\t true")
	} else if a == ternary.False {
		fmt.Println("\t false")
	} else {
		fmt.Println("\t unknown")
	}

	if a == ternary.Unknown {
		fmt.Println("\t unknown")
	} else {
		fmt.Println("\t not unknown")
	}
}

func TriBoolExample() {

	var a, b ternary.Bool

	if a.Parse("True") {
		fmt.Println(a)
	}

	a = ternary.False
	b = ternary.True

	c := a.Or(b.Not())

	fmt.Println(c)
	fmt.Println(b == a)

	fmt.Println(a.Xor(b))
}

func ExampleTryte() {

	a := ternary.NewTryte(-15)
	b := ternary.NewTryte(-7)
	c := new(ternary.Tryte)

	c.Mul(a, b)

	fmt.Println(c)
	fmt.Println(c.Int())
}

func main() {

	TriBoolUse()
}
