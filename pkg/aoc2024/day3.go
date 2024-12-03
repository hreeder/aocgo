package aoc2024

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day3Part1(input string) {
	fmt.Printf("Day 3 Part 1: %v\n", Day3ProcessCommands(input))
}

func Day3Part2(input string) {
	fmt.Printf("Day 3 Part 2: %v\n", Day3ProcessExtendedCommands(input))
}

func Day3ExtractCommands(input string) [][]int {
	var results [][]int

	parser := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	inner := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for _, result := range parser.FindAllStringIndex(input, -1) {
		thisResult := input[result[0]:result[1]]

		matches := inner.FindStringSubmatch(thisResult)
		left, _ := strconv.Atoi(matches[1])
		right, _ := strconv.Atoi(matches[2])

		results = append(results, []int{left, right, result[0]})
	}

	return results
}

func Day3ProcessCommands(input string) int {
	commands := Day3ExtractCommands(input)
	sum := 0

	for _, command := range commands {
		sum += (command[0] * command[1])
	}

	return sum
}

func Day3ProcessExtendedCommands(input string) int {
	commands := Day3ExtractCommands(input)
	enabler := regexp.MustCompile(`do\(\)|don't\(\)`)
	var disabled [][]int

	disabledStart := -1
	for _, enablerRange := range enabler.FindAllStringIndex(input, -1) {
		diff := enablerRange[1] - enablerRange[0]
		if diff != 4 && disabledStart == -1 {
			// dont
			disabledStart = enablerRange[0]
		} else if disabledStart != -1 && diff == 4 {
			// do
			disabled = append(disabled, []int{disabledStart, enablerRange[0]})
			disabledStart = -1
		}
	}

	sum := 0

	for _, command := range commands {
		isEnabled := true
		for _, disableRange := range disabled {
			if disableRange[0] < command[2] && command[2] < disableRange[1] {
				isEnabled = false
				break
			}
		}

		if isEnabled {
			result := (command[0] * command[1])
			sum += result
		}
	}

	return sum
}
