package aoc2016

import "testing"

func TestDay2Part1(t *testing.T) {
	day := NewDay2()

	var tests = []struct {
		move   string
		newPos string
	}{
		{"ULL", "1"},
		{"RRDDD", "9"},
		{"LURDL", "8"},
		{"UUUUD", "5"},
	}

	for _, tt := range tests {
		day.Move(tt.move)
		if day.CurrentPosition != tt.newPos {
			t.Fatalf("Step Failed. Tried to move %v, Expected position %v, Got Position %v", tt.move, tt.newPos, day.CurrentPosition)
		}
	}
}

func TestDay2Part2(t *testing.T) {
	day := NewDay2()
	day.SetAlphaNumericKeypad()

	var tests = []struct {
		move   string
		newPos string
	}{
		{"ULL", "5"},
		{"RRDDD", "D"},
		{"LURDL", "B"},
		{"UUUUD", "3"},
	}

	for _, tt := range tests {
		day.Move(tt.move)
		if day.CurrentPosition != tt.newPos {
			t.Fatalf("Step Failed. Tried to move %v, Expected position %v, Got Position %v", tt.move, tt.newPos, day.CurrentPosition)
		}
	}
}
