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
	})

	if len(conflicts) == 0 {
		return nil
	}

	var b strings.Builder

	for _, c := range conflicts {
		b.WriteString(" - ")
		b.WriteString(c)
		b.WriteByte('\n')
	}

	return fmt.Errorf("%s", b.String())
}
