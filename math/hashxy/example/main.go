package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gitchander/heuristic/math/hashxy"
)

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomPoint(r *rand.Rand) hashxy.Point {

	const c = 10000
	return hashxy.Point{
		c - r.Intn(2*c),
		c - r.Intn(2*c),
	}
}

func main() {

	MatrixExample()
}

func MatrixExample() {

	m := hashxy.NewMatrix(11, 13)

	r := newRand()
	for i := 0; i < 100; i++ {
		m.Set(randomPoint(r), r.Intn(1000))
	}

	/*
		v := hm.GetCell(p)

		k, ok := v.(int)
		if ok {
			fmt.Println(k)
		}
	*/

	count := 0
	for I := m.NewIterator(); !I.Done(); I.Next() {

		p, v := I.Current()
		if k, ok := v.(int); ok {
			fmt.Println(p, k)
		}
		count++
	}
	fmt.Println(count)

	count = 0
	for I := m.NewIterator(); !I.Done(); I.Next() {
		count++
	}
	fmt.Println(count)
}
