package diffsquares

// SquareOfSums sums the first N natural numbers then squares the result
func SquareOfSums(x int) int {
	sum := 0
	for i := 0; i <= x; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares sums the squares of the first N natural numbers
func SumOfSquares(x int) int {
	sum := 0
	for i := 0; i <= x; i++ {
		sum += i * i
	}
	return sum
}

// Difference subtraces SumOfSquares from SquareOfSums
func Difference(x int) int {
	return SquareOfSums(x) - SumOfSquares(x)
}
