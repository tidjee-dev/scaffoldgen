package generator

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

// ScanDirectory recursively scans a directory and builds a Node tree
func ScanDirectory(rootPath string) (*model.Node, error) {
	info, err := os.Stat(rootPath)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, os.ErrInvalid
	}

	rootName := filepath.Base(rootPath)
	root := model.NewDir(rootName)

	err = scanDir(rootPath, root)
	return root, err
}

// ScanDirectoryWithIgnores scans a directory respecting ignore patterns
func ScanDirectoryWithIgnores(rootPath string, ignorePatterns []string) (*model.Node, error) {
	root, err := ScanDirectory(rootPath)
	if err != nil {
		return nil, err
	}

	ignoreRules := model.NewIgnoreRules(ignorePatterns)
	filterTree(root, ignoreRules)
	return root, nil
}

func scanDir(dirPath string, parent *model.Node) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// Sort entries by name for deterministic output
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for _, entry := range entries {
		name := entry.Name()

		// Skip hidden files by default
		if strings.HasPrefix(name, ".") {
			continue
		}

		var child *model.Node
		if entry.IsDir() {
			child = model.NewDir(name)
			subPath := filepath.Join(dirPath, name)
			if err := scanDir(subPath, child); err != nil {
				return err
			}
		} else {
			child = model.NewFile(name)
		}

		parent.AddChild(child)
	}

	return nil
}

// filterTree removes nodes that match ignore rules
func filterTree(n *model.Node, ignoreRules *model.IgnoreRules) {
	if n == nil {
		return
	}

	if ignoreRules.ShouldIgnore(n.Name, n.IsDir()) {
		n.Ignore = true
		return
	}

	// Filter children
	filtered := make([]*model.Node, 0, len(n.Children))
	for _, child := range n.Children {
		filterTree(child, ignoreRules)
		if !child.Ignore {
			filtered = append(filtered, child)
		}
	}
	n.Children = filtered
}
