package main

import (
	"fmt"

	"github.com/gitchander/heuristic/math/ternary/tri"
)

func TriBoolExample() {

	var a, b tri.Bool

	a = tri.StringToBool("True")
	fmt.Println(a)

	a = tri.False
	b = tri.True

	c := a.Or(b.Not())

	fmt.Println(c)

	fmt.Println(a.Xor(b))
	fmt.Println()
}

func TriBoolUnknown() {

	var a tri.Bool

	if a == tri.True {
		fmt.Println("true")
	} else if a == tri.False {
		fmt.Println("false")
	} else {
		fmt.Println("unknown")
	}

	if a == tri.Unknown {
		fmt.Println("unknown")
	} else {
		fmt.Println("not unknown")
	}
	fmt.Println()
}

func main() {
	TriBoolExample()
	TriBoolUnknown()
}
