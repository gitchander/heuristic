package mindshift

import (
	"encoding/json"
	"errors"
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

func cellToPoint(ds []int) (p Point, err error) {

	if len(ds) != 2 {
		err = errors.New("cellToPoint")
		return
	}

	p = Point{ds[0], ds[1]}

	return
}

func makePrimitives(config PuzzleConfig) ([]Primitive, error) {

	var err error
	var startPos, finishPos Point

	var Identifiers = config.Primitives.Identifiers
	var ps = make([]Primitive, len(Identifiers))

	for i, node := range Identifiers {

		if startPos, err = cellToPoint(node.Location.Start); err != nil {
			return nil, err
		}

		if finishPos, err = cellToPoint(node.Location.Finish); err != nil {
			return nil, err
		}

		ps[i] = Primitive{
			//Id: (node.Id),
			//Cells: (new Array()),
			Location: Location{
				Start:   startPos,
				Current: startPos,
				Finish:  finishPos,
			},
		}
	}

	var strArray = config.Primitives.Cells

	for y, str := range strArray {
		for x, id := range str {

			for i := 0; i < len(Identifiers); i++ {

				var prim = ps[i]

				if id == rune(Identifiers[i].Id) {

					prim.Points = append(prim.Points, Point{x, y})
					break
				}
			}
		}
	}

	return ps, nil
}

type Id rune

func (this *Id) MarshalJSON() ([]byte, error) {

	s := string(*this)

	return json.Marshal(&s)
}

func (this *Id) UnmarshalJSON(bs []byte) error {

	var s string

	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}

	r, size := utf8.DecodeRuneInString(s)
	if size != len(s) {
		return errors.New("UnmarshalJSON")
	}

	*this = Id(r)

	return nil
}

type Identifier struct {
	Id       Id
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
