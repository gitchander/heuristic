package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gitchander/heuristic/math/hashxy"
)

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randPoint(r *rand.Rand) hashxy.Point {

	const c = 10000
	return hashxy.Point{
		c - r.Intn(2*c),
		c - r.Intn(2*c),
	}
}

func main() {

	m, err := hashxy.NewMatrix(11, 13)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := newRand()
	for i := 0; i < 100; i++ {
		m.Set(randPoint(r), r.Intn(1000))
	}

	count := 0
	for I := hashxy.NewIterator(m); !I.Done(); I.Next() {

		p, v := I.Current()
		if k, ok := v.(int); ok {
			fmt.Println(p, k)
		}
		count++
	}
	fmt.Println(count)

	count = 0
	for I := hashxy.NewIterator(m); !I.Done(); I.Next() {
		count++
	}
	fmt.Println(count)
}
