package main

import (
	"encoding/json"
	"fmt"

	"github.com/gitchander/heuristic/mindshift"
)

func main() {

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

	var pc mindshift.PuzzleConfig

	err := json.Unmarshal([]byte(str), &pc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bs, err := json.MarshalIndent(&pc, "", "\t")
	//bs, err := json.Marshal(&pc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(bs))
}
