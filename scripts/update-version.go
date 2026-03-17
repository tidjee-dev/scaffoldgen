package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	// Get commit hash
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return "", false, err
	}
	commit := strings.TrimSpace(string(output))

	// Check if working directory is dirty
	cmd = exec.Command("git", "status", "--porcelain")
	cmd.Dir = dir
	output, err = cmd.Output()
	if err != nil {
		return "", false, err
	}
	dirty := len(strings.TrimSpace(string(output))) > 0

	return commit, dirty, nil
}
