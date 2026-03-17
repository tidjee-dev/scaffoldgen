package app

import (
	"strings"
	"testing"
)

func TestIsNewerVersion(t *testing.T) {
	tests := []struct {
		name     string
		current  string
		latest   string
		expected bool
	}{
		{
			name:     "same version",
			current:  "1.0.0",
			latest:   "1.0.0",
			expected: false,
		},
		{
			name:     "newer patch",
			current:  "1.0.0",
			latest:   "1.0.1",
			expected: true,
		},
		{
			name:     "newer minor",
			current:  "1.0.0",
			latest:   "1.1.0",
			expected: true,
		},
		{
			name:     "newer major",
			current:  "1.0.0",
			latest:   "2.0.0",
			expected: true,
		},
		{
			name:     "older version",
			current:  "1.1.0",
			latest:   "1.0.0",
			expected: false,
		},
		{
			name:     "with v prefix",
			current:  "v1.0.0",
			latest:   "v1.0.1",
			expected: true,
		},
		{
			name:     "mixed prefix",
			current:  "1.0.0",
			latest:   "v1.0.1",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isNewerVersion(tt.current, tt.latest)
			if result != tt.expected {
				t.Errorf("isNewerVersion(%s, %s) = %v; want %v", tt.current, tt.latest, result, tt.expected)
			}
		})
	}
}

func TestUpdateCmd(t *testing.T) {
	// Test that update command is properly configured
	if updateCmd.Use != "update" {
		t.Errorf("Expected use 'update', got %s", updateCmd.Use)
	}

	if updateCmd.Short != "Update scaffoldgen to the latest version" {
		t.Errorf("Expected short description 'Update scaffoldgen to the latest version', got %s", updateCmd.Short)
	}
}

func TestGetBinaryInfo(t *testing.T) {
	// Test getBinaryInfo function
	binaryURL, binaryPath, err := getBinaryInfo()
	if err != nil {
		t.Errorf("getBinaryInfo() failed: %v", err)
	}

	if binaryURL == "" {
		t.Error("Expected binary URL to be non-empty")
	}

	if binaryPath == "" {
		t.Error("Expected binary path to be non-empty")
	}

	// Check that URL contains expected patterns
	if !strings.Contains(binaryURL, "github.com/tidjee-dev/scaffoldgen/releases") {
		t.Errorf("Expected URL to contain GitHub releases path, got %s", binaryURL)
	}
}
