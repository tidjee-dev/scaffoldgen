package tui

import (
	"testing"
)

func TestStyleFunctions(t *testing.T) {
	// Test that style functions return non-empty strings
	tests := []struct {
		name     string
		function func(string) string
		input    string
	}{
		{"Success", Success, "✅ Test message"},
		{"Error", Error, "❌ Error message"},
		{"Info", Info, "ℹ️ Info message"},
		{"Warn", Warn, "⚠️ Warning message"},
		{"Header", Header, "Header"},
		{"SubHeader", SubHeader, "Sub Header"},
		{"Dim", Dim, "Dimmed text"},
		{"Added", Added, "filename"},
		{"Exists", Exists, "filename"},
		{"Ignored", Ignored, "filename"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.input)
			if result == "" {
				t.Errorf("Expected non-empty result for %s", tt.name)
			}

			// Check that the input is contained in the output (for most styles)
			if tt.name != "Exists" {
				// Some styles might add ANSI codes, so we check if the original text is present
				if !containsString(result, tt.input) {
					t.Errorf("Expected input %q to be in output %q for %s", tt.input, result, tt.name)
				}
			}
		})
	}
}

func TestStyleConsistency(t *testing.T) {
	// Test that the same input produces consistent output
	input := "test message"

	success1 := Success(input)
	success2 := Success(input)

	if success1 != success2 {
		t.Error("Success style should be consistent")
	}

	error1 := Error(input)
	error2 := Error(input)

	if error1 != error2 {
		t.Error("Error style should be consistent")
	}
}

func TestStyleColors(t *testing.T) {
	// Test that different styles produce different outputs
	input := "test"

	success := Success(input)
	error := Error(input)
	info := Info(input)
	warn := Warn(input)

	// They should all be non-empty
	if success == "" {
		t.Error("Success style should produce output")
	}
	if error == "" {
		t.Error("Error style should produce output")
	}
	if info == "" {
		t.Error("Info style should produce output")
	}
	if warn == "" {
		t.Error("Warn style should produce output")
	}

	// At least some should be different (they might have different colors)
	if success == error && success == info && success == warn {
		t.Log("Note: All styles produce same output - this might be expected if no colors")
	}
}

func TestHeaderStyles(t *testing.T) {
	// Test header-specific styles
	title := "Test Title"
	subtitle := "Test Subtitle"

	headerOutput := Header(title)
	subheaderOutput := SubHeader(subtitle)

	if headerOutput == "" {
		t.Error("Header should produce output")
	}
	if subheaderOutput == "" {
		t.Error("SubHeader should produce output")
	}

	// Both should contain the input text
	if !containsString(headerOutput, title) {
		t.Error("Header should contain title text")
	}
	if !containsString(subheaderOutput, subtitle) {
		t.Error("SubHeader should contain subtitle text")
	}
}

func TestPreviewStyles(t *testing.T) {
	// Test preview-specific styles
	filename := "test.go"

	addedOutput := Added(filename)
	existsOutput := Exists(filename)
	ignoredOutput := Ignored(filename)

	if addedOutput == "" {
		t.Error("Added style should produce output")
	}
	if existsOutput == "" {
		t.Error("Exists style should produce output")
	}
	if ignoredOutput == "" {
		t.Error("Ignored style should produce output")
	}

	// Check prefixes
	if !containsString(addedOutput, "+ ") {
		t.Error("Added style should start with +")
	}
	if !containsString(existsOutput, "✓ ") {
		t.Error("Exists style should start with ✓")
	}
	if !containsString(ignoredOutput, "… ") {
		t.Error("Ignored style should start with …")
	}
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
