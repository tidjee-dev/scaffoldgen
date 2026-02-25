package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tidjee-dev/scaffoldgen/internal/generator"
	"github.com/tidjee-dev/scaffoldgen/internal/tui"

)

var (
	reverseInput  string
	reverseFormat string
	reverseOut    string
)

var reverseCmd = &cobra.Command{
	Use:   "reverse",
	Short: "Scan directory and generate structure file",
	Long: `Scan an existing directory structure and generate a structure file.
Useful for documenting existing project layouts.`,
	Example: `
scaffoldgen reverse --in ./backend --format md --out structure.md
scaffoldgen reverse --in ./src --format json > structure.json
`,
	RunE: runReverse,
}

func init() {
	rootCmd.AddCommand(reverseCmd)

	reverseCmd.Flags().StringVar(&reverseInput, "in", "", "Path to directory to scan")
	reverseCmd.Flags().StringVar(&reverseFormat, "format", "md", "Output format: md, json, or yml")
	reverseCmd.Flags().StringVar(&reverseOut, "out", "", "Output file (default: stdout)")

	_ = reverseCmd.MarkFlagRequired("in")
}

func runReverse(cmd *cobra.Command, args []string) error {
	// Validate directory
	info, err := os.Stat(reverseInput)
	if err != nil {
		return fmt.Errorf("read directory: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("input must be a directory")
	}

	// Scan directory
	root, err := generator.ScanDirectory(reverseInput)
	if err != nil {
		return fmt.Errorf("scan directory: %w", err)
	}

	// Export in requested format
	format := strings.ToLower(reverseFormat)
	var output string

	switch format {
	case "md", "markdown":
		output = generator.ExportMarkdown(root)

	case "json":
		output, err = generator.ExportJSON(root)
		if err != nil {
			return fmt.Errorf("export json: %w", err)
		}

	case "yml", "yaml":
		output, err = generator.ExportYAML(root)
		if err != nil {
			return fmt.Errorf("export yaml: %w", err)
		}

	default:
		return fmt.Errorf("unsupported format: %s (supported: md, json, yml)", format)
	}

	// Write output
	if reverseOut != "" {
		if err := os.WriteFile(reverseOut, []byte(output), 0o644); err != nil {
			return fmt.Errorf("write file: %w", err)
		}
		fmt.Println(tui.Success(fmt.Sprintf("Generated: %s", reverseOut)))
	} else {
		fmt.Print(output)
	}

	return nil
}
