package app

import "testing"

func TestAdd_Subtests(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -5, -3, -8},
		{"mixed signs", -5, 10, 5},
		{"zero and positive", 0, 100, 100},
		{"zero and negative", 0, -50, -50},
		{"both zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMultiply_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"multiply by zero", 100, 0, 0},
		{"multiply by one", 42, 1, 42},
		{"multiply by negative one", 42, -1, -42},
		{"large numbers", 1000, 1000, 1000000},
		{"both negative", -10, -10, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSquare_BoundaryValues(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"square of 1", 1, 1},
		{"square of -1", -1, 1},
		{"square of max int32", 46340, 2147395600},
		{"square of large number", 1000, 1000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Square(tt.input)
			if got != tt.expected {
				t.Errorf("Square(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestIsEven_Comprehensive(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"smallest even", 2, true},
		{"smallest odd", 1, false},
		{"large even", 1000000, true},
		{"large odd", 999999, false},
		{"negative even", -100, true},
		{"negative odd", -99, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.n); got != tt.want {
				t.Errorf("IsEven(%d) = %t, want %t", tt.n, got, tt.want)
			}
		})
	}
}

func TestSubtract_VariousScenarios(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"subtract smaller from larger", 100, 50, 50},
		{"subtract larger from smaller", 50, 100, -50},
		{"subtract from zero", 0, 50, -50},
		{"subtract zero", 50, 0, 50},
		{"subtract negative", 10, -5, 15},
		{"both negative", -10, -5, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestBuildReport_EmptyName(t *testing.T) {
	got := BuildReport("", 5, 5)
	want := "hello  | sum=10 | product=25 | square=25 | difference=0 | even=false"

	if got != want {
		t.Errorf("BuildReport(\"\", 5, 5) = %q, want %q", got, want)
	}
}

func TestBuildReport_LargeNumbers(t *testing.T) {
	got := BuildReport("test", 1000, 2000)
	want := "hello test | sum=3000 | product=2000000 | square=1000000 | difference=-1000 | even=true"

	if got != want {
		t.Errorf("BuildReport(\"test\", 1000, 2000) = %q, want %q", got, want)
	}
}

func TestBuildReport_NegativeNumbers(t *testing.T) {
	got := BuildReport("neg", -5, 3)
	want := "hello neg | sum=-2 | product=-15 | square=25 | difference=-8 | even=false"

	if got != want {
		t.Errorf("BuildReport(\"neg\", -5, 3) = %q, want %q", got, want)
	}
}
