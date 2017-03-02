// Package secret implements a secret handshake for a code
package secret

import (
	"fmt"
)

const (
	testVersion = 1
	runeOne     = 49
)

var ops = map[int]string{
	0: "wink",
	1: "double blink",
	2: "close your eyes",
	3: "jump",
}

// Handshake shows the operations for the secret handshake
func Handshake(code uint) []string {
	var handshake []string

	b := fmt.Sprintf("%b", code)
	b = reverse(b)

	fmt.Println(code | 0x02)

	for index, r := range b {
		if index == 4 {
			handshake = reverseStrArray(handshake)
			return handshake // no need to evaluate greater than 32
		} else if index < 4 {
			if r == runeOne {
				handshake = append(handshake, ops[index])
			}
		}
	}

	return handshake
}

func reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func reverseStrArray(s []string) []string {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
	return s
}
