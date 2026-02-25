package model

import "fmt"

// PrintTree prints the node tree (debug purpose only).
func PrintTree(n *Node, depth int) {
	if n == nil {
		return
	}

	prefix := ""
	for i := 0; i < depth; i++ {
		prefix += "  "
	}

	kind := "DIR"
	if n.IsFile() {
		kind = "FILE"
	}

	fmt.Printf("%s- %s (%s)\n", prefix, n.Name, kind)

	for _, c := range n.Children {
		PrintTree(c, depth+1)
	}
}
