package diffsquares

func SquareOfSumsNaive(in int) int {
	var sum int
	for i := 1; i <= in; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquaresNaive(in int) int {
	var sum int
	for i := 1; i <= in; i++ {
		sum += i * i
	}
	return sum
}

func DifferenceNaive(in int) int {
	return SquareOfSumsNaive(in) - SumOfSquaresNaive(in)
}

// ====

func SquareOfSums(n int) int {
	s := (n * (n + 1)) / 2
	return s * s
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
