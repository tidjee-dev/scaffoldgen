package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tidjee-dev/scaffoldgen/internal/generator"
	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

func ValidateFilesystemConflicts(root *model.Node) error {
	var conflicts []string

	generator.Walk(root, "", func(e generator.Event) {
		path := filepath.FromSlash(e.Path)

		info, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				return
			}
			conflicts = append(conflicts, fmt.Sprintf("stat %s: %v", e.Path, err))
			return
		}

		// Check for file/directory type conflicts
		switch e.Kind {
		case generator.EventDir:
			if !info.IsDir() {
				conflicts = append(conflicts,
					fmt.Sprintf("%s exists as FILE but expected DIRECTORY", e.Path))
			}

		case generator.EventFile:
			if info.IsDir() {
				conflicts = append(conflicts,
					fmt.Sprintf("%s exists as DIRECTORY but expected FILE", e.Path))
			}
		}

		// Check for symlink conflicts
		if info.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(path)
			if err != nil {
				conflicts = append(conflicts,
					fmt.Sprintf("%s is a broken symlink: %v", e.Path, err))
			} else {
				conflicts = append(conflicts,
					fmt.Sprintf("%s is a symlink pointing to: %s", e.Path, target))
			}
		}

		// Check for permission issues
		switch e.Kind {
		case generator.EventFile:
			// Check if we can write to the file location
			if info.Mode()&0200 == 0 { // No write permission for owner
				conflicts = append(conflicts,
					fmt.Sprintf("%s exists but lacks write permissions", e.Path))
			}

		case generator.EventDir:
			// Check if we can create files in the directory
			if info.Mode()&0300 == 0 { // No write/execute permission for owner
				conflicts = append(conflicts,
					fmt.Sprintf("%s exists but lacks write/execute permissions", e.Path))
			}
		}
	})

	if len(conflicts) == 0 {
		return nil
	}

	// Optimize string building with strings.Join
	conflictLines := make([]string, len(conflicts))
	for i, c := range conflicts {
		conflictLines[i] = " - " + c
	}

	return fmt.Errorf("%s", strings.Join(conflictLines, "\n"))
}
