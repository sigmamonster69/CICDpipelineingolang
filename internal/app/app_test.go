package app

import "testing"

func TestMessage(t *testing.T) {
	got := Message()
	want := "CI/CD pipeline learning scaffold is ready."

	if got != want {
		t.Fatalf("Message() = %q, want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
	
	if got := Add(-5, -3); got != -8 {
		t.Fatalf("Add(-5, -3) = %d, want -8", got)
	}
}

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

func TestBuildReport(t *testing.T) {
	got := BuildReport("sam", 2, 3)
	want := "hello sam | sum=5 | product=6 | square=4 | difference=-1 | even=true"

	if got != want {
		t.Fatalf("BuildReport(\"sam\", 2, 3) = %q, want %q", got, want)
	}
	
	got2 := BuildReport("alice", 10, 5)
	want2 := "hello alice | sum=15 | product=50 | square=100 | difference=5 | even=true"
	if got2 != want2 {
		t.Fatalf("BuildReport(\"alice\", 10, 5) = %q, want %q", got2, want2)
	}
}
