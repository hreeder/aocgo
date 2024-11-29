package aoc2016

import (
	"fmt"
	"strings"
)

func Day2Part1(input string) {
	day := NewDay2()
	var code []string

	for _, instructions := range strings.Split(input, "\n") {
		day.Move(instructions)
		code = append(code, fmt.Sprintf("%v", day.CurrentPosition))
	}

	fmt.Printf("Part 1: %v\n", strings.Join(code, ""))
}

func Day2Part2(input string) {
	day := NewDay2()
	day.SetAlphaNumericKeypad()
	var code []string

	for _, instructions := range strings.Split(input, "\n") {
		day.Move(instructions)
		code = append(code, fmt.Sprintf("%v", day.CurrentPosition))
	}

	fmt.Printf("Part 2: %v\n", strings.Join(code, ""))
}

type Day2 struct {
	CurrentPosition string

	nextPositions map[string]map[string]string
}

func NewDay2() *Day2 {
	day2 := &Day2{
		CurrentPosition: "5",

		nextPositions: map[string]map[string]string{
			"1": {
				"R": "2",
				"D": "4",
			},
			"2": {
				"R": "3",
				"L": "1",
				"D": "5",
			},
			"3": {
				"L": "2",
				"D": "6",
			},
			"4": {
				"U": "1",
				"R": "5",
				"D": "7",
			},
			"5": {
				"U": "2",
				"R": "6",
				"L": "4",
				"D": "8",
			},
			"6": {
				"U": "3",
				"L": "5",
				"D": "9",
			},
			"7": {
				"U": "4",
				"R": "8",
			},
			"8": {
				"U": "5",
				"R": "9",
				"L": "7",
			},
			"9": {
				"U": "6",
				"L": "8",
			},
		},
	}

	return day2
}

func (d *Day2) SetAlphaNumericKeypad() {
	d.nextPositions = map[string]map[string]string{
		"1": {
			"D": "3",
		},
		"2": {
			"R": "3",
			"D": "6",
		},
		"3": {
			"U": "1",
			"L": "2",
			"R": "4",
			"D": "7",
		},
		"4": {
			"L": "3",
			"D": "8",
		},
		"5": {
			"R": "6",
		},
		"6": {
			"U": "2",
			"L": "5",
			"R": "7",
			"D": "A",
		},
		"7": {
			"U": "3",
			"L": "6",
			"R": "8",
			"D": "B",
		},
		"8": {
			"U": "4",
			"L": "7",
			"R": "9",
			"D": "C",
		},
		"9": {
			"L": "8",
		},
		"A": {
			"U": "6",
			"R": "B",
		},
		"B": {
			"U": "7",
			"L": "A",
			"R": "C",
			"D": "D",
		},
		"C": {
			"U": "8",
			"L": "B",
		},
		"D": {
			"U": "B",
		},
	}
}

func (d *Day2) Move(instructions string) {
	for _, instruction := range instructions {
		newPos, found := d.nextPositions[d.CurrentPosition][string(instruction)]
		if found {
			d.CurrentPosition = newPos
		}
	}
}
