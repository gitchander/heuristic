package main

import (
	"fmt"

	. "github.com/gitchander/heuristic/math/ternary"
)

func ExampleBoolean() {

	var a, b TriBool

	a = False
	b = True

	c := a.Or(b.Not())
	fmt.Println(c)
	fmt.Println(b == True)

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

	ExampleTryte()
	//ExampleBoolean()
}
