package advent

import "testing"

func TestDay01_Puzzle1(t *testing.T) {
	want := "7"
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	if got := Day01_Puzzle1(input); got != want {
		t.Errorf("Day01_Puzzle1() = %q, want %q", got, want)
	}
}

func TestDay01_Puzzle1_Empty(t *testing.T) {
	want := "0"
	input := []int{}
	if got := Day01_Puzzle1(input); got != want {
		t.Errorf("Day01_Puzzle1() = %q, want %q", got, want)
	}
}

func TestDay01_Puzzle2(t *testing.T) {
	want := "5"
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	if got := Day01_Puzzle2(input); got != want {
		t.Errorf("Day01_Puzzle2() = %q, want %q", got, want)
	}
}

func TestDay01_Puzzle2_Empty(t *testing.T) {
	want := "0"
	input := []int{}
	if got := Day01_Puzzle2(input); got != want {
		t.Errorf("Day01_Puzzle2() = %q, want %q", got, want)
	}
}
