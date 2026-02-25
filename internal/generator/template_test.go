package generator

import (
	"strings"
	"testing"
	"time"
)

func TestGetLanguageFromExtension(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{"main.go", "go"},
		{"models.py", "python"},
		{"app.tsx", "typescript"},
		{"utils.js", "javascript"},
		{"handler.rs", "rust"},
		{"Main.java", "java"},
		{"app.cpp", "cpp"},
		{"config.c", "c"},
		{"header.h", "c-header"},
		{"app.hpp", "cpp-header"},
		{"app.rb", "ruby"},
		{"index.php", "php"},
		{"main.kt", "kotlin"},
		{"app.swift", "swift"},
		{"Program.cs", "csharp"},
		{"script.sh", "shell"},
		{"deploy.ps1", "powershell"},
		{"module.lua", "lua"},
		{"unknown.xyz", "text"},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := GetLanguageFromExtension(tt.filename)
			if result != tt.expected {
				t.Errorf("GetLanguageFromExtension(%q) = %q, want %q", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestGetFileExtension(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{"main.go", ".go"},
		{"Makefile", ""},
		{"archive.tar.gz", ".gz"},
		{".hidden", ".hidden"},
		{"file_without_ext", ""},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			result := GetFileExtension(tt.filename)
			if result != tt.expected {
				t.Errorf("GetFileExtension(%q) = %q, want %q", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"user_handler", "userHandler"},
		{"api_router", "apiRouter"},
		{"db_connection", "dbConnection"},
		{"simple", "simple"},
		{"has_id", "hasId"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ToCamelCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToCamelCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"userHandler", "user_handler"},
		{"APIRouter", "a_p_i_router"},
		{"simple", "simple"},
		{"UserManager", "user_manager"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ToSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToSnakeCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestHasRenderer(t *testing.T) {
	tests := []struct {
		ext      string
		expected bool
	}{
		{".go", true},
		{".py", true},
		{".ts", true},
		{".tsx", true},
		{".xyz", false},
		{".unknown", false},
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			result := HasRenderer(tt.ext)
			if result != tt.expected {
				t.Errorf("HasRenderer(%q) = %v, want %v", tt.ext, result, tt.expected)
			}
		})
	}
}

func TestGoRenderer(t *testing.T) {
	ctx := TemplateContext{
		Package:  "main",
		Filename: "main.go",
		Language: "go",
	}

	result := goRenderer(ctx)

	if !strings.Contains(result, "package") {
		t.Errorf("goRenderer() missing 'package' keyword")
	}
	if !strings.Contains(result, "main") {
		t.Errorf("goRenderer() missing package name")
	}
}

func TestPythonRenderer(t *testing.T) {
	ctx := TemplateContext{
		Package:  "models",
		Filename: "models.py",
		Language: "python",
	}

	result := pythonRenderer(ctx)

	if !strings.Contains(result, "\"\"\"") {
		t.Errorf("pythonRenderer() missing docstring quotes")
	}
	if !strings.Contains(result, "models") {
		t.Errorf("pythonRenderer() missing module name")
	}
}

func TestTypescriptRenderer(t *testing.T) {
	ctx := TemplateContext{
		Package:  "components",
		Filename: "App.tsx",
		Language: "typescript",
	}

	result := typescriptRenderer(ctx)

	if !strings.Contains(result, "/**") || !strings.Contains(result, "*/") {
		t.Errorf("typescriptRenderer() missing JSDoc comments")
	}
	if !strings.Contains(result, "components") {
		t.Errorf("typescriptRenderer() missing module name")
	}
}

func TestApplyVariableSubstitution(t *testing.T) {
	ctx := TemplateContext{
		Package:  "mymodule",
		Filename: "handler.go",
		Language: "go",
		Path:     "cmd/api/handler.go",
		Date:     time.Date(2025, 6, 15, 10, 30, 0, 0, time.UTC),
		Author:   "testauthor",
		Metadata: map[string]interface{}{},
	}

	tests := []struct {
		input    string
		expected string
	}{
		{"{{package}}", "mymodule"},
		{"{{filename}}", "handler.go"},
		{"{{language}}", "go"},
		{"{{path}}", "cmd/api/handler.go"},
		{"{{author}}", "testauthor"},
		{"{{year}}", "2025"},
		{"Module: {{package}}", "Module: mymodule"},
		{"// {{filename}} - {{date}}", "// handler.go - 2025-06-15"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ApplyVariableSubstitution(tt.input, ctx)
			if result != tt.expected {
				t.Errorf("ApplyVariableSubstitution(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestValidateGoTemplate(t *testing.T) {
	validGo := "package main\n\nfunc main() {}"
	if err := validateGoTemplate(validGo); err != nil {
		t.Errorf("validateGoTemplate() failed for valid code: %v", err)
	}

	invalidGo := "func main() {}"
	if err := validateGoTemplate(invalidGo); err == nil {
		t.Errorf("validateGoTemplate() should reject code without package")
	}
}

func TestValidatePythonTemplate(t *testing.T) {
	validPy := `"""Module docstring."""`
	if err := validatePythonTemplate(validPy); err != nil {
		t.Errorf("validatePythonTemplate() failed: %v", err)
	}

	invalidPy := `"""Unbalanced docstring."`
	if err := validatePythonTemplate(invalidPy); err == nil {
		t.Errorf("validatePythonTemplate() should reject unbalanced quotes")
	}
}

func TestRenderTemplateWithContext(t *testing.T) {
	ctx := TemplateContext{
		Package:  "handlers",
		Filename: "router.go",
		Language: "go",
	}

	result := RenderTemplateWithContext(".go", ctx)

	if result == "" {
		t.Error("RenderTemplateWithContext() returned empty string")
	}
	if !strings.Contains(result, "package") {
		t.Error("RenderTemplateWithContext() missing package declaration")
	}
}

func TestRenderTemplate(t *testing.T) {
	result := RenderTemplate(".py", "models.py", "models", "python", "app/models.py")

	if result == "" {
		t.Error("RenderTemplate() returned empty string")
	}
	if !strings.Contains(result, "models") {
		t.Error("RenderTemplate() missing module reference")
	}
}

func TestNewTemplateContext(t *testing.T) {
	ctx := NewTemplateContext("handler.go", "api", "go", "cmd/api/handler.go")

	if ctx.Filename != "handler.go" {
		t.Errorf("NewTemplateContext() filename = %q, want %q", ctx.Filename, "handler.go")
	}
	if ctx.Package != "api" {
		t.Errorf("NewTemplateContext() package = %q, want %q", ctx.Package, "api")
	}
	if ctx.Language != "go" {
		t.Errorf("NewTemplateContext() language = %q, want %q", ctx.Language, "go")
	}
	if ctx.Path != "cmd/api/handler.go" {
		t.Errorf("NewTemplateContext() path = %q, want %q", ctx.Path, "cmd/api/handler.go")
	}
	if ctx.Metadata == nil {
		t.Error("NewTemplateContext() metadata is nil")
	}
}
