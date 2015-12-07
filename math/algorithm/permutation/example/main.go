package main

import (
	"fmt"
	"log"

	"github.com/gitchander/heuristic/math/algorithm/permutation"
)

func main() {
	exampleUse()
	exampleTrace()
	exampleFactorial()
}

func exampleUse() {

	vs := []interface{}{true, -5, "str"}

	p, err := permutation.New(vs)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(vs)
	for p.Next() {
		fmt.Println(vs)
	}

	fmt.Println()
}

func exampleTrace() {

	vs := []interface{}{
		[]int{},
		[]bool{true, false},
		[]int{1, 2, 3},
		[]string{"one", "two", "three", "four"},
	}

	for _, v := range vs {
		traceValue(v)
		fmt.Println()
	}
}

func traceValue(v interface{}) {
	err := permutation.Trace(v,
		func(w interface{}) bool {
			fmt.Println(w)
			return true
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func exampleFactorial() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
}

// Calculates the factorial of a very long :)
func factorial(n int) int {
	a := make([]struct{}, n)
	i := 0
	err := permutation.Trace(a,
		func(_ interface{}) bool {
			i++
			return true
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	return i
}
