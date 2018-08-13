package isogram

import "strings"

// IsIsogram checks if the input word contains repeating letters
func IsIsogram(word string) bool {
	if word == "" {
		return true
	}

	countMap := make(map[rune]int, len(word))

	for _, w := range strings.ToLower(word) {
		countMap[w]++
		if w != '-' && w != ' ' && countMap[w] > 1 {
			return false
		}
	}
	return true
}

// ===

func IsIsogramTwo(inp string) bool {
	if inp == "" {
		return true
	}

	inp = strings.ToLower(inp)

	for i := 0; i < len(inp); i++ {
		if strings.LastIndexByte(inp, inp[i]) > i && inp[i] != '-' && inp[i] != ' ' {
			return false
		}
	}
	return true
}
