package main

import (
	"fmt"

	"github.com/sigmamonster69/CICDpipelineingolang/internal/app"
)

func main() {
	fmt.Println("CICD pipeline summary")
	fmt.Println(app.Message())
	fmt.Printf("2 + 3 = %d\n", app.Add(2, 3))
	fmt.Printf("4 * 5 = %d\n", app.Multiply(4, 5))
}
