package advent

import "testing"

func TestRun(t *testing.T) {
	if got := Run(); len(got) < 1 {
		t.Errorf("Run() = length %q, want length > 0", len(got))
	}
}

func TestFormatResults(t *testing.T) {
	want := "Puzzle: A\nResult: B\n---\nPuzzle: C\nResult: D\n"
	input := []Puzzle{{name: "A", result: "B"}, {name: "C", result: "D"}}
	if got := FormatResults(input); got != want {
		t.Errorf("FormatResults() = %q, want %q", got, want)
	} else {
		t.Logf("FormatResults() want ")
	}
}
