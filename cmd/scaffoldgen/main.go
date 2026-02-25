package main

import (
	"fmt"
	"os"

	"github.com/tidjee-dev/scaffoldgen/internal/app"
)

func main() {
	if err := app.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
