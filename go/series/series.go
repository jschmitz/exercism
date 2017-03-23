// Package slice given a string of digits, output all the contiguous substrings of length `n` in that string.
package slice

const testVersion = 1

// All outputs all contigous substrings
func All(n int, s string) []string {
	var all []string

	for i := 0; i <= len(s); i++ {
		if i+n <= len(s) {
			all = append(all, string(s[i:i+n]))
		} else {
			return all
		}
	}
	return all
}

// UnsafeFirst returns the first substring unsafely of length n
func UnsafeFirst(n int, s string) string {
	return string(s[0:n])
}

// First safely returns the first substring of length n
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		ok = false
	} else {
		first = s[0:n]
		ok = true
	}

	return first, ok
}
