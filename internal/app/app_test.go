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
}

func TestMultiply(t *testing.T) {
	if got := Multiply(4, 5); got != 20 {
		t.Fatalf("Multiply(4, 5) = %d, want 20", got)
	}
}

func TestSquare(t *testing.T) {
	if got := Square(6); got != 36 {
		t.Fatalf("Square(6) = %d, want 36", got)
	}
}

func TestSubtract(t *testing.T) {
	if got := Subtract(10, 4); got != 6 {
		t.Fatalf("Subtract(10, 4) = %d, want 6", got)
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "even", n: 4, want: true},
		{name: "odd", n: 3, want: false},
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
		t.Fatalf("BuildReport() = %q, want %q", got, want)
	}
}
