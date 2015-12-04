package main

import (
	"fmt"

	. "github.com/gitchander/heuristic/math/ternary"
)

func TriBoolUse() {

	var a Trilean

	if a == True {
		fmt.Println("\t true")
	} else if a == False {
		fmt.Println("\t false")
	} else {
		fmt.Println("\t unknown")
	}

	if a == Unknown {
		fmt.Println("\t unknown")
	} else {
		fmt.Println("\t not unknown")
	}
}

func TriBoolExample() {

	var a, b Trilean

	if a.Parse("True") {
		fmt.Println(a)
	}

	a = False
	b = True

	c := a.Or(b.Not())

	fmt.Println(c)
	fmt.Println(b == a)

	fmt.Println(a.Xor(b))
}

func ExampleTryte() {

	a := NewTryte(-15)
	b := NewTryte(-7)
	c := new(Tryte)

	c.Mul(a, b)

	fmt.Println(c)
	fmt.Println(c.Int())
}

func main() {

	TriBoolUse()
}
