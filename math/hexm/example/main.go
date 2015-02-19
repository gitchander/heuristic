package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gitchander/heuristic/math/hexm"
)

func main() {

	ExampleIterator()

	var (
		err error
		a   hexm.Coord
	)

	a, err = hexm.NewCoord(7, 0, 3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("cord sum=", sum(a.GetCoord()))

	fmt.Printf("%+v\n", a)
}

func sum(as ...int) int {

	s := 0
	for _, a := range as {
		s += a
	}
	return s
}

func ExampleIterator() error {

	s, err := hexm.NewSize(5, 5, 5)
	if err != nil {
		return err
	}

	m := hexm.NewMatrix(s)

	r := newRand()
	for I := m.NewIterator(); !I.Done(); I.Next() {

		I.SetCurrent(r.Intn(1000))
	}

	for I := m.NewIterator(); !I.Done(); I.Next() {
		c, v, _ := I.Current()
		fmt.Printf("coord: %v, val: %v\n", c, v)
	}

	var (
		cellIndex hexm.Coord
		c         interface{}
	)

	if cellIndex, err = hexm.NewCoord(0, 2, 4); err != nil {
		return err
	}

	err = m.SetCell(cellIndex, "3254324326")
	if err != nil {
		return err
	}

	if cellIndex, err = hexm.NewCoord(0, 2, 4); err != nil {
		return err
	}

	c, err = m.GetCell(cellIndex)
	fmt.Println("cell val: ", c)

	return nil
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
