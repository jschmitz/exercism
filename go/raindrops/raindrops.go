// Package raindrops converts a number to a string.
package raindrops

import "strconv"

const testVersion = 2

// Convert returns the string based on the factors of the number inputted.
func Convert(i int) string {
	var r string
	if i%3 == 0 {
		r = "Pling"
	}
	if i%5 == 0 {
		r += "Plang"
	}

	if i%7 == 0 {
		r += "Plong"
	}

	if r == "" {
		r = strconv.Itoa(i)
	}

	return r
}
