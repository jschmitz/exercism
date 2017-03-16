// Package accumulate implements an operation on each member of a collection
package accumulate

const testVersion = 1

// Accumulate implements an operation per collection member and returns the result
func Accumulate(c []string, m func(string) string) []string {
	var conversions []string
	for _, s := range c {
		conversions = append(conversions, m(s))
	}
	return conversions
}
