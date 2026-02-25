package generator

import (
	"testing"
)

func TestFormatFilename(t *testing.T) {
	tests := []struct {
		name     string
		language string
		expected string
	}{
		{"user_handler", "java", "UserHandler"},
		{"api_server", "csharp", "ApiServer"},
		{"utils_helper", "go", "utils_helper"},
		{"database_pool", "python", "database_pool"},
	}

	for _, tt := range tests {
		result := FormatFilename(tt.name, tt.language)
		if result != tt.expected {
			t.Errorf("FormatFilename(%s, %s) = %s, want %s", tt.name, tt.language, result, tt.expected)
		}
	}
}

func TestNewTemplateContextMetadata(t *testing.T) {
	ctx := NewTemplateContext("main.go", "main", "go", "cmd/main.go")

	if ctx.Filename != "main.go" {
		t.Errorf("Expected filename 'main.go', got %s", ctx.Filename)
	}
	if ctx.Package != "main" {
		t.Errorf("Expected package 'main', got %s", ctx.Package)
	}
	if ctx.Language != "go" {
		t.Errorf("Expected language 'go', got %s", ctx.Language)
	}
	if ctx.Path != "cmd/main.go" {
		t.Errorf("Expected path 'cmd/main.go', got %s", ctx.Path)
	}
	if ctx.Metadata == nil {
		t.Error("Expected metadata map to be initialized")
	}
}

func TestParseNameFormatting(t *testing.T) {
	tests := []struct {
		name     string
		language string
	}{
		{"user_handler.go", "go"},
		{"UserHandler.java", "java"},
		{"api_client.ts", "typescript"},
	}

	for _, tt := range tests {
		base, className, funcName, varName := ParseNameFormatting(tt.name, tt.language)
		if len(base) == 0 {
			t.Errorf("ParseNameFormatting(%s, %s) should return non-empty base", tt.name, tt.language)
		}
		if len(className) == 0 {
			t.Errorf("ParseNameFormatting(%s, %s) should return non-empty className", tt.name, tt.language)
		}
		if len(funcName) == 0 {
			t.Errorf("ParseNameFormatting(%s, %s) should return non-empty funcName", tt.name, tt.language)
		}
		if len(varName) == 0 {
			t.Errorf("ParseNameFormatting(%s, %s) should return non-empty varName", tt.name, tt.language)
		}
	}
}

func TestParseNameFormattingJava(t *testing.T) {
	base, className, _, _ := ParseNameFormatting("user_handler.java", "java")
	expectedClass := "UserHandler"
	if className != expectedClass {
		t.Errorf("Java class name should be '%s', got '%s'", expectedClass, className)
	}
	if len(base) == 0 {
		t.Error("Base name should not be empty")
	}
}
