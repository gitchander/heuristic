package main

import (
	"math"

	"github.com/gitchander/heuristic/math/cubspl"
)

func vectorsNear(a, b cubspl.Point) bool {

	const e = 4.0

	if math.Abs(a.X-b.X) > e {
		return false
	}

	if math.Abs(a.Y-b.Y) > e {
		return false
	}

	return true
}

func insertByIndex(ps []cubspl.Point, j int, p cubspl.Point) []cubspl.Point {

	n := len(ps)
	if cap(ps) == n {
		temp := make([]cubspl.Point, n+1, (n+1)*2)
		copy(temp[:n], ps[:n])
		ps = temp
	}

	ps = ps[:n+1]

	for i := n; i > j; i-- {
		ps[i] = ps[i-1]
	}

	ps[j] = p

	return ps
}

func removeByIndex(ps []cubspl.Point, j int) (cubspl.Point, []cubspl.Point) {

	p := ps[j]

	n := len(ps)
	for i := j; i < n-1; i++ {
		ps[i] = ps[i+1]
	}

	ps = ps[:n-1]

	return p, ps
}
