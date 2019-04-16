// Package hamming provides function to calculate the Hamming
// distance between two DNA strands
package hamming

import "errors"

// Distance receives two DNA strands and returns the Hamming distance
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA strands have different length")
	}
	if len(a) == 0 {
		return 0, nil
	}
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			distance++
		}
	}
	return distance, nil
}
