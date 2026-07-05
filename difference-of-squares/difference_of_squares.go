// Package differenceofsquares provides functions to calculate the difference between
// the square of the sum and the sum of the squares of the first N natural numbers.
//
// This problem demonstrates the importance of choosing efficient algorithms:
// - Naive approach: O(n) time, O(1) space
// - Mathematical approach: O(1) time, O(1) space using closed-form formulas
//
// Mathematical formulas used:
//
//	Sum of first n natural numbers: n(n+1)/2
//	Sum of squares of first n natural numbers: n(n+1)(2n+1)/6
package differenceofsquares

// SquareOfSum calculates the square of the sum of the first n natural numbers.
//
// Uses the closed-form formula: (n(n+1)/2)²
//
// Parameters:
//   - n: a positive integer
//
// Returns:
//   - int: the square of the sum
//
// Example:
//
//	SquareOfSum(10) // returns 3025
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first n natural numbers.
//
// Uses the closed-form formula: n(n+1)(2n+1)/6
//
// Parameters:
//   - n: a positive integer
//
// Returns:
//   - int: the sum of squares
//
// Example:
//
//	SumOfSquares(10) // returns 385
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between the square of the sum and the
// sum of the squares of the first n natural numbers.
//
// The function uses an O(1) mathematical approach:
//
//	squareOfSum = (sum of numbers)²
//	sumOfSquares = sum of (numbers²)
//	result = squareOfSum - sumOfSquares
//
// This is significantly more efficient than iterating through all numbers,
// especially for large values of n.
//
// Parameters:
//   - n: a positive integer (number of natural numbers to consider)
//
// Returns:
//   - int: the difference (square of sum - sum of squares)
//
// Example:
//
//	Difference(10) // returns 2640
//	Difference(1)  // returns 0 (1² - 1² = 0)
//	Difference(2)  // returns 4 ((1+2)² - (1²+2²) = 9-5 = 4)
//
// Complexity:
//   - Time: O(1) - constant time regardless of input size
//   - Space: O(1) - uses only a few variables
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
