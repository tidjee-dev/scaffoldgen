package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type VersionConfig struct {
	Version  string `json:"version"`
	Build    Build  `json:"build"`
	Metadata Meta   `json:"metadata"`
}

type Build struct {
	Date   string `json:"date"`
	Commit string `json:"commit"`
	Dirty  bool   `json:"dirty"`
}

type Meta struct {
	GoVersion string `json:"go_version"`
	Module    string `json:"module"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run scripts/update-version.go <version>")
		fmt.Println("Example: go run scripts/update-version.go 1.0.1")
		os.Exit(1)
	}

	newVersion := os.Args[1]

	// Validate version format (semantic versioning)
	if !isValidVersion(newVersion) {
		fmt.Printf("Error: Invalid version format '%s'. Expected semantic versioning format (e.g., 1.0.0, 2.1.3-beta)\n", newVersion)
		os.Exit(1)
	}

	rootDir := "."
	versionFile := filepath.Join(rootDir, "version.json")
	versionGoFile := filepath.Join(rootDir, "internal", "app", "version.go")

	// Read current version config
	config, err := readVersionConfig(versionFile)
	if err != nil {
		fmt.Printf("Error reading version config: %v\n", err)
		os.Exit(1)
	}

	// Update version
	config.Version = newVersion
	config.Build.Date = time.Now().Format(time.RFC3339)

	// Get git info
	if commit, dirty, err := getGitInfo(rootDir); err == nil {
		config.Build.Commit = commit
		config.Build.Dirty = dirty
	} else {
		// Log git error but don't fail - version update is more important
		fmt.Printf("Warning: Could not get git info: %v\n", err)
	}

	// Write updated version config
	if err := writeVersionConfig(versionFile, config); err != nil {
		fmt.Printf("Error writing version config: %v\n", err)
		os.Exit(1)
	}

	// Update Go version file
	if err := updateGoVersionFile(versionGoFile, newVersion); err != nil {
		fmt.Printf("Error updating Go version file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Version updated to %s\n", newVersion)
	fmt.Printf("📝 Updated: %s\n", versionFile)
	fmt.Printf("📝 Updated: %s\n", versionGoFile)
}

func isValidVersion(version string) bool {
	// Semantic versioning regex: X.Y.Z where X, Y, Z are numbers
	// Optional pre-release and build metadata supported
	// Examples: 1.0.0, 2.1.3-beta, 1.0.0-alpha.1, 1.0.0+build.1
	semverRegex := `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
	matched, err := regexp.MatchString(semverRegex, version)
	if err != nil {
		return false
	}
	return matched
}

func readVersionConfig(path string) (*VersionConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config VersionConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func writeVersionConfig(path string, config *VersionConfig) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func updateGoVersionFile(path string, version string) error {
	template := `package app

import "fmt"

var Version = "%s"

func PrintVersion() {
	fmt.Println("scaffoldgen", Version)
}
`

	content := fmt.Sprintf(template, version)
	return os.WriteFile(path, []byte(content), 0644)
}

func getGitInfo(dir string) (string, bool, error) {
	// Get commit hash and check if dirty in a single command to avoid race conditions
	// Use git status --porcelain=v1 -b to get branch info and dirty status together
	cmd := exec.Command("git", "status", "--porcelain=v1", "-b")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return "", false, err
	}

	// Parse output to check if dirty
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	dirty := false
	for _, line := range lines {
		// Skip branch header line (starts with ##)
		if strings.HasPrefix(line, "##") {
			continue
		}
		// Any other line indicates a dirty working directory
		if strings.TrimSpace(line) != "" {
			dirty = true
			break
		}
	}

	// Get commit hash separately (this is safe as commit hash doesn't change during status check)
	cmd = exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = dir
	output, err = cmd.Output()
	if err != nil {
		return "", false, err
	}
	commit := strings.TrimSpace(string(output))

	return commit, dirty, nil
}
