package main

import (
	"fmt"
	"time"

	"github.com/sigmamonster69/CICDpipelineingolang/internal/app"
)

// main prints a short notes-style summary that is easy to validate in CI.
func main() {
	// Reuse the existing app package so the new command stays aligned with the rest of the repo.
	fmt.Println("CICD notes")
	fmt.Println(app.Message())

	// Include a timestamped line so the command feels like a small utility.
	fmt.Printf("generated at %s\n", time.Now().UTC().Format(time.RFC3339))
}
