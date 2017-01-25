//Package hamming provides the distance between two DNA strands
package hamming

import "errors"

const testVersion = 5

// Distance returns the number of byte differences for two DNA strands. If the strands are different lengths an error is returned.
func Distance(a, b string) (int, error) {
	var distance int
	var err error

	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				distance += 1
			}
		}
	} else {
		err = errors.New("Strings are not the same length")
	}

	return distance, err
}
