package raindrops

import "strconv"

// Convert retuns "Pling", "Plang" or "Plong" if the input is 3, 5 or 7 correspondingly.
// In other cases, it returns the input integer converted into string.
func Convert(num int) string {

	out := ""

	if num%3 == 0 {
		out += "Pling"
	}
	if num%5 == 0 {
		out += "Plang"
	}
	if num%7 == 0 {
		out += "Plong"
	}

	if out == "" {
		out = strconv.Itoa(num)
	}

	return out
}
