// Package hamming provides function to calculate the Hamming
// distance between two DNA strands
package hamming

import "errors"

// Distance receives two DNA strands and returns the Hamming distance
func Distance(strand1, strand2 string) (distance int, err error) {
	if len(strand1) != len(strand2) {
		return 0, errors.New("DNA strands have different length")
	}
	if len(strand1) == 0 {
		return 0, nil
	}
	for index := 0; index < len(strand1); index++ {
		if strand1[index] != strand2[index] {
			distance++
		}
	}
	return distance, nil
}
