package app

import (
	"fmt"
	"strings"
)

// Message returns the status text displayed by the demo application.
// This function is used as a health check indicator in CI/CD pipelines.
// When the application starts successfully, this message confirms proper initialization.
// Returns:
//   - string: a status message indicating the pipeline scaffold is ready
func Message() string {
	return "CI/CD pipeline learning scaffold is ready."
}

// Add returns the sum of two integers.
// This is a fundamental arithmetic operation used throughout the application.
// The CI pipeline includes automated tests to verify this function's correctness.
// Any regression here would indicate a serious issue with basic operations.
// Parameters:
//   - a: first integer operand
//   - b: second integer operand
//
// Returns:
//   - int: the sum of a and b
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers.
// Multiplication is used in various calculations within the application.
// This function demonstrates how simple operations are tested in CI/CD workflows.
// Parameters:
//   - a: first integer operand (multiplicand)
//   - b: second integer operand (multiplier)
//
// Returns:
//   - int: the product of a and b
func Multiply(a, b int) int {
	return a * b
}

// Square returns n multiplied by itself (n²).
// Squaring is a common mathematical operation used in statistics and geometry.
// This function shows how utility functions are documented and tested.
// Parameters:
//   - n: the integer to square
//
// Returns:
//   - int: n squared (n × n)
func Square(n int) int {
	return n * n
}

// Subtract returns the difference between two integers (a - b).
// Subtraction complements the Add function for complete arithmetic support.
// Both operations are validated by the automated test suite.
// Parameters:
//   - a: minuend (the number from which another is subtracted)
//   - b: subtrahend (the number to subtract)
//
// Returns:
//   - int: the result of subtracting b from a
func Subtract(a, b int) int {
	return a - b
}

// IsEven reports whether an integer n is even (divisible by 2).
// This function demonstrates boolean return types and modulo operations.
// Even/odd checks are common in algorithms and data processing.
// Parameters:
//   - n: the integer to check
//
// Returns:
//   - bool: true if n is even, false if odd
func IsEven(n int) bool {
	return n%2 == 0
}

// BuildReport formats a comprehensive learning report from the basic helper functions.
// This function demonstrates string formatting and composition of multiple operations.
// It's used in the dashboard to show example output from all arithmetic functions.
// The report includes: greeting, sum, product, square, difference, and even check.
// Parameters:
//   - name: the name to include in the greeting
//   - a: first integer for arithmetic operations
//   - b: second integer for arithmetic operations
//
// Returns:
//   - string: a formatted report with pipe-separated values
func BuildReport(name string, a, b int) string {
	// Build a slice of report parts, each showing a different calculation
	parts := []string{
		fmt.Sprintf("hello %s", name),                // Greeting with provided name
		fmt.Sprintf("sum=%d", Add(a, b)),             // Sum of a and b
		fmt.Sprintf("product=%d", Multiply(a, b)),    // Product of a and b
		fmt.Sprintf("square=%d", Square(a)),          // Square of a
		fmt.Sprintf("difference=%d", Subtract(a, b)), // Difference: a - b
		fmt.Sprintf("even=%t", IsEven(a)),            // Whether a is even
	}

	// Join all parts with pipe separator for easy parsing
	return strings.Join(parts, " | ")
}
