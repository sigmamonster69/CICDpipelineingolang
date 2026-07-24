package demo

import "testing"

func TestMessage(t *testing.T) {
\tgot := Message()
\twant := "CI/CD pipeline learning scaffold is ready."

\tif got != want {
\t\tt.Fatalf("Message() = %q, want %q", got, want)
\t}
}
