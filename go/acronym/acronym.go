// Package acronym implements a string to acronym conversion
package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Abbreviate implements string parsing rules to create an acronym
func Abbreviate(s string) string {
	var acronym string
	acronym = string(s[0])

	for i, r := range s {
		if (unicode.IsUpper(r) || string(s[i-1]) == " " || string(s[i-1]) == "-") && (i > 0 && !unicode.IsUpper(rune(s[i-1]))) {
			acronym += string(s[i])
		}

	}
	return strings.ToUpper(acronym)
}
