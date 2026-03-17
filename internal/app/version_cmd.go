package app

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	verboseVersion bool
)

// getBuildInfo returns embedded build information
func getBuildInfo() *BuildInfo {
	// Try to get embedded build info first
	info := getEmbeddedBuildInfo()

	// If we're in development and version.json exists, use that
	if info.Version == "dev" {
		if data, err := os.ReadFile("version.json"); err == nil {
			var config BuildInfo
			if err := json.Unmarshal(data, &config); err == nil {
				info = &config
			}
		}
	}

	return info
}

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
	info := getBuildInfo()

	fmt.Printf("scaffoldgen %s\n", info.Version)

	fmt.Println("\nBuild Information:")
	if info.Build.Commit != "unknown" {
		shortCommit := info.Build.Commit
		if len(shortCommit) > 8 {
			shortCommit = shortCommit[:8]
		}
		fmt.Printf("  Commit: %s", shortCommit)
		if info.Build.Dirty {
			fmt.Printf(" (dirty)")
		}
		fmt.Println()
	}

	if info.Build.Date != "unknown" {
		if parsed, err := time.Parse(time.RFC3339, info.Build.Date); err == nil {
			fmt.Printf("  Built:  %s\n", parsed.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf("  Built:  %s\n", info.Build.Date)
		}
	}

	fmt.Println("\nRuntime:")
	fmt.Printf("  Go:     %s\n", info.Metadata.GoVersion)
	fmt.Printf("  Module: %s\n", info.Metadata.Module)

	return nil
}
