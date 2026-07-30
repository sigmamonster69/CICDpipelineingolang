package main

import "testing"

// TestAdd verifies that the Add function correctly sums two integers.
// This test runs automatically in the CI pipeline on every code push.
// It validates the core arithmetic functionality of our application.
// If this test fails, the CI pipeline will block the deployment,
// preventing buggy code from reaching production.
func TestAdd(t *testing.T) {
	// Test case: adding two positive integers
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

// TestSubtract verifies that the Subtract function correctly calculates the difference.
// This test is part of our automated quality gate in the CI/CD pipeline.
// A failure here indicates a regression in basic arithmetic operations.
// The pipeline will halt and report the error to developers immediately.
func TestSubtract(t *testing.T) {
	// Test case: subtracting a smaller number from a larger number
	result := Subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Subtract(10, 4) = %d; expected %d", result, expected)
	}
}

// TestAddNegativeNumbers tests edge cases with negative integer operands.
// Good CI pipelines test both normal cases and edge cases like negative numbers.
// This ensures our functions handle the full range of integer inputs correctly.
// Edge case testing helps catch subtle bugs that might appear in production.
func TestAddNegativeNumbers(t *testing.T) {
	// Test case: adding two negative integers should yield a more negative result
	result := Add(-5, -3)
	expected := -8
	if result != expected {
		t.Errorf("Add(-5, -3) = %d; expected %d", result, expected)
	}
	
	// Test case: adding a positive and negative integer
	result2 := Add(5, -3)
	expected2 := 2
	if result2 != expected2 {
		t.Errorf("Add(5, -3) = %d; expected %d", result2, expected2)
	}
}
