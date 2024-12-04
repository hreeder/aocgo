package aoc2024

import "testing"

const (
	day4input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
)

func TestDay4Part1(t *testing.T) {
	got := Day4CountXmas(day4input)
	want := 18

	if got != want {
		t.Fatalf("Did not calculate correct amount of xmas, want %v got %v", want, got)
	}
}

func TestDay4Part2(t *testing.T) {
	got := Day4CountMas(day4input)
	want := 9

	if got != want {
		t.Fatalf("Did not calculate the correct amount of x-mas, want %v got %v", want, got)
	}
}
