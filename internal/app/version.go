package app

import "fmt"

var Version = "1.1.0"

func PrintVersion() {
	fmt.Println("scaffoldgen", Version)
}
