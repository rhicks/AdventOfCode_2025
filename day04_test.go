package main

import "testing"

func TestCheckInput(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	got := day04_1(input)
	want := input

	if got != want {
		t.Errorf("day04_1(example) = %s, want %s", got, want)
	}
}

func TestInitializeEmptyGrid(t *testing.T) {

}
