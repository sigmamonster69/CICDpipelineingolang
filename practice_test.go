package cicdpipelineingolang

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("sam")
	want := "Hello, sam!"
	if got != want {
		t.Fatalf("Greeting() = %q, want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
}

func TestFactorial(t *testing.T) {
	if got := Factorial(5); got != 120 {
		t.Fatalf("Factorial(5) = %d, want 120", got)
	}
}

func TestFibonacci(t *testing.T) {
	if got := Fibonacci(6); got != 8 {
		t.Fatalf("Fibonacci(6) = %d, want 8", got)
	}
}

func TestReverseString(t *testing.T) {
	if got := ReverseString("codex"); got != "xedoc" {
		t.Fatalf("ReverseString() = %q, want %q", got, "xedoc")
	}
}
