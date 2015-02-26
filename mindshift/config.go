package mindshift

import (
	"unicode/utf8"
)

type Identifier struct {
	Id       string
	Location struct {
		Start  []int
		Finish []int
	}
}

type PuzzleConfig struct {
	Name        string
	Author      string
	Description string
	Difficulty  int
	Primitives  struct {
		Cells       []string
		Identifiers []Identifier
	}
}

func intsToPoint(ds []int) (p Point, err error) {

	if len(ds) != 2 {
		err = newError("intsToPoint")
		return
	}

	p = Point{ds[0], ds[1]}

	return
}

func stringToRune(s string) rune {

	r, runeSize := utf8.DecodeRuneInString(s)
	if runeSize != len(s) {
		return utf8.RuneError
	}

	return r
}
