package app

import "testing"

func TestMessage(t *testing.T) {
\tgot := Message()
\twant := "CI/CD pipeline learning scaffold is ready."

\tif got != want {
\t\tt.Fatalf("Message() = %q, want %q", got, want)
\t}
}

func TestAdd(t *testing.T) {
\tif got := Add(2, 3); got != 5 {
\t\tt.Fatalf("Add(2, 3) = %d, want 5", got)
\t}
}

func TestMultiply(t *testing.T) {
\tif got := Multiply(4, 5); got != 20 {
\t\tt.Fatalf("Multiply(4, 5) = %d, want 20", got)
\t}
}

func TestIsEven(t *testing.T) {
\ttests := []struct {
\t\tname string
\t\tn    int
\t\twant bool
\t}{
\t\t{name: "even", n: 4, want: true},
\t\t{name: "odd", n: 3, want: false},
\t}

\tfor _, tt := range tests {
\t\tt.Run(tt.name, func(t *testing.T) {
\t\t\tif got := IsEven(tt.n); got != tt.want {
\t\t\t\tt.Fatalf("IsEven(%d) = %t, want %t", tt.n, got, tt.want)
\t\t\t}
\t\t})
\t}
}

func TestBuildReport(t *testing.T) {
\tgot := BuildReport("sam", 2, 3)
\twant := "hello sam | sum=5 | product=6 | even=true"

\tif got != want {
\t\tt.Fatalf("BuildReport() = %q, want %q", got, want)
\t}
}
