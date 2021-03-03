package main

import (
	"fmt"
	"log"

	"github.com/gitchander/heuristic/math/ternary/tri"
)

func TriExample() {

	a, err := tri.ParseTri("True")
	checkError(err)

	fmt.Println(a)

	a = tri.False
	b := tri.True

	c := a.Or(b.Not())

	fmt.Println(c)

	fmt.Println(a.Xor(b))
	fmt.Println()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TriUnknown() {

	var a tri.Tri

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
	TriExample()
	TriUnknown()
}
