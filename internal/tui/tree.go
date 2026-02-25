package tui

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tidjee-dev/scaffoldgen/internal/model"

)

func RenderTree(n *model.Node, prefix string, last bool, currentPath string) {
	if n == nil {
		return
	}

	fullPath := n.Name
	if currentPath != "" {
		fullPath = filepath.Join(currentPath, n.Name)
	}

	line := treeLine(prefix, last, renderNodeLabel(n, fullPath))
	fmt.Println(line)

	childPrefix := prefix
	if last {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	var visible []*model.Node
	for _, c := range n.Children {
		if !c.Ignore {
			visible = append(visible, c)
		}
	}

	for i, c := range n.Children {
		RenderTree(c, childPrefix, i == len(n.Children)-1, fullPath)
	}
}

func treeLine(prefix string, last bool, label string) string {
	if prefix == "" {
		return label
	}

	if last {
		return prefix + "└── " + label
	}
	return prefix + "├── " + label
}

func renderNodeLabel(n *model.Node, fullPath string) string {
	if n.Ignore {
		return Ignored(n.Name)
	}

	_, err := os.Stat(filepath.FromSlash(fullPath))

	if err == nil {
		return Exists(n.Name)
	}
	if os.IsNotExist(err) {
		return Added(n.Name)
	}
	return Info("+ " + n.Name)
}
