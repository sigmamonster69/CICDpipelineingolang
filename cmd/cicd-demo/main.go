package main

import (
\t"fmt"

\t"github.com/sigmamonster69/CICDpipelineingolang/internal/app"
)

func main() {
\tfmt.Println(app.Message())
\tfmt.Printf("%s\\n", app.BuildReport("sam", 2, 3))
}
