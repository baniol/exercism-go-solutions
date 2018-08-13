package hamming

import (
	"errors"
)

// Distance calculates Hamming difference between two strings
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("Strands not equal")
	}

	mismatchCount := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			mismatchCount++
		}
	}

	return mismatchCount, nil

}
