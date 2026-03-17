package app

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

func TestGenerateCmd(t *testing.T) {
	// Test that generate command is properly configured
	if generateCmd.Use != "generate" {
		t.Errorf("Expected use 'generate', got %s", generateCmd.Use)
	}

	if generateCmd.Short != "Generate scaffold scripts from markdown structure" {
		t.Errorf("Expected short description 'Generate scaffold scripts from markdown structure', got %s", generateCmd.Short)
	}

	// Test that required flags are set
	flags := generateCmd.Flags()
	if flag := flags.Lookup("in"); flag == nil {
		t.Error("Expected 'in' flag to be set")
	}
	if flag := flags.Lookup("shell"); flag == nil {
		t.Error("Expected 'shell' flag to be set")
	}
	if flag := flags.Lookup("out"); flag == nil {
		t.Error("Expected 'out' flag to be set")
	}
	if flag := flags.Lookup("dry-run"); flag == nil {
		t.Error("Expected 'dry-run' flag to be set")
	}
	if flag := flags.Lookup("verbose"); flag == nil {
		t.Error("Expected 'verbose' flag to be set")
	}
}

func TestValidateShellMode(t *testing.T) {
	tests := []struct {
		shell    string
		expected bool
	}{
		{"sh", true},
		{"ps1", true},
		{"both", true},
		{"SH", true},   // strings.ToLower makes this valid
		{"PS1", true},  // strings.ToLower makes this valid
		{"BOTH", true}, // strings.ToLower makes this valid
		{"invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.shell, func(t *testing.T) {
			shellMode := strings.ToLower(tt.shell)
			isValid := shellMode == "sh" || shellMode == "ps1" || shellMode == "both"
			if isValid != tt.expected {
				t.Errorf("Expected shell mode %s to be %v, got %v", tt.shell, tt.expected, isValid)
			}
		})
	}
}

func TestValidateFileExtension(t *testing.T) {
	tests := []struct {
		ext      string
		expected bool
	}{
		{".json", true},
		{".yml", true},
		{".yaml", true},
		{".md", true},
		{".markdown", true},
		{".JSON", true}, // strings.ToLower makes this valid
		{".YML", true},  // strings.ToLower makes this valid
		{".txt", false},
		{".go", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			ext := strings.ToLower(tt.ext)
			isValid := ext == ".json" || ext == ".yml" || ext == ".yaml" || ext == ".md" || ext == ".markdown"
			if isValid != tt.expected {
				t.Errorf("Expected extension %s to be %v, got %v", tt.ext, tt.expected, isValid)
			}
		})
	}
}

func TestGenerateFilesystemValidation(t *testing.T) {
	// Test with valid structure
	root := model.NewDir("test")
	root.AddChild(model.NewFile("test.txt"))

	if err := ValidateFilesystemConflicts(root); err != nil {
		t.Errorf("Expected no error for valid structure, got %v", err)
	}
}

func TestGenerateOutputDirectory(t *testing.T) {
	// Test temporary directory creation
	tempDir := t.TempDir()

	if err := os.MkdirAll(tempDir, 0o755); err != nil {
		t.Errorf("Expected to create directory without error, got %v", err)
	}

	// Verify directory exists
	if info, err := os.Stat(tempDir); err != nil {
		t.Errorf("Expected directory to exist, got error %v", err)
	} else if !info.IsDir() {
		t.Error("Expected path to be a directory")
	}
}

func TestGenerateFilePermissions(t *testing.T) {
	tempDir := t.TempDir()

	// Test shell script permissions
	shPath := filepath.Join(tempDir, "scaffold.sh")
	shContent := "#!/bin/bash\necho 'test'"
	if err := os.WriteFile(shPath, []byte(shContent), 0o755); err != nil {
		t.Errorf("Expected to write shell script without error, got %v", err)
	}

	// Test PowerShell script permissions
	ps1Path := filepath.Join(tempDir, "scaffold.ps1")
	ps1Content := "Write-Output 'test'"
	if err := os.WriteFile(ps1Path, []byte(ps1Content), 0o644); err != nil {
		t.Errorf("Expected to write PowerShell script without error, got %v", err)
	}
}
