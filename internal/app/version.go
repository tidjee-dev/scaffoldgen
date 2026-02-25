package app

import "fmt"

var Version = "dev" 

func PrintVersion() {
	fmt.Println("scaffoldgen", Version)
}
