package aoc2024

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Day2Direction int

const (
	D2Up   Day2Direction = iota
	D2Down Day2Direction = iota
)

func Day2Part1(input string) {
	safeReports := 0
	for _, line := range strings.Split(input, "\n") {
		parsed := Day2ParseReport(line)
		if parsed.IsSafe() {
			safeReports += 1
		}
	}

	fmt.Printf("Day 2 Part 1: %v\n", safeReports)
}

func Day2Part2(input string) {
	safeReports := Day2GetDampenedSafe(input)
	fmt.Printf("Day 2 Part 2: %v\n", safeReports)
}

type Day2Report struct {
	numbers []int
}

func Day2ParseReport(line string) *Day2Report {
	var numbers []int

	for _, char := range strings.Split(line, " ") {
		number, _ := strconv.Atoi(char)
		numbers = append(numbers, number)
	}

	return &Day2Report{numbers: numbers}
}

func (d *Day2Report) Direction() (Day2Direction, error) {
	var currentDirection Day2Direction

	first := d.numbers[0]
	second := d.numbers[1]

	if first == second {
		return 0, errors.New("first numbers are the same")
	}

	if first > second {
		currentDirection = D2Down
	} else {
		currentDirection = D2Up
	}

	for index := range d.numbers {
		if index == 0 {
			continue
		}

		if d.numbers[index] == d.numbers[index-1] {
			return 0, errors.New("number was the same as previous")
		}

		if currentDirection == D2Down {
			if d.numbers[index] > d.numbers[index-1] {
				return 0, errors.New("number went up while descending")
			}
		} else {
			if d.numbers[index] < d.numbers[index-1] {
				return 0, errors.New("number went down while ascending")
			}
		}
	}

	return currentDirection, nil
}

func (d *Day2Report) IsSafe() bool {
	direction, err := d.Direction()
	if err != nil {
		return false
	}

	for index := range d.numbers {
		if index == 0 {
			continue
		}

		var difference int

		if direction == D2Down {
			difference = d.numbers[index-1] - d.numbers[index]
		} else {
			difference = d.numbers[index] - d.numbers[index-1]
		}

		if difference < 1 || difference > 3 {
			return false
		}
	}

	return true
}

func Day2GetDampenedSafe(input string) int {
	count := 0

outer:
	for _, line := range strings.Split(input, "\n") {
		regular := Day2ParseReport(line)
		if regular.IsSafe() {
			count += 1
			continue
		}

		for index := range regular.numbers {
			dampened := Day2ParseReport(line)
			dampened.numbers = append(dampened.numbers[:index], dampened.numbers[index+1:]...)
			if dampened.IsSafe() {
				count += 1
				continue outer
			}
		}
	}

	return count
}
