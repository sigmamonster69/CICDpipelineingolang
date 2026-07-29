package main

import "testing"

// TestAdd verifies that the Add function works correctly
// This test runs automatically in the CI pipeline on every code push
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

// TestSubtract verifies that the Subtract function works correctly
// If this test fails, the CI pipeline will block the deployment
func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

// TestAddNegativeNumbers tests edge cases with negative numbers
// Good CI pipelines test both normal and edge cases
func TestAddNegativeNumbers(t *testing.T) {
	result := Add(-5, -3)
	if result != -8 {
		t.Errorf("Expected -8, got %d", result)
	}
}
