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

var (
	validateInput string
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check structure against existing directory",
	Long: `Validate a structure file against the current filesystem to detect conflicts.
Useful for checking what would be overwritten or conflicted before generation.`,
	Example: `
scaffoldgen validate --in structure.md
scaffoldgen validate --in structure.yml
`,
	RunE: runValidate,
}

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringVar(&validateInput, "in", "", "Path to structure file (md, json, or yml)")
	_ = validateCmd.MarkFlagRequired("in")
}

func runValidate(cmd *cobra.Command, args []string) error {
	f, err := os.Open(validateInput)
	if err != nil {
		return fmt.Errorf("open input: %w", err)
	}
	defer f.Close()

	ext := strings.ToLower(filepath.Ext(validateInput))
	var root *model.Node

	switch ext {
	case ".md":
		root, err = inputpkg.ParseMarkdown(f)
	case ".json":
		root, err = inputpkg.ParseJSON(f)
	case ".yml", ".yaml":
		root, err = inputpkg.ParseYAML(f)
	default:
		return fmt.Errorf("unsupported file format: %s (supported: .md, .json, .yml, .yaml)", ext)
	}

	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	// Validate filesystem conflicts
	if err := ValidateFilesystemConflicts(root); err != nil {
		fmt.Println(tui.Error("❌ Validation failed:"))
		fmt.Println(err.Error())
		return fmt.Errorf("validation failed")
	}

	fmt.Println(tui.Success("✅ Validation passed - no conflicts detected"))
	return nil
}
