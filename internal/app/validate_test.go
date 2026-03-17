package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

func TestValidateFilesystemConflicts(t *testing.T) {
	// Test with valid structure (no conflicts)
	root := model.NewDir("test")
	root.AddChild(model.NewFile("test.txt"))

	if err := ValidateFilesystemConflicts(root); err != nil {
		t.Errorf("Expected no error for valid structure, got %v", err)
	}
}

func TestValidateFilesystemConflictsEmpty(t *testing.T) {
	// Test with empty structure
	root := model.NewDir("empty")

	if err := ValidateFilesystemConflicts(root); err != nil {
		t.Errorf("Expected no error for empty structure, got %v", err)
	}
}

func TestValidateFilesystemConflictsNested(t *testing.T) {
	// Test with nested structure
	root := model.NewDir("test")
	subdir := model.NewDir("subdir")
	subdir.AddChild(model.NewFile("nested.txt"))
	root.AddChild(subdir)

	if err := ValidateFilesystemConflicts(root); err != nil {
		t.Errorf("Expected no error for nested structure, got %v", err)
	}
}

func TestValidateFilesystemConflictsWithExistingFiles(t *testing.T) {
	// Create temporary directory for testing
	tempDir := t.TempDir()

	// Create an existing file
	existingFile := filepath.Join(tempDir, "existing.txt")
	if err := os.WriteFile(existingFile, []byte("test content"), 0o644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test structure that expects a directory where file exists
	root := model.NewDir(tempDir)
	root.AddChild(model.NewDir("existing.txt")) // Expecting directory but file exists

	err := ValidateFilesystemConflicts(root)
	if err == nil {
		t.Error("Expected error for directory/file conflict")
	}

	// Check that error message contains expected information
	errStr := err.Error()
	if !contains(errStr, "existing.txt") {
		t.Errorf("Expected error message to mention existing.txt, got: %s", errStr)
	}
	if !contains(errStr, "FILE but expected DIRECTORY") {
		t.Errorf("Expected error message to mention file/directory conflict, got: %s", errStr)
	}
}

func TestValidateFilesystemConflictsWithExistingDirectory(t *testing.T) {
	// Create temporary directory for testing
	tempDir := t.TempDir()

	// Create an existing directory
	existingDir := filepath.Join(tempDir, "existingdir")
	if err := os.MkdirAll(existingDir, 0o755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Test structure that expects a file where directory exists
	root := model.NewDir(tempDir)
	root.AddChild(model.NewFile("existingdir")) // Expecting file but directory exists

	err := ValidateFilesystemConflicts(root)
	if err == nil {
		t.Error("Expected error for file/directory conflict")
	}

	// Check that error message contains expected information
	errStr := err.Error()
	if !contains(errStr, "existingdir") {
		t.Errorf("Expected error message to mention existingdir, got: %s", errStr)
	}
	if !contains(errStr, "DIRECTORY but expected FILE") {
		t.Errorf("Expected error message to mention directory/file conflict, got: %s", errStr)
	}
}

func TestValidateFilesystemConflictsMultipleConflicts(t *testing.T) {
	// Create temporary directory for testing
	tempDir := t.TempDir()

	// Create existing file and directory
	existingFile := filepath.Join(tempDir, "conflict1.txt")
	if err := os.WriteFile(existingFile, []byte("test"), 0o644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	existingDir := filepath.Join(tempDir, "conflict2")
	if err := os.MkdirAll(existingDir, 0o755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Test structure with multiple conflicts
	root := model.NewDir(tempDir)
	root.AddChild(model.NewDir("conflict1.txt")) // Dir expected, file exists
	root.AddChild(model.NewFile("conflict2"))    // File expected, dir exists

	err := ValidateFilesystemConflicts(root)
	if err == nil {
		t.Error("Expected error for multiple conflicts")
	}

	// Check that error message contains both conflicts
	errStr := err.Error()
	if !contains(errStr, "conflict1.txt") || !contains(errStr, "conflict2") {
		t.Errorf("Expected error message to mention both conflicts, got: %s", errStr)
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && indexOf(s, substr) >= 0))
}

// Simple indexOf implementation for substring search
func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
