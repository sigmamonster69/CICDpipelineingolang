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

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("level") {
		t.Fatal("IsPalindrome(level) = false, want true")
	}
	if IsPalindrome("hello") {
		t.Fatal("IsPalindrome(hello) = true, want false")
	}
}

func TestMax(t *testing.T) {
	if got := Max(2, 3); got != 3 {
		t.Fatalf("Max(2, 3) = %d, want 3", got)
	}
	if got := Max(9, 4); got != 9 {
		t.Fatalf("Max(9, 4) = %d, want 9", got)
	}
}

func TestMin(t *testing.T) {
	if got := Min(2, 3); got != 2 {
		t.Fatalf("Min(2, 3) = %d, want 2", got)
	}
	if got := Min(9, 4); got != 4 {
		t.Fatalf("Min(9, 4) = %d, want 4", got)
	}
}
