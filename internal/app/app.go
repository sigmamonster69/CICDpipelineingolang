package app

import (
\t"fmt"
\t"strings"
)

// Message returns the text the demo app prints.
func Message() string {
\treturn "CI/CD pipeline learning scaffold is ready."
}

// Add returns the sum of two numbers.
func Add(a, b int) int {
\treturn a + b
}

// Multiply returns the product of two numbers.
func Multiply(a, b int) int {
\treturn a * b
}

// IsEven reports whether n is even.
func IsEven(n int) bool {
\treturn n%2 == 0
}

// BuildReport formats a short learning report from the basic helpers.
func BuildReport(name string, a, b int) string {
\tparts := []string{
\t\tfmt.Sprintf("hello %s", name),
\t\tfmt.Sprintf("sum=%d", Add(a, b)),
\t\tfmt.Sprintf("product=%d", Multiply(a, b)),
\t\tfmt.Sprintf("even=%t", IsEven(a)),
\t}

\treturn strings.Join(parts, " | ")
}
