package aoc2024

import (
	"reflect"
	"testing"
)

const (
	day6input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
)

func TestDay6Part1(t *testing.T) {
	day := Day6Load(day6input)

	want := 41
	got := day.WalkUntilExited()
	if want != got {
		t.Fatalf("visited incorrect number of locations, want %v got %v", want, got)
	}
}

func TestDay6Parser(t *testing.T) {
	day := Day6Load(day6input)

	wantLenGrid := 10
	gotLenGrid := len(day.grid)
	if gotLenGrid != wantLenGrid {
		t.Fatalf("grid was wrong height, want %v got %v", wantLenGrid, gotLenGrid)
	}

	wantLenFirstRow := 10
	gotLenFirstRow := len(day.grid[0])
	if gotLenFirstRow != wantLenFirstRow {
		t.Fatalf("grid was wrong width, want %v got %v", wantLenFirstRow, gotLenFirstRow)
	}

	wantPosition := []int{6, 4}
	gotPosition := day.currentPosition
	if !reflect.DeepEqual(wantPosition, gotPosition) {
		t.Fatalf("did not get correct initial position, want %v got %v", wantPosition, gotPosition)
	}

	if day.grid[6][4] != '.' {
		t.Fatalf("didnt replace start position with a .")
	}
}

func TestDay6SingleWalk(t *testing.T) {
	day := Day6Load(day6input)
	day.walk()

	wantPosition := []int{1, 4}
	gotPosition := day.currentPosition

	if !reflect.DeepEqual(wantPosition, gotPosition) {
		t.Fatalf("not at the expected position, want %v got %v", wantPosition, gotPosition)
	}

	wantDirection := 'R'
	gotDirection := day.currentDirection
	if wantDirection != gotDirection {
		t.Fatalf("not facing the correct direction, want %v got %v", wantDirection, gotDirection)
	}

	wantCount := 6
	gotCount := day.countVisited()
	if wantCount != gotCount {
		t.Fatalf("incorrect number of locations visited, want %v got %v", wantCount, gotCount)
	}
}

func TestDay6LoopDetection(t *testing.T) {
	day := Day6Load(day6input)
	day.grid[6][3] = 'O'

	want := -1
	got := day.WalkUntilExited()

	if want != got {
		t.Fatalf("didnt detect loop, uh oh, want %v got %v", want, got)
	}
}

func TestDay6GetLoopCounter(t *testing.T) {
	want := 6
	got := Day6FindAllLoops(day6input)
	if want != got {
		t.Fatalf("didnt find all loops, want %v got %v", want, got)
	}
}

func TestDay6BlockageDetection(t *testing.T) {
	day := Day6Load(day6input)
	day.grid[1][5] = 'O'

	want := -2
	got := day.WalkUntilExited()

	if want != got {
		t.Fatalf("didnt detect blockage correctly and interpreted a loop, want %v got %v", want, got)
	}
}
