package aoc2024

import (
	"reflect"
	"testing"
)

const (
	day7input = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
)

func TestDay7Part1(t *testing.T) {
	day := Day7Parse(day7input)

	want := 3749
	got := day.SumValid(false)

	if want != got {
		t.Fatalf("did not sum valid operations correctly, want %v got %v", want, got)
	}
}

func TestDay7Parser(t *testing.T) {
	got := Day7Parse(day7input)

	firstOpDesiredWant := 190
	firstOpDesiredGot := got.operations[0].wantedResult

	if firstOpDesiredGot != firstOpDesiredWant {
		t.Fatalf("first operation did not have the right wanted result, want %v got %v", firstOpDesiredWant, firstOpDesiredGot)
	}

	firstOpElementsWant := []int{10, 19}
	firstOpElementsGot := got.operations[0].elements

	if !reflect.DeepEqual(firstOpElementsGot, firstOpElementsWant) {
		t.Fatalf("first operation did not parse elements correctly, want %v got %v", firstOpElementsWant, firstOpElementsGot)
	}
}

func TestDay7CalculatesSingles(t *testing.T) {
	day := Day7Parse(day7input)

	want := true
	got := day.operations[0].IsValid(false)

	if want != got {
		t.Fatalf("did not calculate correctly")
	}
}

func TestDay7Part1Calculator(t *testing.T) {
	wantedResults := []bool{
		true,
		true,
		false,
		false,
		false,
		false,
		false,
		false,
		true,
	}

	day := Day7Parse(day7input)

	for i, operation := range day.operations {
		want := wantedResults[i]
		got := operation.IsValid(false)
		if want != got {
			t.Fatalf("did not get correct result for operation %v (%v), want %v got %v", i, operation.String(), want, got)
		}
	}
}

func TestDay7WithConcatenators(t *testing.T) {
	input := `156: 15 6
7290: 6 8 6 15
192: 17 8 14`

	day := Day7Parse(input)

	for _, op := range day.operations {
		if !op.IsValid(true) {
			t.Fatalf("%v was invalid even with concatenation", op.String())
		}
	}
}
