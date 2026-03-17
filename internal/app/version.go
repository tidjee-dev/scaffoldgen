package app

import (
	"fmt"
)

// Version is set during build using ldflags
var Version = "1.1.3"

func PrintVersion() {
	fmt.Printf("scaffoldgen %s\n", Version)
}
