package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestAddNegativeNumbers(t *testing.T) {
	result := Add(-5, -3)
	if result != -8 {
		t.Errorf("Expected -8, got %d", result)
	}
}
