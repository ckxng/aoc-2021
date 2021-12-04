package advent

import (
	"fmt"
	"strconv"
	"strings"
)

// using string builder method for puzzle 1
func Day03_Puzzle1(input []string) string {
	// no input means the product must be 0 anyway
	if len(input) < 1 {
		return "0"
	}

	// strings to append "bits" to
	var sbGamma strings.Builder
	var sbEpsilon strings.Builder

	// totals
	var totZeros []int
	var totOnes []int

	// constants to compare
	zero := byte(48) // "0"
	one := byte(49)  // "1"

	// use the length of the first item as our bit size for gamma and epsilon
	length := len(input[0])

	for pos := 0; pos < length; pos++ {
		totZeros = append(totZeros, 0)
		totOnes = append(totOnes, 0)

		// check this position of each array element
		// skip any non-existant positions if the element is undersized
		// this behavior is undefined by the puzzle
		for i := 0; i < len(input) && pos < len(input[i]); i++ {
			switch input[i][pos] {
			case zero:
				totZeros[pos]++
			case one:
				totOnes[pos]++
			}
		}

		// if there is the same number of zeros as ones, the behavior is undefined
		// by the puzzle
		if totZeros[pos] > totOnes[pos] {
			sbGamma.WriteString("0")
			sbEpsilon.WriteString("1")
		} else {
			sbGamma.WriteString("1")
			sbEpsilon.WriteString("0")
		}
	}

	iGamma, _ := strconv.ParseInt(sbGamma.String(), 2, 64)
	iEpsilon, _ := strconv.ParseInt(sbEpsilon.String(), 2, 64)

	return fmt.Sprintf("%d", iGamma*iEpsilon)
}

// using bit manipulation for puzzle 2
func Day03_Puzzle2(input []string) string {
	// no input means the product must be 0 anyway
	if len(input) < 1 {
		return "0"
	}

	return "0"
}
