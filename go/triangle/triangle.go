//Package triangle determines if three length create a equilateral, isosceles or scalene triangle.
package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

// KindFromSides implements the check to determine the kind of triangle.
func KindFromSides(a, b, c float64) Kind {

	if (a == math.Inf(1) || b == math.Inf(1) || c == math.Inf(1)) ||
		(a < 0 || b < 0 || c < 0) ||
		(a == 0 && b == 0 && c == 0) ||
		!inEquality([]float64{a, b, c}) {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}

// Kind is the return type for a type of triangle.
type Kind int

// Constants for the kinds of triangles
const (
	NaT = 0
	Equ = 1
	Iso = 2
	Sca = 3
)

func inEquality(sides []float64) bool {
	sort.Float64s(sides)
	return sides[2] <= sides[1]+sides[0]
}
