// Package panagram provides a boolean result of panagram analysis
package pangram

import (
	"strings"
)

const testVersion = 1

// IntSet is a set of integers
type IntSet struct {
	set map[int]bool
}

// NewIntSet instantiates a new set of integers
func NewIntSet() *IntSet {
	return &IntSet{make(map[int]bool)}
}

// IsPanagram implements a string analysis to determine if panagram
func IsPangram(s string) bool {
	panagram := true
	set := NewIntSet()

	if len(s) > 27 {
		s = strings.ToLower(s)
		for _, l := range s {
			set.Add(int(l))
		}

		for i := 97; i < 123 && panagram; i++ {
			if !set.Contains(i) {
				panagram = false
			}
		}
	} else {
		panagram = false
	}
	return panagram
}

// Add adds an integer to the set
func (set *IntSet) Add(i int) {
	set.set[i] = true
}

// Contains returns true if the integer is in the set, false otherwise
func (set *IntSet) Contains(i int) bool {
	return set.set[i]
}
