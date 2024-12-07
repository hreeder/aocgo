package aoc2024

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func Day6Part1(input string) {
	day6 := Day6Load(input)
	fmt.Printf("Day 6 Part 1: %v\n", day6.WalkUntilExited())
}

func Day6Part2(input string) {

	loops := Day6FindAllLoops(input)
	fmt.Printf("Day 6 Part 2: %v\n", loops)
}

type Day6Map struct {
	grid [][]rune

	currentPosition  []int
	currentDirection rune

	turnLocations [][]int
}

func Day6FindAllLoops(input string) int {
	sum := 0
	blank := Day6Load(input)

	bar := progressbar.Default(int64(len(blank.grid)))

	for i := range blank.grid {
		for j := range blank.grid[i] {
			inner := Day6Load(input)
			if inner.grid[i][j] == '^' {
				continue
			}

			inner.grid[i][j] = 'O'
			if inner.WalkUntilExited() == -1 {
				sum += 1
			}
		}
		bar.Add(1)
	}

	return sum
}

func Day6Load(input string) *Day6Map {
	currentPosition := []int{0, 0}
	var grid [][]rune
	for x, line := range strings.Split(input, "\n") {
		var gridLine []rune

		for y, char := range line {
			thisChar := char
			if char == '^' {
				currentPosition = []int{x, y}
				thisChar = '.'
			}
			gridLine = append(gridLine, thisChar)
		}

		grid = append(grid, gridLine)
	}

	return &Day6Map{
		grid:             grid,
		currentPosition:  currentPosition,
		currentDirection: 'U',
	}
}

func (d *Day6Map) countVisited() int {
	count := 0

	for _, row := range d.grid {
		for _, char := range row {
			if char == 'X' {
				count += 1
			}
		}
	}

	return count
}

func (d *Day6Map) WalkUntilExited() int {
	var err error

	for err == nil {
		loop := d.walk()
		if loop != nil {
			if loop.Error() == "loop detected" {
				return -1
			}
			return -2
		}

		_, err = d.peek()
	}
	return d.countVisited()
}

func (d *Day6Map) walk() error {
	keepGoing := true
	moveCount := 0

	for keepGoing {
		d.grid[d.currentPosition[0]][d.currentPosition[1]] = 'X'

		nextItem, err := d.peek()
		if err != nil {
			keepGoing = false
			break
		}

		if nextItem != '.' && nextItem != 'X' {
			// if moveCount == 0 {
			// 	return errors.New("invalid grid")
			// }

			priorTurnCount := 0
			for _, priorTurn := range d.turnLocations {
				if reflect.DeepEqual(priorTurn, d.currentPosition) {
					priorTurnCount += 1
				}
			}
			if priorTurnCount > 3 {
				return errors.New("loop detected")
			}
			d.turnLocations = append(d.turnLocations, d.currentPosition)

			switch d.currentDirection {
			case 'U':
				d.currentDirection = 'R'
			case 'R':
				d.currentDirection = 'D'
			case 'D':
				d.currentDirection = 'L'
			case 'L':
				d.currentDirection = 'U'
			}

			keepGoing = false
			break
		} else {
			nextPosition, _ := d.nextPosition()
			d.currentPosition = nextPosition
			moveCount += 1
		}
	}

	return nil
}

func (d *Day6Map) peek() (rune, error) {
	nextPosition, err := d.nextPosition()
	if err != nil {
		return 0, err
	}

	return d.grid[nextPosition[0]][nextPosition[1]], nil
}

func (d *Day6Map) nextPosition() ([]int, error) {
	var next []int
	switch d.currentDirection {
	case 'U':
		next = []int{d.currentPosition[0] - 1, d.currentPosition[1]}
	case 'R':
		next = []int{d.currentPosition[0], d.currentPosition[1] + 1}
	case 'D':
		next = []int{d.currentPosition[0] + 1, d.currentPosition[1]}
	case 'L':
		next = []int{d.currentPosition[0], d.currentPosition[1] - 1}
	}

	if next[0] < 0 || next[0] >= len(d.grid) || next[1] < 0 || next[1] >= len(d.grid[0]) {
		return nil, errors.New("next position out of bounds")
	}

	return next, nil
}
