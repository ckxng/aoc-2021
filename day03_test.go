package advent

import "testing"

func TestDay03_Puzzle1(t *testing.T) {
	want := "198"
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	if got := Day03_Puzzle1(input); got != want {
		t.Errorf("Day03_Puzzle1() = %q, want %q", got, want)
	}
}

// behavior undefined by puzzle
func TestDay03_Puzzle1_Mismatched(t *testing.T) {
	want := "10"
	input := []string{
		"101",
		"01",
		"1010",
	}
	if got := Day03_Puzzle1(input); got != want {
		t.Errorf("Day03_Puzzle1() = %q, want %q", got, want)
	}
}

// behavior undefined by puzzle
func TestDay03_Puzzle1_Tie(t *testing.T) {
	want := "240"
	input := []string{
		"00011",
		"01101",
	}
	if got := Day03_Puzzle1(input); got != want {
		t.Errorf("Day03_Puzzle1() = %q, want %q", got, want)
	}
}

func TestDay03_Puzzle1_Empty(t *testing.T) {
	want := "0"
	input := []string{}
	if got := Day03_Puzzle1(input); got != want {
		t.Errorf("Day03_Puzzle1() = %q, want %q", got, want)
	}
}

func TestDay03_Puzzle2(t *testing.T) {
	want := "230"
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	if got := Day03_Puzzle2(input); got != want {
		t.Errorf("Day03_Puzzle2() = %q, want %q", got, want)
	}
}

func TestDay03_Puzzle2_Empty(t *testing.T) {
	want := "0"
	input := []string{}
	if got := Day03_Puzzle2(input); got != want {
		t.Errorf("Day03_Puzzle2() = %q, want %q", got, want)
	}
}
