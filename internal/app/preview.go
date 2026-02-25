package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	inputpkg "github.com/tidjee-dev/scaffoldgen/internal/input"
	"github.com/tidjee-dev/scaffoldgen/internal/model"
	"github.com/tidjee-dev/scaffoldgen/internal/tui"
)

var previewInput string

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview scaffold structure (tree view)",
	Example: `
scaffoldgen preview --in structure.md
scaffoldgen preview --in structure.json
scaffoldgen preview --in structure.yml
`,
	RunE: runPreview,
}

func init() {
	rootCmd.AddCommand(previewCmd)

	previewCmd.Flags().StringVar(&previewInput, "in", "", "Path to structure file (md, json, or yml)")
	_ = previewCmd.MarkFlagRequired("in")
}

func runPreview(cmd *cobra.Command, args []string) error {
	f, err := os.Open(previewInput)
	if err != nil {
		return fmt.Errorf("open input: %w", err)
	}
	defer f.Close()

	ext := strings.ToLower(filepath.Ext(previewInput))
	var root *model.Node

	switch ext {
	case ".json":
		root, err = inputpkg.ParseJSON(f)
		if err != nil {
			return fmt.Errorf("parse json: %w", err)
		}
	case ".yml", ".yaml":
		root, err = inputpkg.ParseYAML(f)
		if err != nil {
			return fmt.Errorf("parse yaml: %w", err)
		}
	case ".md", ".markdown":
		root, err = inputpkg.ParseMarkdown(f)
		if err != nil {
			return fmt.Errorf("parse markdown: %w", err)
		}
	default:
		return fmt.Errorf("unsupported file format: %s (supported: .md, .json, .yml, .yaml)", ext)
	}

	fmt.Printf("%s\n", tui.Header("Preview"))
	tui.RenderTree(root, "", true, "")
	fmt.Println()

	return nil
}
