package app

import (
	"fmt"
)

// Version will be set by build flags during compilation
var Version = "1.1.0" // Set during build

func PrintVersion() {
	fmt.Printf("scaffoldgen %s\n", Version)
}
