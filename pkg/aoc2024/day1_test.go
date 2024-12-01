package aoc2024

import (
	"reflect"
	"testing"
)

const day1InputData = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestDay1Part1(t *testing.T) {
	wanted := 11
	got := day1Part1GetTotal(day1InputData)

	if wanted != got {
		t.Fatalf("Did not calculate total correctly. Wanted %v, Got %v\n", wanted, got)
	}
}

func TestDay1ListParser(t *testing.T) {
	wanted := [][]int{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}

	gotLeft, gotRight := day1ParseList(day1InputData)

	if !reflect.DeepEqual([][]int{gotLeft, gotRight}, wanted) {
		t.Fatalf("Could not parse list correctly\n")
	}
}

func TestDay1SimilarityScore(t *testing.T) {
	wanted := 31
	got := day1Part2GetSimilarity(day1InputData)

	if wanted != got {
		t.Fatalf("Incorrect similarity score. Wanted %v, Got %v\n", wanted, got)
	}
}
