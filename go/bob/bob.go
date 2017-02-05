// Package bob implements teenage responses
package bob

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	testVersion      = 2
	period           = 46
	exclamationPoint = 33
	questionMark     = 63
	none             = 0
)

func Hey(s string) string {
	var response string
	var upper bool

	s = strings.TrimSpace(s)
	if s == "" {
		response = "Fine. Be that way!"
	} else {
		endMark := endMark(s)
		letterUpper := hasUpper(s)

		if endMark == 0 {
			upper = allUpper(s)
		} else {
			upper = allUpper(s[0 : len(s)-1])
		}

		fmt.Println(endMark, upper, letterUpper)

		if endMark == period {
			response = "Whatever."
		}

		if endMark == exclamationPoint {
			if upper {
				response = "Whoa, chill out!"
			} else {
				response = "Whatever."
			}
		}

		if endMark == none {
			if upper && letterUpper {
				response = "Whoa, chill out!"
			} else {
				response = "Whatever."
			}
		}

		if endMark == questionMark {
			if upper && letterUpper {
				response = "Whoa, chill out!"
			} else {
				response = "Sure."
			}
		}
	}

	return response
}

func allSpaces(s string) bool {
	as := true
	for _, r := range s {
		if !unicode.IsSpace(r) {
			as = false
		}
	}
	return as
}

func endMark(s string) int {
	endMark := none
	if len(s) > 0 {
		endMark = int(s[len(s)-1])
		if endMark != period && endMark != exclamationPoint && endMark != questionMark {
			endMark = none
		}
	}
	return int(endMark)
}

func hasEndMark(s string) bool {
	e := s[len(s)-1]
	return e == 46

}
func allUpper(s string) bool {
	upper := true
	for _, r := range s {
		if upper {
			fmt.Println(s, r)
			if unicode.IsUpper(r) || unicode.IsSpace(r) || unicode.IsDigit(r) || unicode.IsPunct(r) || unicode.IsMark(r) || unicode.IsControl(r) || unicode.IsSymbol(r) {
				upper = true
			} else {
				upper = false
			}
		}
	}

	return upper
}

func hasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}
