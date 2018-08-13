package scrabble

import "strings"

var mapping = map[string]int{
	"A": 1,
	"E": 1,
	"I": 1,
	"O": 1,
	"U": 1,
	"L": 1,
	"N": 1,
	"R": 1,
	"S": 1,
	"T": 1,
	"D": 2,
	"G": 2,
	"B": 3,
	"C": 3,
	"M": 3,
	"P": 3,
	"F": 4,
	"H": 4,
	"V": 4,
	"W": 4,
	"Y": 4,
	"K": 5,
	"J": 8,
	"X": 8,
	"Q": 10,
	"Z": 10,
}

// Score computes the value of an input word,
// summing up values of all its letters,
// based on the `mapping` map.
func Score(input string) int {

	score := 0

	for _, i := range strings.ToUpper(input) {
		score += mapping[string(i)]
	}

	return score
}
