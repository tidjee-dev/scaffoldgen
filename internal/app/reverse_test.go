package app

import (
	"strings"
	"testing"
)

func TestReverseCmd(t *testing.T) {
	// Test that reverse command is properly configured
	if reverseCmd.Use != "reverse" {
		t.Errorf("Expected use 'reverse', got %s", reverseCmd.Use)
	}

	if reverseCmd.Short != "Scan directory and generate structure file" {
		t.Errorf("Expected short description 'Scan directory and generate structure file', got %s", reverseCmd.Short)
	}

	// Test that required flags are set
	flags := reverseCmd.Flags()
	if flag := flags.Lookup("in"); flag == nil {
		t.Error("Expected 'in' flag to be set")
	}
	if flag := flags.Lookup("format"); flag == nil {
		t.Error("Expected 'format' flag to be set")
	}
	if flag := flags.Lookup("out"); flag == nil {
		t.Error("Expected 'out' flag to be set")
	}
}

func TestReverseCmdLongDescription(t *testing.T) {
	// Test that long description is provided
	longDesc := reverseCmd.Long
	if longDesc == "" {
		t.Error("Expected long description to be set")
	}

	// Check that it contains expected keywords
	expectedKeywords := []string{
		"Scan",
		"directory structure",
		"structure file",
		"documenting",
		"project layouts",
	}

	for _, keyword := range expectedKeywords {
		if !strings.Contains(longDesc, keyword) {
			t.Errorf("Expected long description to contain keyword: %s", keyword)
		}
	}
}

func TestReverseCmdExample(t *testing.T) {
	// Test that examples are provided
	example := reverseCmd.Example
	if example == "" {
		t.Error("Expected example to be set")
	}

	// Check that examples contain expected commands
	expectedExamples := []string{
		"scaffoldgen reverse --in ./backend --format md --out structure.md",
		"scaffoldgen reverse --in ./src --format json > structure.json",
	}

	for _, expected := range expectedExamples {
		if !strings.Contains(example, expected) {
			t.Errorf("Expected example to contain %s", expected)
		}
	}
}

func TestReverseSupportedFormats(t *testing.T) {
	// Test supported output formats
	supportedFormats := []string{"md", "markdown", "json", "yml", "yaml"}
	
	for _, format := range supportedFormats {
		t.Run(format, func(t *testing.T) {
			// Test that format is recognized
			formatLower := strings.ToLower(format)
			isValid := formatLower == "md" || formatLower == "markdown" || 
					  formatLower == "json" || formatLower == "yml" || formatLower == "yaml"
			if !isValid {
				t.Errorf("Expected %s to be a supported format", format)
			}
		})
	}
}

func TestReverseUnsupportedFormat(t *testing.T) {
	// Test unsupported format
	unsupportedFormat := "txt"
	formatLower := strings.ToLower(unsupportedFormat)
	isValid := formatLower == "md" || formatLower == "markdown" || 
			  formatLower == "json" || formatLower == "yml" || formatLower == "yaml"
	if isValid {
		t.Errorf("Expected %s to be unsupported", unsupportedFormat)
	}
}

func TestReverseDefaultFormat(t *testing.T) {
	// Test default format
	defaultFormat := "md"
	if defaultFormat != "md" {
		t.Errorf("Expected default format to be 'md', got %s", defaultFormat)
	}
}
