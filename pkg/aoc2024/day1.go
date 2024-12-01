package aoc2024

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day1Part1(input string) {
	fmt.Printf("Day 1 Part 1: %v\n", day1Part1GetTotal(input))
}

func Day1Part2(input string) {
	fmt.Printf("Day 1 Part 2: %v\n", day1Part2GetSimilarity(input))
}

func day1ParseList(data string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, "   ")

		newLeft, _ := strconv.Atoi(parts[0])
		left = append(left, newLeft)

		newRight, _ := strconv.Atoi(parts[1])
		right = append(right, newRight)
	}

	return left, right
}

func day1Part1GetTotal(data string) int {
	left, right := day1ParseList(data)

	slices.Sort(left)
	slices.Sort(right)

	currentSum := 0

	for index := range left {
		distance := int(math.Abs(float64(left[index] - right[index])))
		// fmt.Printf("Comparing left[%v], right[%v] - %v-%v=%v\n", index, index, left[index], right[index], distance)
		currentSum += distance
	}

	return currentSum
}

func day1Part2GetSimilarity(data string) int {
	left, right := day1ParseList(data)
	score := 0

	for _, leftVal := range left {
		count := 0
		for _, rightVal := range right {
			if leftVal == rightVal {
				count += 1
			}
		}
		score += (count * leftVal)
	}

	return score
}
