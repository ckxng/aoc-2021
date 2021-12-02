package advent

import "fmt"

func Day01_Puzzle1(input []int) string {
	count := 0

	// skip the first element, as it has no predecessor
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

func Day01_Puzzle2(input []int) string {
	// arrays with less than 4 elements do not have enough data for two sliding
	// windows to show an increase
	if len(input) < 4 {
		return "0"
	}

	count := 0
	previous := input[0] + input[1] + input[2]

	// skip the first three elements, as we are using a 3-element sliding window
	// and comparison begins with the second window
	for i := 3; i < len(input); i++ {
		current := input[i] + input[i-1] + input[i-2]
		if current > previous {
			count++
		}

		// setup previous for next iteration
		previous = current
	}

	return fmt.Sprintf("%d", count)
}
