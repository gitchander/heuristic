package mindshift

import (
	"github.com/gitchander/heuristic/math/hashxy"
)

type IPuzzle interface {
	MoveUp() bool
	MoveDown() bool
	MoveLeft() bool
	MoveRight() bool

	Select(Point) bool
	IsSolved() bool
	Reset()
}

type Puzzle struct {
	indexActive int
	units       []Unit
	hm           hashxy.Matrix
}

func NewPuzzle(config PuzzleConfig) (*Puzzle, error) {

	units, err := makeUnits(config)
	if err != nil {
		return nil, err
	}

	if err = checkDuplicatePoints(units); err != nil {
		return nil, err
	}

	if err = checkLocations(units); err != nil {
		return nil, err
	}

	hm := hashxy.NewMatrix(13, 11)
	hm.

	return &Puzzle{
		indexActive: 0,
		units:       units,
	}, nil
}

func (this *Puzzle) MoveUp() bool {
	return this.moveActiveUnit(Point{X: 0, Y: -1})
}

func (this *Puzzle) MoveDown() bool {
	return this.moveActiveUnit(Point{X: 0, Y: +1})
}

func (this *Puzzle) MoveLeft() bool {
	return this.moveActiveUnit(Point{X: -1, Y: 0})
}

func (this *Puzzle) MoveRight() bool {
	return this.moveActiveUnit(Point{X: +1, Y: 0})
}

func (this *Puzzle) moveActiveUnit(shift Point) bool {

	index := this.indexActive

	if (0 > index) || (index >= len(this.units)) {
		return false
	}

	if collisionShiftUnit(this.units, index, shift) {
		return false
	}

	unit := &(this.units[index])
	unit.Location.Current = unit.Location.Current.Add(shift)

	return true
}

func (this *Puzzle) Select(p Point) bool {

	if i := this.unitIndexByPoint(p); i != -1 {
		this.indexActive = i
		return true
	}
	return false
}

func (this *Puzzle) unitIndexByPoint(q Point) int {

	for i, unit := range this.units {

		ps := unit.Points
		curr := unit.Location.Current

		for _, p := range ps {
			if q.Equal(p.Add(curr)) {
				return i
			}
		}
	}

	return -1
}

func (this *Puzzle) IsSolved() bool {
	return checkSolved(this.units)
}

func (this *Puzzle) Reset() {
	for i := 0; i < len(this.units); i++ {
		unit := &(this.units[i])
		unit.Location.Current = unit.Location.Start
	}
}
