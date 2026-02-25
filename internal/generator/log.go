package generator

import "fmt"

func logVerbose(enabled bool, kind, path string) {
	if !enabled {
		return
	}

	fmt.Printf("[scaffoldgen] %-4s %s\n", kind, path)
}
