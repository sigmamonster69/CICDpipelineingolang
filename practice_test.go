package cicdpipelineingolang

import "testing"

// TestGreeting verifies that the Greeting function returns the correct formatted message.
// This test ensures string formatting works as expected with different input names.
// The CI pipeline runs this test automatically to catch any regressions in string handling.
func TestGreeting(t *testing.T) {
	got := Greeting("sam")
	want := "Hello, sam!"
	if got != want {
		t.Fatalf("Greeting(\"sam\") = %q, want %q", got, want)
	}
	
	// Additional test case: empty name
	got2 := Greeting("")
	want2 := "Hello, !"
	if got2 != want2 {
		t.Fatalf("Greeting(\"\") = %q, want %q", got2, want2)
	}
}

// TestAdd verifies that the Add function correctly sums two integers.
// This is a fundamental test that validates basic arithmetic operations.
// The CI/CD pipeline uses this test as a quality gate before deployment.
func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
	
	// Test with negative numbers
	if got := Add(-5, -3); got != -8 {
		t.Fatalf("Add(-5, -3) = %d, want -8", got)
	}
}

// TestFactorial verifies the factorial calculation for various inputs.
// Factorial is a recursive function, so we test both base cases and recursive cases.
// This test helps ensure the recursion logic works correctly without stack overflow.
func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"factorial of 0", 0, 1},    // Base case: 0! = 1
		{"factorial of 1", 1, 1},    // Base case: 1! = 1
		{"factorial of 5", 5, 120},  // Recursive case: 5! = 120
		{"factorial of 6", 6, 720},  // Larger recursive case
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.input); got != tt.expected {
				t.Fatalf("Factorial(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}

// TestFibonacci verifies the Fibonacci sequence calculation.
// The Fibonacci sequence is commonly used to demonstrate recursion.
// We test multiple positions in the sequence to validate the recursive logic.
func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"fibonacci(0)", 0, 0},  // Base case
		{"fibonacci(1)", 1, 1},  // Base case
		{"fibonacci(6)", 6, 8},  // F(6) = 8
		{"fibonacci(10)", 10, 55}, // F(10) = 55
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.input); got != tt.expected {
				t.Fatalf("Fibonacci(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}

// TestReverseString verifies that strings are correctly reversed.
// This test includes ASCII and Unicode characters to ensure proper rune handling.
// String manipulation is common in text processing applications.
func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"reverse codex", "codex", "xedoc"},
		{"reverse hello", "hello", "olleh"},
		{"reverse empty", "", ""},
		{"reverse single char", "a", "a"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.input); got != tt.expected {
				t.Fatalf("ReverseString(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// TestIsPalindrome verifies palindrome detection for various strings.
// A palindrome reads the same forwards and backwards.
// This test covers both positive and negative cases to ensure accuracy.
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"level is palindrome", "level", true},
		{"radar is palindrome", "radar", true},
		{"hello is not palindrome", "hello", false},
		{"empty string is palindrome", "", true},
		{"single char is palindrome", "a", true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.input); got != tt.expected {
				t.Fatalf("IsPalindrome(%q) = %t, want %t", tt.input, got, tt.expected)
			}
		})
	}
}

// TestMax verifies that the Max function returns the larger of two integers.
// This utility function is tested with various combinations including negatives.
func TestMax(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 3},
		{9, 4, 9},
		{-5, -3, -3},
		{0, 0, 0},
		{7, 7, 7},
	}
	
	for _, tt := range tests {
		if got := Max(tt.a, tt.b); got != tt.expected {
			t.Fatalf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}

// TestMin verifies that the Min function returns the smaller of two integers.
// This is the counterpart to Max and is tested with similar edge cases.
func TestMin(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 2},
		{9, 4, 4},
		{-5, -3, -5},
		{0, 0, 0},
		{7, 7, 7},
	}
	
	for _, tt := range tests {
		if got := Min(tt.a, tt.b); got != tt.expected {
			t.Fatalf("Min(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}

// TestAbs verifies the absolute value calculation for positive, negative, and zero.
// Absolute value is the distance from zero, always non-negative.
func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"abs of -7", -7, 7},
		{"abs of 5", 5, 5},
		{"abs of 0", 0, 0},
		{"abs of -100", -100, 100},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.input); got != tt.expected {
				t.Fatalf("Abs(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
