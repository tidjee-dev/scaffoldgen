package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Create test directory structure
	testDirs := []string{"src", "pkg", "docs"}
	for _, dir := range testDirs {
		err := os.MkdirAll(filepath.Join(tmpDir, dir), 0755)
		if err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	// Create test files
	testFiles := map[string]string{
		"src/main.go":    "package main\n\nfunc main() {}",
		"src/config.go":  "package main\n\ntype Config struct{}",
		"pkg/utils.go":   "package pkg\n\nfunc Helper() {}",
		"docs/README.md": "# Documentation",
	}

	for file, content := range testFiles {
		err := os.WriteFile(filepath.Join(tmpDir, file), []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to create file %s: %v", file, err)
		}
	}

	// Test scanning
	root, err := ScanDirectory(tmpDir)
	if err != nil {
		t.Fatalf("ScanDirectory failed: %v", err)
	}

	// Verify structure
	if root.Name != filepath.Base(tmpDir) {
		t.Errorf("Expected root name %s, got %s", filepath.Base(tmpDir), root.Name)
	}

	// Check that directories are found
	var foundSrc, foundPkg, foundDocs bool
	for _, child := range root.Children {
		if child.Name == "src" && child.IsDir() {
			foundSrc = true
		}
		if child.Name == "pkg" && child.IsDir() {
			foundPkg = true
		}
		if child.Name == "docs" && child.IsDir() {
			foundDocs = true
		}
	}

	if !foundSrc {
		t.Error("Expected to find src directory")
	}
	if !foundPkg {
		t.Error("Expected to find pkg directory")
	}
	if !foundDocs {
		t.Error("Expected to find docs directory")
	}
}

func TestScanDirectoryEmpty(t *testing.T) {
	// Create empty temporary directory
	tmpDir := t.TempDir()

	// Test scanning empty directory
	root, err := ScanDirectory(tmpDir)
	if err != nil {
		t.Fatalf("ScanDirectory failed: %v", err)
	}

	if len(root.Children) != 0 {
		t.Errorf("Expected no children in empty directory, got %d", len(root.Children))
	}
}

func TestScanDirectoryNonExistent(t *testing.T) {
	// Test scanning non-existent directory
	_, err := ScanDirectory("/nonexistent/directory")
	if err == nil {
		t.Error("Expected error for non-existent directory")
	}
}

func TestScanDirectoryWithIgnores(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Create test files including ones that should be ignored
	testFiles := map[string]string{
		"src/main.go":    "package main",
		"src/main.go~":   "backup file",
		"build/output.o": "binary file",
	}

	for file, content := range testFiles {
		dir := filepath.Dir(file)
		fullDir := filepath.Join(tmpDir, dir)
		err := os.MkdirAll(fullDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}

		err = os.WriteFile(filepath.Join(tmpDir, file), []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to create file %s: %v", file, err)
		}
	}

	// Test scanning with ignore patterns
	ignorePatterns := []string{"*.go~", "*.o"}
	root, err := ScanDirectoryWithIgnores(tmpDir, ignorePatterns)
	if err != nil {
		t.Fatalf("ScanDirectoryWithIgnores failed: %v", err)
	}

	// Just verify the structure is created correctly
	var foundMainGo bool
	for _, child := range root.Children {
		if child.Name == "src" && child.IsDir() {
			for _, grandchild := range child.Children {
				if grandchild.Name == "main.go" {
					foundMainGo = true
					break
				}
			}
		}
	}

	if !foundMainGo {
		t.Error("Expected to find main.go file")
	}
}
