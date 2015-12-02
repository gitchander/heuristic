package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gitchander/heuristic/math/hexm"
)

func main() {

	s := hexm.Size{5, 5, 5}

	m, err := hexm.NewMatrix(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for I := hexm.NewIterator(m); !I.Done(); I.Next() {

		I.SetCurrent(r.Intn(1000))
	}

	for I := hexm.NewIterator(m); !I.Done(); I.Next() {
		c, v, _ := I.Current()
		fmt.Printf("coord: %v, val: %v\n", c, v)
	}

	var (
		cellIndex hexm.Coord
		c         interface{}
	)

	cellIndex = hexm.Coord{0, 2, 4}
	if err != nil {
		log.Fatal(err.Error())
	}

	err = m.Set(cellIndex, "3254324326")
	if err != nil {
		log.Fatal(err.Error())
	}

	cellIndex = hexm.Coord{0, 2, 4}

	if !cellIndex.IsValid() {
		log.Fatal("is not valid")
	}

	c, err = m.Get(cellIndex)
	fmt.Println("cell val: ", c)
}
