package main

import (
	"fmt"

	"github.com/sigmamonster69/CICDpipelineingolang/internal/app"
)

// main is the entry point for the cicd-demo command-line application.
// This demo showcases basic CI/CD pipeline concepts with a simple Go application.
// When executed, it displays a status message and a formatted learning report.
// The CI pipeline automatically builds and tests this binary on every commit.
func main() {
	fmt.Println(app.Message())
	fmt.Printf("%s\n", app.BuildReport("sam", 2, 3))
}
