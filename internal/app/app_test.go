package app

import "testing"

// TestMessage verifies that the Message function returns the expected status text.
// This is a basic smoke test to ensure the application initializes correctly.
// The CI pipeline runs this test to confirm the build is functioning.
func TestMessage(t *testing.T) {
	got := Message()
	want := "CI/CD pipeline learning scaffold is ready."

	if got != want {
		t.Fatalf("Message() = %q, want %q", got, want)
	}
}

// TestAdd verifies that the Add function correctly sums two integers.
// This fundamental arithmetic test validates core functionality.
// The CI/CD pipeline uses this as a quality gate before deployment.
func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
	
	// Additional test with negative numbers
	if got := Add(-5, -3); got != -8 {
		t.Fatalf("Add(-5, -3) = %d, want -8", got)
	}
}

// TestMultiply verifies the multiplication operation with various inputs.
// Multiplication is used in calculations throughout the application.
func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{4, 5, 20},
		{0, 100, 0},
		{-3, 7, -21},
		{1, 999, 999},
	}
	
	for _, tt := range tests {
		if got := Multiply(tt.a, tt.b); got != tt.expected {
			t.Fatalf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}

// TestSquare verifies that squaring an integer produces the correct result.
// Square is n multiplied by itself (n²).
func TestSquare(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{6, 36},
		{0, 0},
		{-5, 25},
		{1, 1},
		{10, 100},
	}
	
	for _, tt := range tests {
		if got := Square(tt.input); got != tt.expected {
			t.Fatalf("Square(%d) = %d, want %d", tt.input, got, tt.expected)
		}
	}
}

// TestSubtract verifies the subtraction operation with various test cases.
// Subtraction complements Add for complete arithmetic support.
func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{10, 4, 6},
		{5, 10, -5},
		{0, 0, 0},
		{-5, -3, -2},
	}
	
	for _, tt := range tests {
		if got := Subtract(tt.a, tt.b); got != tt.expected {
			t.Fatalf("Subtract(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}

// TestIsEven verifies the even number detection logic.
// This test covers both even and odd numbers including edge cases.
func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"even number 4", 4, true},
		{"odd number 3", 3, false},
		{"zero is even", 0, true},
		{"negative even", -2, true},
		{"negative odd", -7, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.n); got != tt.want {
				t.Fatalf("IsEven(%d) = %t, want %t", tt.n, got, tt.want)
			}
		})
	}
}

// TestBuildReport verifies the report formatting function.
// BuildReport combines multiple arithmetic operations into a formatted string.
// This test ensures all components are correctly calculated and formatted.
func TestBuildReport(t *testing.T) {
	got := BuildReport("sam", 2, 3)
	want := "hello sam | sum=5 | product=6 | square=4 | difference=-1 | even=true"

	if got != want {
		t.Fatalf("BuildReport(\"sam\", 2, 3) = %q, want %q", got, want)
	}
	
	// Additional test with different values
	got2 := BuildReport("alice", 10, 5)
	want2 := "hello alice | sum=15 | product=50 | square=100 | difference=5 | even=true"
	if got2 != want2 {
		t.Fatalf("BuildReport(\"alice\", 10, 5) = %q, want %q", got2, want2)
	}
}
