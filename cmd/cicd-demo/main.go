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
	// Print the application status message
	// This confirms the binary was built correctly and is running
	fmt.Println(app.Message())
	
	// Generate and display a sample learning report
	// The report demonstrates multiple arithmetic operations from the app package
	// Parameters: name="sam", a=2, b=3
	fmt.Printf("%s\n", app.BuildReport("sam", 2, 3))
}

//hi
