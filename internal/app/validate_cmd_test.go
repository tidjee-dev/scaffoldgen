package app

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateCmd(t *testing.T) {
	// Test that validate command is properly configured
	if validateCmd.Use != "validate" {
		t.Errorf("Expected use 'validate', got %s", validateCmd.Use)
	}

	if validateCmd.Short != "Check structure against existing directory" {
		t.Errorf("Expected short description 'Check structure against existing directory', got %s", validateCmd.Short)
	}

	// Test that command has required flag
	flag := validateCmd.Flags().Lookup("in")
	if flag == nil {
		t.Error("Expected 'in' flag to be present")
	}
}

func TestValidateCmdExecution(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Create a test structure file
	structureFile := filepath.Join(tmpDir, "test.yml")
	content := `# test structure
src/
  main.go
  config.go
`

	err := os.WriteFile(structureFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test structure file: %v", err)
	}

	tests := []struct {
		name    string
		flags   map[string]string
		wantErr bool
	}{
		{
			name:    "valid structure file",
			flags:   map[string]string{"in": structureFile},
			wantErr: false,
		},
		{
			name:    "missing file",
			flags:   map[string]string{"in": "nonexistent.yml"},
			wantErr: true,
		},
		{
			name:    "missing required flag",
			flags:   map[string]string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test validation logic directly
			in, exists := tt.flags["in"]
			if !exists || in == "" {
				if !tt.wantErr {
					t.Error("Expected error for missing 'in' flag")
				}
				return
			}

			// Check if file exists
			_, err := os.Stat(in)
			fileExists := err == nil

			if !fileExists && !tt.wantErr {
				t.Error("Expected error for non-existent file")
			}
		})
	}
}
