package aoc2016

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day1Part1(input string) {
	sample := NewDay1()
	sample.MoveMany(input)
	fmt.Printf("Part 1: %d\n", sample.Distance())
}

func Day1Part2(input string) {
	sample := NewDay1()
	sample.MoveStepped(input)
	fmt.Printf("Part 2: %d\n", sample.Distance())
}

type Day1 struct {
	CurrentDirection string
	X, Y             int
}

func NewDay1() *Day1 {
	sample := &Day1{
		CurrentDirection: "NORTH",
		X:                0,
		Y:                0,
	}

	return sample
}

func (d *Day1) updateDirection(turn byte) {
	if d.CurrentDirection == "NORTH" {
		if turn == 'L' {
			d.CurrentDirection = "WEST"
		} else if turn == 'R' {
			d.CurrentDirection = "EAST"
		}
	} else if d.CurrentDirection == "EAST" {
		if turn == 'L' {
			d.CurrentDirection = "NORTH"
		} else if turn == 'R' {
			d.CurrentDirection = "SOUTH"
		}
	} else if d.CurrentDirection == "SOUTH" {
		if turn == 'L' {
			d.CurrentDirection = "EAST"
		} else if turn == 'R' {
			d.CurrentDirection = "WEST"
		}
	} else if d.CurrentDirection == "WEST" {
		if turn == 'L' {
			d.CurrentDirection = "SOUTH"
		} else if turn == 'R' {
			d.CurrentDirection = "NORTH"
		}
	}
}

func (d *Day1) Move(instruction string) {
	d.updateDirection(instruction[0])

	steps, _ := strconv.Atoi(string(instruction[1:]))
	if d.CurrentDirection == "NORTH" {
		d.X += steps
	} else if d.CurrentDirection == "EAST" {
		d.Y += steps
	} else if d.CurrentDirection == "SOUTH" {
		d.X -= steps
	} else if d.CurrentDirection == "WEST" {
		d.Y -= steps
	}
}

func (d *Day1) MoveMany(instructions string) {
	for _, instruction := range strings.Split(instructions, ", ") {
		d.Move(instruction)
	}
}

func (d *Day1) Distance() int {
	return int(math.Abs(float64(d.X)) + math.Abs(float64(d.Y)))
}

func (d *Day1) MoveUntilVisitedTwice(instructions string) {
	visited := make(map[string]struct{})

	for _, instruction := range strings.Split(instructions, ", ") {
		d.Move(instruction)
		thisVisit := fmt.Sprintf("%d,%d", d.X, d.Y)

		_, alreadyVisited := visited[thisVisit]
		if alreadyVisited {
			break
		}

		visited[thisVisit] = struct{}{}
	}
}

func (d *Day1) MoveStepped(instructions string) {
	visited := make(map[string]struct{})
instructions:
	for _, instruction := range strings.Split(instructions, ", ") {
		d.updateDirection(instruction[0])
		steps, _ := strconv.Atoi(string(instruction[1:]))
		for range steps {
			if d.CurrentDirection == "NORTH" {
				d.X += 1
			} else if d.CurrentDirection == "EAST" {
				d.Y += 1
			} else if d.CurrentDirection == "SOUTH" {
				d.X -= 1
			} else if d.CurrentDirection == "WEST" {
				d.Y -= 1
			}

			thisVisit := fmt.Sprintf("%d,%d", d.X, d.Y)

			_, alreadyVisited := visited[thisVisit]
			if alreadyVisited {
				break instructions
			}

			visited[thisVisit] = struct{}{}
		}
	}
}
