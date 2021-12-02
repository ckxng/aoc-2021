package advent

import "testing"

func TestDay02_Puzzle1(t *testing.T) {
	want := "150"
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	if got := Day02_Puzzle1(input); got != want {
		t.Errorf("Day02_Puzzle1() = %q, want %q", got, want)
	}
}

func TestDay02_Puzzle2(t *testing.T) {
	want := "900"
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	if got := Day02_Puzzle2(input); got != want {
		t.Errorf("Day02_Puzzle2() = %q, want %q", got, want)
	}
}
