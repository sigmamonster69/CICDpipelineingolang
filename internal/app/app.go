package app

import (
	"fmt"
	"strings"
)

// Message returns the status text displayed by the demo application.
func Message() string {
	return "CI/CD pipeline learning scaffold is ready."
}

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
	return a * b
}

// Square returns n multiplied by itself (n²).
func Square(n int) int {
	return n * n
}

// Subtract returns the difference between two integers (a - b).
func Subtract(a, b int) int {
	return a - b
}

// IsEven reports whether an integer n is even (divisible by 2).
func IsEven(n int) bool {
	return n%2 == 0
}

// BuildReport formats a comprehensive learning report from the basic helper functions.
func BuildReport(name string, a, b int) string {
	parts := []string{
		fmt.Sprintf("hello %s", name),
		fmt.Sprintf("sum=%d", Add(a, b)),
		fmt.Sprintf("product=%d", Multiply(a, b)),
		fmt.Sprintf("square=%d", Square(a)),
		fmt.Sprintf("difference=%d", Subtract(a, b)),
		fmt.Sprintf("even=%t", IsEven(a)),
	}

	return strings.Join(parts, " | ")
}
