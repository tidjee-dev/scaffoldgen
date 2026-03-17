package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

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

	// Update Makefile VERSION
	if err := updateMakefile(newVersion); err != nil {
		fmt.Printf("Error updating Makefile: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Version updated to %s\n", newVersion)
	fmt.Printf("📝 Updated: Makefile\n")
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

func updateMakefile(version string) error {
	// Read Makefile
	data, err := os.ReadFile("Makefile")
	if err != nil {
		return fmt.Errorf("reading Makefile: %w", err)
	}

	content := string(data)

	// Update VERSION line
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "VERSION ?=") {
			lines[i] = fmt.Sprintf("VERSION ?= %s", version)
			break
		}
	}

	// Write back
	newContent := strings.Join(lines, "\n")
	return os.WriteFile("Makefile", []byte(newContent), 0644)
}
