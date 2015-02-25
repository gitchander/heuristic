package mindshift

import (
	"unicode/utf8"
)

type Puzzle struct {
	indexActive int
	primitives  []Primitive
}

func NewPuzzle(config PuzzleConfig) (*Puzzle, error) {

	ps, err := makePrimitives(config)
	if err != nil {
		return nil, err
	}

	return &Puzzle{
		indexActive: 0,
		primitives:  ps,
	}, nil
}

func makePrimitives(config PuzzleConfig) ([]Primitive, error) {

	pm := make(map[rune]*Primitive)

	var identifiers = config.Primitives.Identifiers
	for _, identifier := range identifiers {

		r, err := stringToRune(identifier.Id)
		if err != nil {
			return nil, err
		}

		start, err := intsToPoint(identifier.Location.Start)
		if err != nil {
			return nil, err
		}

		finish, err := intsToPoint(identifier.Location.Finish)
		if err != nil {
			return nil, err
		}

		pm[r] = &Primitive{
			Location: Location{
				Start:   start,
				Finish:  finish,
				Current: start,
			},
			Points: make([]Point, 0),
		}
	}

	var cs = config.Primitives.Cells
	for y, s := range cs {

		runeIndex := 0
		for x := 0; true; x++ {
			r, runeSize := utf8.DecodeRuneInString(s[runeIndex:])
			if runeSize == 0 {
				break
			}
			runeIndex += runeSize

			if p, ok := pm[r]; ok {
				p.Points = append(p.Points, Point{x, y})
			}
		}
	}

	var ps []Primitive
	for _, p := range pm {
		ps = append(ps, *p)
	}

	return ps, nil
}
