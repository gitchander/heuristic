package mindshift

import (
	"unicode/utf8"
)

type Location struct {
	Start   Point
	Finish  Point
	Current Point
}

/*
	- unit
	- soother
	- bimbo
	- aggregate
	- group
*/

type Unit struct {
	Location Location
	Points   []Point
}

func collisionShiftUnit(units []Unit, index int, shift Point) bool {

	var (
		unitIndex  = units[index]
		shiftIndex = unitIndex.Location.Current.Add(shift)
	)

	for i, unit := range units {
		if i != index {
			if checkCollision(unitIndex.Points, shiftIndex,
				unit.Points, unit.Location.Current) {
				return true
			}
		}
	}

	return false
}

func checkSolved(units []Unit) bool {

	for _, unit := range units {
		if !unit.Location.Current.Equal(unit.Location.Finish) {
			return false
		}
	}

	return true
}

func makeUnits(config PuzzleConfig) ([]Unit, error) {

	pm := make(map[rune]*Unit)

	var identifiers = config.Primitives.Identifiers
	for _, identifier := range identifiers {

		r := stringToRune(identifier.Id)
		if r == utf8.RuneError {
			return nil, newError("wrong id value")
		}

		start, err := intsToPoint(identifier.Location.Start)
		if err != nil {
			return nil, err
		}

		finish, err := intsToPoint(identifier.Location.Finish)
		if err != nil {
			return nil, err
		}

		pm[r] = &Unit{
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

	var units []Unit
	for _, p := range pm {
		units = append(units, *p)
	}

	return units, nil
}

func checkDuplicatePoints(units []Unit) error {

	for _, unit := range units {
		if err := errorDuplicatePoints(unit.Points); err != nil {
			return err
		}
	}
	return nil
}

func checkLocations(units []Unit) error {

	ls := []string{"Start", "Finish", "Current"}
	for _, l := range ls {
		if err := checkCollisionLocation(units, l); err != nil {
			return err
		}
	}
	return nil
}

func checkCollisionLocation(units []Unit, location string) error {

	var f func(unit *Unit) (sliceA []Point, shiftA Point)

	switch location {
	case "Start":
		f = func(unit *Unit) (sliceA []Point, shiftA Point) {
			if unit != nil {
				sliceA = unit.Points
				shiftA = unit.Location.Start
			}
			return
		}

	case "Finish":
		f = func(unit *Unit) (sliceA []Point, shiftA Point) {
			if unit != nil {
				sliceA = unit.Points
				shiftA = unit.Location.Finish
			}
			return
		}

	case "Current":
		f = func(unit *Unit) (sliceA []Point, shiftA Point) {
			if unit != nil {
				sliceA = unit.Points
				shiftA = unit.Location.Current
			}
			return
		}

	default:
		return newError("Wrong location name")
	}

	const format = "unit collision \"%s\": p[%d] and p[%d]"

	n := len(units)
	for indexA := 0; indexA < n; indexA++ {
		sliceA, shiftA := f(&(units[indexA]))
		for indexB := indexA + 1; indexB < n; indexB++ {
			if indexA != indexB {
				sliceB, shiftB := f(&(units[indexB]))
				if checkCollision(sliceA, shiftA, sliceB, shiftB) {
					return newErrorf(format, location, indexA, indexB)
				}
			}
		}
	}

	return nil
}
