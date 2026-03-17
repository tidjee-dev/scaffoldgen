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

// validateInputFile performs comprehensive validation of the input file
func validateInputFile(path string) error {
	if path == "" {
		return fmt.Errorf("input file path cannot be empty")
	}

	// Check if file exists
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("input file does not exist: %s", path)
		}
		return fmt.Errorf("cannot access input file: %v", err)
	}

	// Check if it's a regular file
	if !info.Mode().IsRegular() {
		return fmt.Errorf("input path must be a regular file, not a directory or special file: %s", path)
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(path))
	supportedExts := map[string]bool{
		".md":       true,
		".json":     true,
		".yml":      true,
		".yaml":     true,
		".markdown": true,
	}

	if !supportedExts[ext] {
		return fmt.Errorf("unsupported file format: %s (supported: .md, .json, .yml, .yaml)", ext)
	}

	// Check if file is readable
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("cannot open input file for reading: %v", err)
	}
	file.Close()

	return nil
}

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
	// Perform comprehensive input validation first
	if err := validateInputFile(validateInput); err != nil {
		return fmt.Errorf("input validation failed: %w", err)
	}

	f, err := os.Open(validateInput)
	if err != nil {
		return fmt.Errorf("open input: %w", err)
	}
	defer f.Close()

	ext := strings.ToLower(filepath.Ext(validateInput))
	var root *model.Node

	switch ext {
	case ".md", ".markdown":
		root, err = inputpkg.ParseMarkdown(f)
	case ".json":
		root, err = inputpkg.ParseJSON(f)
	case ".yml", ".yaml":
		root, err = inputpkg.ParseYAML(f)
	default:
		// This should never happen due to validateInputFile, but keep as safeguard
		return fmt.Errorf("unsupported file format: %s (supported: .md, .json, .yml, .yaml)", ext)
	}

	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	// Validate that root is not nil before proceeding
	if root == nil {
		return fmt.Errorf("parsed structure is empty or invalid")
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
