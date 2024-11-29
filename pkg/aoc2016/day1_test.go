package aoc2016

import "testing"

func TestDay1Part1MovesAsExpected(t *testing.T) {
	sample := NewDay1()
	sample.Move("R2")

	if sample.CurrentDirection != "EAST" {
		t.Fatalf("New direction is not EAST, got %s", sample.CurrentDirection)
	}

	if sample.Y != 2 {
		t.Fatalf("New position is incorrect, expected y=2, got y=%d", sample.Y)
	}
}

func TestDay1Part1MovesMultiple(t *testing.T) {
	sample := NewDay1()
	sample.MoveMany("R2, L3")

	if sample.Y != 2 {
		t.Fatalf("Y != 2")
	}
	if sample.X != 3 {
		t.Fatalf("X != 3")
	}
}

func TestDay1Part1GetsCorrectDistances(t *testing.T) {
	var tests = []struct {
		steps    string
		distance int
	}{
		{"R2, L3", 5},
		{"R2, R2, R2", 2},
		{"R5, L5, R5, R3", 12},
	}

	for _, tt := range tests {
		sample := NewDay1()
		sample.MoveMany(tt.steps)

		want := tt.distance
		got := sample.Distance()

		if want != got {
			t.Fatalf("Got %d, wanted %d", got, want)
		}
	}
}

func TestDay1Part2(t *testing.T) {
	input := "R8, R4, R4, R8"
	expected := 4

	sample := NewDay1()
	sample.MoveStepped(input)

	if expected != sample.Distance() {
		t.Fatalf("Not Right")
	}
}
