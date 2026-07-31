package cicdpipelineingolang

import (
	"fmt"
)

// Greeting returns a personalized greeting message for the given name.
// This function demonstrates basic string formatting in Go.
// It's commonly used as a simple example in CI/CD pipelines to show
// how automated tests validate string manipulation functions.
// Parameters:
//   - name: the name to include in the greeting (can be any string)
// Returns:
//   - string: a formatted greeting like "Hello, {name}!"
func Greeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Add returns the sum of two integers.
// This is a fundamental arithmetic function used to demonstrate
// how CI/CD pipelines catch regressions in basic operations.
// The function is tested by our automated test suite on every commit.
// Parameters:
//   - a: first integer operand
//   - b: second integer operand
// Returns:
//   - int: the sum of a and b
func Add(a, b int) int {
	return a + b
}

// Factorial calculates the factorial of a non-negative integer n (n!).
// Factorial is the product of all positive integers less than or equal to n.
// For example: 5! = 5 × 4 × 3 × 2 × 1 = 120
// This recursive implementation demonstrates algorithmic thinking in Go.
// Note: For large values of n, this may cause stack overflow.
// Parameters:
//   - n: a non-negative integer (n >= 0)
// Returns:
//   - int: the factorial of n (returns 1 for n <= 1)
func Factorial(n int) int {
	// Base case: factorial of 0 and 1 is defined as 1
	if n <= 1 {
		return 1
	}
	// Recursive case: n! = n × (n-1)!
	return n * Factorial(n-1)
}

// Fibonacci returns the nth Fibonacci number using recursive calculation.
// The Fibonacci sequence starts with 0 and 1, and each subsequent number
// is the sum of the two preceding ones: 0, 1, 1, 2, 3, 5, 8, 13, 21...
// This implementation is educational but inefficient for large n.
// In production, consider using memoization or iterative approaches.
// Parameters:
//   - n: the position in the Fibonacci sequence (0-indexed)
// Returns:
//   - int: the nth Fibonacci number (returns n for n <= 1)
func Fibonacci(n int) int {
	// Base cases: F(0) = 0, F(1) = 1
	if n <= 1 {
		return n
	}
	// Recursive case: F(n) = F(n-1) + F(n-2)
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// ReverseString reverses a string by converting it to runes and swapping characters.
// This function handles Unicode characters correctly by using rune slices.
// It demonstrates string manipulation techniques in Go, which treats strings
// as immutable byte sequences. Converting to []rune ensures proper handling
// of multi-byte UTF-8 characters.
// Parameters:
//   - s: the input string to reverse
// Returns:
//   - string: the reversed string (empty string if input is empty)
func ReverseString(s string) string {
	// Convert string to rune slice to handle Unicode characters properly
	runes := []rune(s)
	// Swap characters from both ends moving toward the center
	for f, j := 0, len(runes)-1; f < j; f, j = f+1, j-1 {
		runes[f], runes[j] = runes[j], runes[f]
	}
	return string(runes)
}

// IsPalindrome reports whether a string reads the same forwards and backwards.
// This function leverages ReverseString to perform the check efficiently.
// Palindrome detection is a common interview question and testing exercise.
// Note: This is case-sensitive and includes spaces/punctuation.
// Examples: "level", "radar", "hello" (false)
// Parameters:
//   - s: the string to check for palindrome property
// Returns:
//   - bool: true if the string is a palindrome, false otherwise
func IsPalindrome(s string) bool {
	// A string is a palindrome if it equals its reverse
	return ReverseString(s) == s
}

// PrintGreeting prints a simple "Hello, World!" message to stdout.
// This is a classic introductory function demonstrating I/O operations in Go.
// It's often used as a smoke test to verify the application runs correctly.
// In CI/CD contexts, this might be captured and validated in integration tests.
func PrintGreeting() {
	fmt.Println("Hello, World!")
}

// Max returns the larger of two integers.
// This utility function demonstrates conditional logic and comparison operations.
// It's commonly used in algorithms and data processing tasks.
// The function is tested with various input combinations in our test suite.
// Parameters:
//   - a: first integer to compare
//   - b: second integer to compare
// Returns:
//   - int: the greater value between a and b (returns a if equal)
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two integers.
// This utility function is the counterpart to Max, returning the minimum value.
// Both Min and Max are fundamental building blocks for sorting and selection algorithms.
// Our CI pipeline tests these functions with edge cases like negative numbers.
// Parameters:
//   - a: first integer to compare
//   - b: second integer to compare
// Returns:
//   - int: the lesser value between a and b (returns a if equal)
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Abs returns the absolute value of an integer.
// The absolute value is the non-negative value of a number without regard to sign.
// For example: Abs(-7) = 7, Abs(5) = 5, Abs(0) = 0
// This function demonstrates sign checking and negation operations in Go.
// Parameters:
//   - n: the integer to get the absolute value of
// Returns:
//   - int: the absolute (non-negative) value of n
func Abs(n int) int {
	// If n is negative, return its negation to make it positive
	if n < 0 {
		return -n
	}
	return n
}

// IsOdd reports whether n is odd.
func IsOdd(n int) bool {
	return n%2 != 0
}
