package main

import (
	"fmt"
	"time"

	perm "github.com/gitchander/permutation"
)

func Example1() {

	a := []interface{}{"A", "B", "C"}
	for p := perm.New(a); p.Valid(); p.Next() {
		fmt.Println(a)
	}
}

func Example2() {

	a := []interface{}{"this is string", 17, -0.56, "qwertyuiop"}
	for p := perm.New(a); p.Valid(); p.Next() {
		fmt.Println(a)
	}
}

func Example3() {

	begin := time.Now()
	a := []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
	for p := perm.New(a); p.Valid(); p.Next() {
	}
	fmt.Println(time.Since(begin))
}

func main() {

	Example1()
	Example2()
	Example3()
}
