package aoc2024

import (
	"fmt"
	"strconv"
	"strings"
)

func Day5Part1(input string) {
	d := Day5{}
	d.Parse(input)
	fmt.Printf("Day 5 Part 1: %v\n", d.SumMids(d.GetValidUpdates()))
}

func Day5Part2(input string) {
	d := Day5{}
	d.Parse(input)

	fmt.Printf("Day 5 Part 2: %v\n", d.SumMids(d.GetFixedUpdates()))
}

type Day5 struct {
	rules   [][]int
	updates [][]int
}

func (d *Day5) Parse(input string) {
	sections := strings.Split(input, "\n\n")
	d.rules = d.parseLinesByCharacter(sections[0], "|")
	d.updates = d.parseLinesByCharacter(sections[1], ",")
}

func (d *Day5) parseLinesByCharacter(input string, splitter string) [][]int {
	var results [][]int
	for _, line := range strings.Split(input, "\n") {
		var thisRule []int

		for _, page := range strings.Split(line, splitter) {
			thisPage, _ := strconv.Atoi(page)
			thisRule = append(thisRule, thisPage)
		}

		results = append(results, thisRule)
	}

	return results
}

func (d *Day5) GetValidUpdates() [][]int {
	var validUpdates [][]int

	for _, update := range d.updates {
		if !d.isUpdateValid(update) {
			continue
		}

		validUpdates = append(validUpdates, update)
	}

	return validUpdates
}

func (d *Day5) isUpdateValid(update []int) bool {
	for i := 0; i < len(update); i++ {
		for j := (i + 1); j < len(update); j++ {
			for _, rule := range d.rules {
				if rule[1] == update[i] && rule[0] == update[j] {
					return false
				}
			}
		}
	}

	return true
}

func (d *Day5) SumMids(collection [][]int) int {
	sum := 0

	for _, update := range collection {
		midpoint := len(update) / 2
		sum += update[midpoint]
	}

	return sum
}

func (d *Day5) Fix(update []int) []int {
	result := make([]int, len(update))

	for _, entry := range update {
		count := 0
		for _, otherEntry := range update {
			for _, rule := range d.rules {
				if rule[0] == entry && rule[1] == otherEntry {
					count += 1
				}
			}
		}

		pos := len(update) - count - 1

		result[pos] = entry
	}

	return result
}

func (d *Day5) GetFixedUpdates() [][]int {
	var invalidUpdates [][]int

	for _, update := range d.updates {
		if d.isUpdateValid(update) {
			continue
		}

		invalidUpdates = append(invalidUpdates, d.Fix(update))
	}

	return invalidUpdates
}
