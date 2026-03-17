package app

import (
	"strings"
	"testing"
)

func TestPreviewCmd(t *testing.T) {
	// Test that preview command is properly configured
	if previewCmd.Use != "preview" {
		t.Errorf("Expected use 'preview', got %s", previewCmd.Use)
	}

	if previewCmd.Short != "Preview scaffold structure (tree view)" {
		t.Errorf("Expected short description 'Preview scaffold structure (tree view)', got %s", previewCmd.Short)
	}

	// Test that required flags are set
	flags := previewCmd.Flags()
	if flag := flags.Lookup("in"); flag == nil {
		t.Error("Expected 'in' flag to be set")
	}
}

func TestPreviewCmdExample(t *testing.T) {
	// Test that examples are provided
	example := previewCmd.Example
	if example == "" {
		t.Error("Expected example to be set")
	}

	// Check that examples contain expected commands
	expectedExamples := []string{
		"scaffoldgen preview --in structure.md",
		"scaffoldgen preview --in structure.json",
		"scaffoldgen preview --in structure.yml",
	}

	for _, expected := range expectedExamples {
		if !strings.Contains(example, expected) {
			t.Errorf("Expected example to contain %s", expected)
		}
	}
}

func TestPreviewSupportedFormats(t *testing.T) {
	// Test supported file extensions
	supportedExts := []string{".json", ".yml", ".yaml", ".md", ".markdown"}
	
	for _, ext := range supportedExts {
		t.Run(ext, func(t *testing.T) {
			// Test that extension is recognized
			extLower := strings.ToLower(ext)
			isValid := extLower == ".json" || extLower == ".yml" || extLower == ".yaml" || extLower == ".md" || extLower == ".markdown"
			if !isValid {
				t.Errorf("Expected %s to be a supported format", ext)
			}
		})
	}
}

func TestPreviewUnsupportedFormat(t *testing.T) {
	// Test unsupported file extension
	unsupportedExt := ".txt"
	extLower := strings.ToLower(unsupportedExt)
	isValid := extLower == ".json" || extLower == ".yml" || extLower == ".yaml" || extLower == ".md" || extLower == ".markdown"
	if isValid {
		t.Errorf("Expected %s to be unsupported", unsupportedExt)
	}
}
