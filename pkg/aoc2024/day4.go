package aoc2024

import (
	"fmt"
	"strings"
)

func Day4Part1(input string) {
	fmt.Printf("Day 4 Part 1: %d\n", Day4CountXmas(input))
}

func Day4Part2(input string) {
	fmt.Printf("Day 4 Part 2: %d\n", Day4CountMas(input))
}

func day4ParseGrid(input string) [][]rune {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {

		var thisLine []rune
		for _, char := range line {
			thisLine = append(thisLine, char)
		}

		grid = append(grid, thisLine)
	}

	return grid
}

func Day4CountXmas(input string) int {
	grid := day4ParseGrid(input)
	maxX := len(grid)
	maxY := len(grid[0])
	sum := 0

	fmt.Printf("max: x=%d, y=%d\n", maxX, maxY)

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] != 'X' {
				continue
			}

			if x >= 3 && grid[x-1][y] == 'M' && grid[x-2][y] == 'A' && grid[x-3][y] == 'S' {
				// N
				sum += 1
				// fmt.Printf("N XMAS at %d,%d\n", x, y)
			}

			if x >= 3 && (y+3) < maxY && grid[x-1][y+1] == 'M' && grid[x-2][y+2] == 'A' && grid[x-3][y+3] == 'S' {
				// NE
				sum += 1
				// fmt.Printf("NE XMAS at %d,%d\n", x, y)
			}

			if (y+3) < maxY && grid[x][y+1] == 'M' && grid[x][y+2] == 'A' && grid[x][y+3] == 'S' {
				// E
				sum += 1
				// fmt.Printf("E XMAS at %d,%d\n", x, y)
			}

			if (x+3) < maxX && (y+3) < maxY && grid[x+1][y+1] == 'M' && grid[x+2][y+2] == 'A' && grid[x+3][y+3] == 'S' {
				// SE
				sum += 1
				// fmt.Printf("SE XMAS at %d,%d\n", x, y)
			}

			if (x+3) < maxX && grid[x+1][y] == 'M' && grid[x+2][y] == 'A' && grid[x+3][y] == 'S' {
				// S
				sum += 1
				// fmt.Printf("S XMAS at %d,%d\n", x, y)
			}

			if (x+3) < maxX && y >= 3 && grid[x+1][y-1] == 'M' && grid[x+2][y-2] == 'A' && grid[x+3][y-3] == 'S' {
				// SW
				sum += 1
				// fmt.Printf("SW XMAS at %d,%d\n", x, y)
			}

			if y >= 3 && grid[x][y-1] == 'M' && grid[x][y-2] == 'A' && grid[x][y-3] == 'S' {
				// W
				sum += 1
				// fmt.Printf("W XMAS at %d,%d\n", x, y)
			}

			if y >= 3 && x >= 3 && grid[x-1][y-1] == 'M' && grid[x-2][y-2] == 'A' && grid[x-3][y-3] == 'S' {
				// NW
				sum += 1
				// fmt.Printf("NW XMAS at %d,%d\n", x, y)
			}
		}
	}

	return sum
}

func Day4CountMas(input string) int {
	grid := day4ParseGrid(input)
	maxX := len(grid)
	maxY := len(grid[0])
	sum := 0
	fmt.Printf("max: x=%d, y=%d\n", maxX, maxY)

	for x := range grid {
		if x == 0 || x == (maxX-1) {
			continue
		}

		for y := range grid[x] {
			if y == 0 || y == (maxY-1) {
				continue
			}

			if grid[x][y] != 'A' {
				continue
			}

			if (grid[x-1][y-1] == 'M' || grid[x+1][y+1] == 'M') && (grid[x-1][y-1] == 'S' || grid[x+1][y+1] == 'S') && (grid[x-1][y+1] == 'M' || grid[x+1][y-1] == 'M') && (grid[x-1][y+1] == 'S' || grid[x+1][y-1] == 'S') {
				sum += 1
			}
		}
	}

	return sum
}
