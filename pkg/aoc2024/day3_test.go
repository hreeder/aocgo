package aoc2024

import (
	"testing"
)

func TestDay3Part1(t *testing.T) {
}

func TestDay3CommandExtractor(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	wantAll := [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}}
	got := Day3ExtractCommands(input)

	for idx, want := range wantAll {
		thisGot := got[idx]
		if want[0] != thisGot[0] || want[1] != thisGot[1] {
			t.Fatalf("Did not extract commands successfully - want %v, got %v", want, thisGot)
		}
	}
}

func TestDay3CommandProcessor(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	want := 161
	got := Day3ProcessCommands(input)

	if want != got {
		t.Fatalf("Did not process the right commands, want %v, got %v", want, got)
	}
}

func TestDay3ExtendedCommandProcessor(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()dont()?mul(8,5))"
	want := 48
	got := Day3ProcessExtendedCommands(input)

	if want != got {
		t.Fatalf("Did not process the right commands, want %v, got %v", want, got)
	}
}
