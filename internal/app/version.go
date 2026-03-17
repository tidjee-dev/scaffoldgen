package app

import "fmt"

var Version = "1.0.2"

func PrintVersion() {
	fmt.Println("scaffoldgen", Version)
}
