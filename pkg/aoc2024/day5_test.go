package aoc2024

import (
	"reflect"
	"testing"
)

const (
	day5Input = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
)

func TestDay5Part1(t *testing.T) {
	day := Day5{}
	day.Parse(day5Input)
	want := 143
	got := day.SumMids(day.GetValidUpdates())
	if want != got {
		t.Fatalf("Could not sum correctly, want %v got %v", want, got)
	}
}

func TestDay5Parser(t *testing.T) {
	day := Day5{}
	day.Parse(day5Input)

	wantLen := 21
	if len(day.rules) != wantLen {
		t.Fatalf("Not enough rules were parsed, want %v got %v", wantLen, len(day.rules))
	}

	wantFirst := []int{47, 53}
	if !reflect.DeepEqual(day.rules[0], wantFirst) {
		t.Fatalf("First rule was not as expected, want %v got %v", wantFirst, day.rules[0])
	}

	wantUpdatesLen := 6
	if len(day.updates) != wantUpdatesLen {
		t.Fatalf("Not enough updates were parsed, want %v got %v", wantUpdatesLen, len(day.updates))
	}

	wantUpdatesFirst := []int{75, 47, 61, 53, 29}
	if !reflect.DeepEqual(day.updates[0], wantUpdatesFirst) {
		t.Fatalf("First update was not as expected, want %v got %v", wantUpdatesFirst, day.updates[0])
	}
}

func TestDay5Validator(t *testing.T) {
	day := Day5{}
	day.Parse(day5Input)

	expected := []bool{true, true, true, false, false, false}
	for i, want := range expected {
		got := day.isUpdateValid(day.updates[i])
		if got != want {
			t.Fatalf("Update validation failed for update %v - want %v got %v", i, want, got)
		}
	}
}

func TestDay5Fixer(t *testing.T) {
	day := Day5{}
	day.Parse(day5Input)

	var tests = []struct {
		input []int
		want  []int
	}{
		{
			[]int{75, 97, 47, 61, 53},
			[]int{97, 75, 47, 61, 53},
		},
		{
			[]int{61, 13, 29},
			[]int{61, 29, 13},
		},
		{
			[]int{97, 13, 75, 29, 47},
			[]int{97, 75, 47, 29, 13},
		},
	}
	for _, tt := range tests {
		got := day.Fix(tt.input)

		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("Could not fix correctly, want %v, got %v", tt.want, got)
		}
	}
}
