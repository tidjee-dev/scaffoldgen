package generator

import (
	"strings"
	"testing"
	"time"
)

func TestRenderTemplateWithSubstitution(t *testing.T) {
	// Test with unknown extension (should return empty)
	ctx := TemplateContext{
		Package:  "testpkg",
		Filename: "test.txt",
		Language: "go",
		Path:     "path/to/test.txt",
		Date:     time.Now(),
		Author:   "Test Author",
		Metadata: make(map[string]interface{}),
	}

	result := RenderTemplateWithSubstitution(".unknown", ctx)
	if result != "" {
		t.Errorf("Expected empty result for unknown extension, got %q", result)
	}
}

func TestValidateTemplate(t *testing.T) {
	tests := []struct {
		extension string
		content   string
		expectErr bool
		errMsg    string
	}{
		{".go", "package main\n\nfunc main() {}", false, ""},
		{".go", "func main() {}", true, "missing 'package' keyword"},
		{".py", "\"\"\"Docstring\"\"\"", false, ""},
		{".py", "\"\"\"Unclosed docstring", true, "unbalanced triple quotes"},
		{".java", "public class Test {}", false, ""},
		{".java", "package com.example;", false, ""},
		{".java", "invalid content", true, "missing class/interface or package declaration"},
		{".js", "function test() {}", false, ""},
		{".js", "function test() {", true, "unbalanced braces"},
		{".h", "#ifndef TEST_H\n#define TEST_H\ncontent", false, ""},
		{".h", "content without guards", true, "missing header guards"},
		{".unknown", "any content", false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.extension, func(t *testing.T) {
			err := ValidateTemplate(tt.extension, tt.content)
			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected error for %s, got nil", tt.extension)
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("Expected error to contain %q, got %q", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error for %s, got %v", tt.extension, err)
				}
			}
		})
	}
}

func TestStripComments(t *testing.T) {
	tests := []struct {
		extension string
		input     string
		expected  string
	}{
		{".go", "package main\n// comment\nfunc main() {}", "package main\n\nfunc main() {}"},
		{".go", "package main\n/* multi\nline comment */\nfunc main() {}", "package main\n/* multi\nline comment */\nfunc main() {}"},
		{".py", "import os\n# comment\nprint('hello')", "import os\n\nprint('hello')"},
		{".java", "public class Test {\n// comment\n}", "public class Test {\n\n}"},
		{".unknown", "content with // comment", "content with // comment"},
	}

	for _, tt := range tests {
		t.Run(tt.extension, func(t *testing.T) {
			result := stripComments(tt.input, tt.extension)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestTemplateSyntaxError(t *testing.T) {
	err := NewTemplateSyntaxError("go", "test error")

	// Test error type
	templateErr, ok := err.(*TemplateSyntaxError)
	if !ok {
		t.Error("Expected error to be of type *TemplateSyntaxError")
	}

	// Test error fields
	if templateErr.Language != "go" {
		t.Errorf("Expected language 'go', got %s", templateErr.Language)
	}

	if templateErr.Message != "test error" {
		t.Errorf("Expected message 'test error', got %s", templateErr.Message)
	}

	// Test error string
	expected := "template syntax error (go): test error"
	if err.Error() != expected {
		t.Errorf("Expected error string %q, got %q", expected, err.Error())
	}
}

func TestValidateJavaScriptTemplate(t *testing.T) {
	tests := []struct {
		content   string
		expectErr bool
	}{
		{"function test() {}", false},
		{"function test() {", true},
		{"{ balanced { braces } }", false},
		{"{ unbalanced { braces", true},
		{"no braces", false},
	}

	for _, tt := range tests {
		t.Run(tt.content, func(t *testing.T) {
			err := validateJavaScriptTemplate(tt.content)
			if tt.expectErr && err == nil {
				t.Error("Expected error, got nil")
			} else if !tt.expectErr && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func TestValidateJavaTemplate(t *testing.T) {
	tests := []struct {
		content   string
		expectErr bool
	}{
		{"public class Test {}", false},
		{"package com.example;", false},
		{"invalid content", true},
		{"interface Test {}", false},
	}

	for _, tt := range tests {
		t.Run(tt.content, func(t *testing.T) {
			err := validateJavaTemplate(tt.content)
			if tt.expectErr && err == nil {
				t.Error("Expected error, got nil")
			} else if !tt.expectErr && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func TestValidateHeaderTemplate(t *testing.T) {
	tests := []struct {
		content   string
		expectErr bool
	}{
		{"#ifndef TEST_H\n#define TEST_H\ncontent", false},
		{"content without guards", true},
		{"#define TEST_H\ncontent", true}, // missing #ifndef
		{"#ifndef TEST_H\ncontent", true}, // missing #define
	}

	for _, tt := range tests {
		t.Run(tt.content, func(t *testing.T) {
			err := validateHeaderTemplate(tt.content)
			if tt.expectErr && err == nil {
				t.Error("Expected error, got nil")
			} else if !tt.expectErr && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}
