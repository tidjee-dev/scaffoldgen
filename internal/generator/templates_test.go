package generator

import (
	"strings"
	"testing"
)

func TestGoTemplate(t *testing.T) {
	result := goTemplate("main")
	if !strings.Contains(result, "package main") {
		t.Errorf("goTemplate should contain 'package main', got: %s", result)
	}
}

func TestPythonTemplate(t *testing.T) {
	result := pythonTemplate("utils")
	if !strings.Contains(result, "utils") {
		t.Errorf("pythonTemplate should contain 'utils', got: %s", result)
	}
	if !strings.Contains(result, "module") {
		t.Errorf("pythonTemplate should contain 'module', got: %s", result)
	}
}

func TestTypescriptTemplate(t *testing.T) {
	result := typescriptTemplate("client")
	if !strings.Contains(result, "client") {
		t.Errorf("typescriptTemplate should contain 'client', got: %s", result)
	}
	if !strings.Contains(result, "//") {
		t.Errorf("typescriptTemplate should contain comment, got: %s", result)
	}
}

func TestJavaTemplate(t *testing.T) {
	result := javaTemplate("handler")
	if !strings.Contains(result, "public class") {
		t.Errorf("javaTemplate should contain 'public class', got: %s", result)
	}
}

func TestRustTemplate(t *testing.T) {
	result := rustTemplate("app")
	if !strings.Contains(result, "//") {
		t.Errorf("rustTemplate should contain comment, got: %s", result)
	}
	if !strings.Contains(result, "app") {
		t.Errorf("rustTemplate should contain 'app', got: %s", result)
	}
}

func TestCppTemplate(t *testing.T) {
	result := cppTemplate("core")
	if !strings.Contains(result, "include") {
		t.Errorf("cppTemplate should contain 'include', got: %s", result)
	}
}

func TestCTemplate(t *testing.T) {
	result := cTemplate("lib")
	if !strings.Contains(result, "include") {
		t.Errorf("cTemplate should contain 'include', got: %s", result)
	}
}

func TestHeaderTemplate(t *testing.T) {
	result := cHeaderTemplate("MyLib")
	if !strings.Contains(result, "#ifndef") {
		t.Errorf("cHeaderTemplate should contain '#ifndef', got: %s", result)
	}
	if !strings.Contains(result, "#endif") {
		t.Errorf("cHeaderTemplate should contain '#endif', got: %s", result)
	}
}

func TestGetTemplate(t *testing.T) {
	tests := []struct {
		ext      string
		expected bool
	}{
		{".go", true},
		{".py", true},
		{".ts", true},
		{".java", true},
		{".unknown", false},
	}

	for _, tt := range tests {
		result := GetTemplate(tt.ext)
		if tt.expected && result == nil {
			t.Errorf("GetTemplate(%s) should return a provider", tt.ext)
		}
		if !tt.expected && result != nil {
			t.Errorf("GetTemplate(%s) should return nil", tt.ext)
		}
	}
}

func TestJavaRenderer(t *testing.T) {
	ctx := TemplateContext{
		Package:  "handler",
		Filename: "Handler.java",
		Language: "java",
	}
	result := javaRenderer(ctx)
	if !strings.Contains(result, "public class") {
		t.Errorf("javaRenderer should contain 'public class', got: %s", result)
	}
}

func TestRustRenderer(t *testing.T) {
	ctx := TemplateContext{
		Package:  "app",
		Filename: "lib.rs",
		Language: "rust",
	}
	result := rustRenderer(ctx)
	if !strings.Contains(result, "//!") {
		t.Errorf("rustRenderer should contain '//!', got: %s", result)
	}
}

func TestShellRenderer(t *testing.T) {
	ctx := TemplateContext{
		Package:  "deploy",
		Filename: "deploy.sh",
		Language: "shell",
	}
	result := shellRenderer(ctx)
	if !strings.Contains(result, "#!/usr/bin/env bash") {
		t.Errorf("shellRenderer should contain shebang, got: %s", result)
	}
}

func TestLuaRenderer(t *testing.T) {
	ctx := TemplateContext{
		Package:  "app",
		Filename: "app.lua",
		Language: "lua",
	}
	result := luaRenderer(ctx)
	if !strings.Contains(result, "--") {
		t.Errorf("luaRenderer should contain comment, got: %s", result)
	}
	if !strings.Contains(result, "return") {
		t.Errorf("luaRenderer should contain return, got: %s", result)
	}
}

func TestCsharpRenderer(t *testing.T) {
	ctx := TemplateContext{
		Package:  "api_handler",
		Filename: "Handler.cs",
		Language: "csharp",
	}
	result := csharpRenderer(ctx)
	if !strings.Contains(result, "namespace") {
		t.Errorf("csharpRenderer should contain namespace, got: %s", result)
	}
	if !strings.Contains(result, "public class") {
		t.Errorf("csharpRenderer should contain public class, got: %s", result)
	}
}

func TestPHPRenderer(t *testing.T) {
	ctx := TemplateContext{
		Package:  "handlers",
		Filename: "Handler.php",
		Language: "php",
	}
	result := phpRenderer(ctx)
	if !strings.Contains(result, "<?php") {
		t.Errorf("phpRenderer should contain PHP tag, got: %s", result)
	}
	if !strings.Contains(result, "namespace") {
		t.Errorf("phpRenderer should contain namespace, got: %s", result)
	}
}

func TestTemplateRendererMap(t *testing.T) {
	expectedRenderers := []string{".go", ".py", ".ts", ".js", ".java"}

	for _, ext := range expectedRenderers {
		if _, exists := renderers[ext]; !exists {
			t.Errorf("Renderer map should contain %s", ext)
		}
	}
}
