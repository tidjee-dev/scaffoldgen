package app

import (
	"fmt"
)

// Version is set during build using ldflags
var Version = "dev"

func PrintVersion() {
	fmt.Printf("scaffoldgen %s\n", Version)
}
