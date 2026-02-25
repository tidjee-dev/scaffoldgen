package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tidjee-dev/scaffoldgen/internal/generator"
	inputpkg "github.com/tidjee-dev/scaffoldgen/internal/input"
	"github.com/tidjee-dev/scaffoldgen/internal/model"
	"github.com/tidjee-dev/scaffoldgen/internal/tui"
)

var (
	input   string
	shell   string
	outDir  string
	dryRun  bool
	verbose bool
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate scaffold scripts from markdown structure",
	RunE:  runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVar(&input, "in", "", "Path to structure file (md, json, or yml)")
	generateCmd.Flags().StringVar(&shell, "shell", "sh", "sh|ps1|both")
	generateCmd.Flags().StringVar(&outDir, "out", ".", "Output directory")
	generateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print script without writing file")
	generateCmd.Flags().BoolVar(&verbose, "verbose", false, "Verbose output")

	_ = generateCmd.MarkFlagRequired("in")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	f, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("open input: %w", err)
	}
	defer f.Close()

	ext := strings.ToLower(filepath.Ext(input))
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

	if err := ValidateFilesystemConflicts(root); err != nil {
		return fmt.Errorf("%s", tui.Error(err.Error()))
	}

	shellMode := strings.ToLower(shell)
	if shellMode != "sh" && shellMode != "ps1" && shellMode != "both" {
		return fmt.Errorf("invalid --shell value: %s", shellMode)
	}

	var shScript, ps1Script string

	if shellMode == "sh" || shellMode == "both" {
		shScript, err = generator.GenerateSH(root, verbose)
		if err != nil {
			return err
		}
	}

	if shellMode == "ps1" || shellMode == "both" {
		ps1Script, err = generator.GeneratePS1(root, verbose)
		if err != nil {
			return err
		}
	}

	if dryRun {
		if shScript != "" {
			fmt.Printf("%s\n\n", tui.Header("scaffold.sh"))
			fmt.Printf("%s\n\n", tui.Dim(shScript))
		}
		if ps1Script != "" {
			fmt.Printf("%s\n\n", tui.Header("scaffold.ps1"))
			fmt.Printf("%s", tui.Dim(ps1Script))
		}
		return nil
	}

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create out dir: %w", err)
	}

	if shScript != "" {
		path := filepath.Join(outDir, "scaffold.sh")
		if err := os.WriteFile(path, []byte(shScript), 0o755); err != nil {
			return fmt.Errorf("write scaffold.sh: %w", err)
		}
		fmt.Println(tui.Success("Generated: " + path))
		fmt.Println()
		fmt.Printf("%s\n", tui.Info("Run '.\\"+path+"' to execute"))
	}

	if ps1Script != "" {
		path := filepath.Join(outDir, "scaffold.ps1")
		if err := os.WriteFile(path, []byte(ps1Script), 0o644); err != nil {
			return fmt.Errorf("write scaffold.ps1: %w", err)
		}
		fmt.Println(tui.Success("Generated: " + path))
		fmt.Println()
		fmt.Printf("%s\n", tui.Info("Run '.\\"+path+"' to execute"))
	}

	return nil
}
