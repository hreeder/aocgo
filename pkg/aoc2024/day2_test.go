package aoc2024

import (
	"reflect"
	"testing"
)

func TestDay2Part1(t *testing.T) {

}

func TestDay2ReportParser(t *testing.T) {
	var tests = []struct {
		input   string
		numbers []int
	}{
		{"7 6 4 2 1", []int{7, 6, 4, 2, 1}},
		{"1 2 7 8 9", []int{1, 2, 7, 8, 9}},
		{"9 7 6 2 1", []int{9, 7, 6, 2, 1}},
		{"1 3 2 4 5", []int{1, 3, 2, 4, 5}},
		{"8 6 4 4 1", []int{8, 6, 4, 4, 1}},
		{"1 3 6 7 9", []int{1, 3, 6, 7, 9}},
	}

	for _, tt := range tests {
		parsed := Day2ParseReport(tt.input)
		if !reflect.DeepEqual(tt.numbers, parsed.numbers) {
			t.Fatalf("Did not parse %v correctly. Wanted %v, Got %v", tt.input, tt.numbers, parsed.numbers)
		}
	}
}

func TestDay2DirectionParse(t *testing.T) {
	var tests = []struct {
		input     string
		direction Day2Direction
		valid     bool
	}{
		{"7 6 4 2 1", D2Down, true},
		{"1 2 7 8 9", D2Up, true},
		{"9 7 6 2 1", D2Down, true},
		{"1 3 2 4 5", 0, false},
		{"8 6 4 4 1", 0, false},
		{"1 3 6 7 9", D2Up, true},
	}

	for _, tt := range tests {
		parsed := Day2ParseReport(tt.input)
		direction, err := parsed.Direction()

		if err != nil && tt.valid == true {
			t.Fatalf("Threw error parsing direction for %v", tt.input)
		}

		if err == nil && tt.valid == false {
			t.Fatalf("Didn't throw error parsing direction for %v", tt.input)
		}

		if err == nil && tt.direction != direction {
			t.Fatalf("Wrong Direction for %v. Wanted %v, Got %v", tt.input, tt.direction, direction)
		}
	}
}

func TestDay2GetSafeties(t *testing.T) {
	var tests = []struct {
		input string
		safe  bool
	}{
		{"7 6 4 2 1", true},
		{"1 2 7 8 9", false},
		{"9 7 6 2 1", false},
		{"1 3 2 4 5", false},
		{"8 6 4 4 1", false},
		{"1 3 6 7 9", true},
	}

	for _, tt := range tests {
		parsed := Day2ParseReport(tt.input)
		isSafe := parsed.IsSafe()
		if isSafe != tt.safe {
			t.Fatalf("Incorrect safety reported for %v - wanted %v, got %v", tt.input, tt.safe, isSafe)
		}
	}
}

func TestDay2GetDampenedSafeties(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	want := 4
	got := Day2GetDampenedSafe(input)

	if want != got {
		t.Fatalf("Incorrect dampened safeties, wanted %v, got %v", want, got)
	}
}
