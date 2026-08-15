package cicdpipelineingolang

import (
	"fmt"
)

// Greeting returns a personalized greeting message for the given name.
func Greeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Factorial calculates the factorial of a non-negative integer n (n!).
// For example: 5! = 5 × 4 × 3 × 2 × 1 = 120
// Note: For large values of n, this may cause stack overflow.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci returns the nth Fibonacci number using recursive calculation.
// The Fibonacci sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21...
// Note: This implementation is educational but inefficient for large n.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// ReverseString reverses a string by converting it to runes and swapping characters.
// This function handles Unicode characters correctly.
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome reports whether a string reads the same forwards and backwards.
// This check is case-sensitive and includes spaces/punctuation.
func IsPalindrome(s string) bool {
	return ReverseString(s) == s
}

// PrintGreeting prints "Hello, World!" to stdout.
func PrintGreeting() {
	fmt.Println("Hello, World!")
}

// Max returns the larger of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Abs returns the absolute value of an integer.
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// IsOdd reports whether n is odd.
func IsOdd(n int) bool {
	return n%2 != 0
}
