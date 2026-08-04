package main

import (
	"fmt"

	"github.com/sigmamonster69/CICDpipelineingolang/internal/app"
)

func main() {
	fmt.Println(app.Message())
	fmt.Printf("%s\n", app.BuildReport("sam", 2, 3))
}
