package cicdpipelineingolang

import (
	"fmt"
)

// Greeting returns a greeting message for the given name.
func Greeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Factorial calculates the factorial of a non-negative integer.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci returns the nth Fibonacci number.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// ReverseString reverses a string.
func ReverseString(s string) string {
	runes := []rune(s)
	for f, j := 0, len(runes)-1; f < j; f, j = f+1, j-1 {
		runes[f], runes[j] = runes[j], runes[f]
	}
	return string(runes)
}

// IsPalindrome reports whether a string reads the same forwards and backwards.
func IsPalindrome(s string) bool {
	return ReverseString(s) == s
}

// PrintGreeting prints a greeting to stdout.
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
