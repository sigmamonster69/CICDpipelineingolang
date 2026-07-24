package app

import (
	"fmt"
	"strings"
)

// Message returns the text the demo app prints.
func Message() string {
	return "CI/CD pipeline learning scaffold is ready."
}

// Add returns the sum of two numbers.
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two numbers.
func Multiply(a, b int) int {
	return a * b
}

// Subtract returns the difference between two numbers.
func Subtract(a, b int) int {
	return a - b
}

// IsEven reports whether n is even.
func IsEven(n int) bool {
	return n%2 == 0
}

// BuildReport formats a short learning report from the basic helpers.
func BuildReport(name string, a, b int) string {
	parts := []string{
		fmt.Sprintf("hello %s", name),
		fmt.Sprintf("sum=%d", Add(a, b)),
		fmt.Sprintf("product=%d", Multiply(a, b)),
		fmt.Sprintf("difference=%d", Subtract(a, b)),
		fmt.Sprintf("even=%t", IsEven(a)),
	}

	return strings.Join(parts, " | ")
}
