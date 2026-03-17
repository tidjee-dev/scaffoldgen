package app

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var (
	verboseVersion bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long: `Display version information for scaffoldgen.
Use --verbose to see detailed build information including commit and build date.`,
	RunE: runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&verboseVersion, "verbose", "V", false, "Show detailed version information")
}

func runVersion(cmd *cobra.Command, args []string) error {
	if verboseVersion {
		return showDetailedVersion()
	}
	PrintVersion()
	return nil
}

func showDetailedVersion() error {
	fmt.Printf("scaffoldgen %s\n", Version)

	// Try to read version.json for additional info
	versionFile := filepath.Join(".", "version.json")
	if data, err := os.ReadFile(versionFile); err == nil {
		var config struct {
			Build struct {
				Date   string `json:"date"`
				Commit string `json:"commit"`
				Dirty  bool   `json:"dirty"`
			} `json:"build"`
			Metadata struct {
				GoVersion string `json:"go_version"`
				Module    string `json:"module"`
			} `json:"metadata"`
		}

		if err := json.Unmarshal(data, &config); err == nil {
			fmt.Println("\nBuild Information:")
			if config.Build.Commit != "" {
				shortCommit := config.Build.Commit
				if len(shortCommit) > 8 {
					shortCommit = shortCommit[:8]
				}
				fmt.Printf("  Commit: %s", shortCommit)
				if config.Build.Dirty {
					fmt.Printf(" (dirty)")
				}
				fmt.Println()
			}

			if config.Build.Date != "" {
				if parsed, err := time.Parse(time.RFC3339, config.Build.Date); err == nil {
					fmt.Printf("  Built:  %s\n", parsed.Format("2006-01-02 15:04:05"))
				} else {
					fmt.Printf("  Built:  %s\n", config.Build.Date)
				}
			}

			fmt.Println("\nRuntime:")
			fmt.Printf("  Go:     %s\n", config.Metadata.GoVersion)
			fmt.Printf("  Module: %s\n", config.Metadata.Module)
		}
	}

	return nil
}
