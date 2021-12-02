package advent

import (
	"fmt"
	"strconv"
	"strings"
)

func Day02_Puzzle1(input []string) string {
	horizontal := 0
	depth := 0

	for i := 0; i < len(input); i++ {
		command := strings.Split(input[i], " ")
		factor, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "forward":
			horizontal += factor
		case "down":
			depth += factor
		case "up":
			depth -= factor
		}
	}

	return fmt.Sprintf("%d", horizontal*depth)
}

func Day02_Puzzle2(input []string) string {
	horizontal := 0
	depth := 0
	aim := 0

	for i := 0; i < len(input); i++ {
		command := strings.Split(input[i], " ")
		factor, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "forward":
			horizontal += factor
			depth += factor * aim
		case "down":
			aim += factor
		case "up":
			aim -= factor
		}
	}

	return fmt.Sprintf("%d", horizontal*depth)
}
