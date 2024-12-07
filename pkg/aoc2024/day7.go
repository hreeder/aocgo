package aoc2024

import (
	"fmt"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func Day7Part1(input string) {
	day := Day7Parse(input)
	fmt.Printf("Day 7 Part 1: %v\n", day.SumValid(false))
}

func Day7Part2(input string) {
	day := Day7Parse(input)
	fmt.Printf("Day 7 Part 2: %v\n", day.SumValid(true))
}

type Day7 struct {
	operations []Day7Operation
}

type Day7Operation struct {
	wantedResult int
	elements     []int
}

func Day7Parse(input string) *Day7 {
	day := Day7{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")

		wanted, _ := strconv.Atoi(parts[0])
		var elements []int

		for _, element := range strings.Split(parts[1], " ") {
			thisElement, _ := strconv.Atoi(element)
			elements = append(elements, thisElement)
		}

		day.operations = append(day.operations, Day7Operation{
			wantedResult: wanted,
			elements:     elements,
		})
	}

	return &day
}

func (o *Day7Operation) String() string {
	output := fmt.Sprintf("%v: ", o.wantedResult)
	for _, elem := range o.elements {
		output += fmt.Sprintf("%v ", elem)
	}

	return strings.Trim(output, " ")
}

func (o *Day7Operation) IsValid(withConcat bool) bool {
	operatorPositionCount := len(o.elements) - 1
	if operatorPositionCount == 1 {
		if (o.elements[0] + o.elements[1]) == o.wantedResult {
			return true
		}

		if withConcat && fmt.Sprintf("%v%v", o.elements[0], o.elements[1]) == strconv.Itoa(o.wantedResult) {
			return true
		}

		return (o.elements[0] * o.elements[1]) == o.wantedResult
	}

	validOperators := []rune{'+', '*'}
	if withConcat {
		validOperators = append(validOperators, '|')
	}

	operatorPositions := make([]int, operatorPositionCount)
	for i := range operatorPositions {
		operatorPositions[i] = len(validOperators)
	}
	productGen := combin.NewCartesianGenerator(operatorPositions)

	for productGen.Next() {
		product := productGen.Product(nil)

		elements := make([]int, len(o.elements))
		copy(elements, o.elements)

		current := elements[0]
		for position, operator := range product {
			nextValue := elements[position+1]

			if validOperators[operator] == '+' {
				current += nextValue
			}

			if validOperators[operator] == '*' {
				current = current * nextValue
			}

			if validOperators[operator] == '|' {
				newNumStr := fmt.Sprintf("%v%v", current, nextValue)
				newNum, _ := strconv.Atoi(newNumStr)
				current = newNum
			}
		}

		if current == o.wantedResult {
			return true
		}
	}

	return false
}

func (d *Day7) SumValid(withConcat bool) int {
	sum := 0

	for _, operation := range d.operations {
		if operation.IsValid(withConcat) {
			sum += operation.wantedResult
		}
	}

	return sum
}
