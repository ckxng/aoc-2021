package advent

import "fmt"

type Puzzle struct {
	name   string
	result string
}

func Run() []Puzzle {
	results := []Puzzle{
		{
			name:   "Advent of Code 2021 Day 1 Puzzle 1",
			result: Day01_Puzzle1(input_day01),
		},
		{
			name:   "Advent of Code 2021 Day 1 Puzzle 2",
			result: Day01_Puzzle2(input_day01),
		},
		{
			name:   "Advent of Code 2021 Day 2 Puzzle 1",
			result: Day02_Puzzle1(input_day02),
		},
		{
			name:   "Advent of Code 2021 Day 2 Puzzle 2",
			result: Day02_Puzzle2(input_day02),
		},
		{
			name:   "Advent of Code 2021 Day 3 Puzzle 1",
			result: Day03_Puzzle1(input_day03),
		},
		{
			name:   "Advent of Code 2021 Day 3 Puzzle 2",
			result: Day03_Puzzle2(input_day03),
		},
	}

	return results
}

func FormatResults(results []Puzzle) string {
	var sresults string

	for i := 0; i < len(results); i++ {
		if sresults == "" {
			sresults = fmt.Sprintf("Puzzle: %s\nResult: %s\n", results[i].name, results[i].result)
		} else {
			sresults = fmt.Sprintf("%s---\nPuzzle: %s\nResult: %s\n", sresults, results[i].name, results[i].result)
		}
	}
	return sresults
}
