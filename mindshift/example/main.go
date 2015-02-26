package main

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	"github.com/gitchander/heuristic/mindshift"
)

func StringToRunes(s string) []rune {

	rs := make([]rune, 0)
	runeIndex := 0
	for {
		r, runeSize := utf8.DecodeRuneInString(s[runeIndex:])
		if runeSize == 0 {
			break
		}
		runeIndex += runeSize
		rs = append(rs, r)
	}

	return rs
}

func StringTest() {

	s := "漢語客家普通话ІіЇїЄє"
	rs := StringToRunes(s)
	for i, r := range rs {
		fmt.Printf("%d: %c\n", i, r)
	}
}

func Test() {

	/*

		str := `{
			"Name": "Puzzle 1",
			"Author": "Chander",
			"Description": "Description Puzzle",
			"Difficulty": 1,
			"Primitives": {
				"Cells": [
					"223",
					"111",
					"332"
				],
				"Identifiers": [
					{ "Id": "1", "Location": { "Start": [0, 4], "Finish": [0, 0] } },
					{ "Id": "2", "Location": { "Start": [4, -2], "Finish": [0, 0] } },
					{ "Id": "3", "Location": { "Start": [-4, -2], "Finish": [0, 0] } }
				]
			}
		}`

	*/

	str := `{
			"Name": "Cross 2",
			"Author": "Chander",
			"Description": "",
			"Difficulty": 1,
			"Primitives": {
				"Cells": [
					"----456----",
					"----456----",
					"----456----",
					"----456----",
					"11115231111",
					"22224-52222",
					"33333163333",
					"----456----",
					"----456----",
					"----456----",
					"----456----"
				],
				"Identifiers": [
					{ "Id": "1", "Location": { "Start": [-14, -4], "Finish": [0, 0] } },
					{ "Id": "2", "Location": { "Start": [-14, 0], "Finish": [0, 0] } },
					{ "Id": "3", "Location": { "Start": [-14, 4], "Finish": [0, 0] } },
					{ "Id": "4", "Location": { "Start": [20, 0], "Finish": [0, 0] } },
					{ "Id": "5", "Location": { "Start": [14, 0], "Finish": [0, 0] } },
					{ "Id": "6", "Location": { "Start": [8, 0], "Finish": [0, 0] } }
				]
			}
		}`

	var pc mindshift.PuzzleConfig

	err := json.Unmarshal([]byte(str), &pc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	puzzle, err := mindshift.NewPuzzle(pc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", puzzle)

}

func main() {

	Test()
}
