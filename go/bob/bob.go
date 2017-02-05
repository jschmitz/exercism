// Package bob implements teenage responses
package bob

import (
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

// Hey responsds to a conversation like a teenager
func Hey(s string) string {
	var response string
	var upper, letterUpper bool

	s = strings.TrimSpace(s)
	if s == "" {
		response = "Fine. Be that way!"
	} else {
		endMark := endMark(s)
		upper, letterUpper = allUpper(s)

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

func allUpper(s string) (bool, bool) {
	upper := true
	upperLetter := false
	for _, r := range s {
		// check for an upper case letter
		if unicode.IsUpper(r) {
			upperLetter = true
		}

		//Check all upper case letters with special char exceptions
		if upper {
			if unicode.IsUpper(r) || unicode.IsSpace(r) || unicode.IsDigit(r) || unicode.IsPunct(r) || unicode.IsMark(r) || unicode.IsControl(r) || unicode.IsSymbol(r) {
				upper = true
			} else {
				upper = false
			}
		}
	}

	return upper, upperLetter
}
