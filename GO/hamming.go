package hamming

import (
	"errors"
)

//Calculates hamming distance between two strings if they have equal length else throws an error
func Distance(a, b string) (int, error) {

	distance := 0
	if len(a) != len(b) {
		return 0, errors.New("unequal string length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
