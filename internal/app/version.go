package app

import (
	"fmt"
)

// Version is set during build using ldflags, defaults to current release
var Version = "1.1.3"

func PrintVersion() {
	fmt.Printf("scaffoldgen %s\n", Version)
}
