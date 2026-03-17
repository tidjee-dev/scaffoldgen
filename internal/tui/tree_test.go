package tui

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

func TestRenderTreeNil(t *testing.T) {
	// Test with nil node
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	RenderTree(nil, "", true, "")

	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)

	// Should not print anything for nil node
	if buf.Len() > 0 {
		t.Errorf("Expected no output for nil node, got %q", buf.String())
	}
}

func TestRenderTreeSingleFile(t *testing.T) {
	// Test with single file node
	root := model.NewFile("test.txt")

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	RenderTree(root, "", true, "")

	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)

	output := buf.String()
	if !contains(output, "test.txt") {
		t.Errorf("Expected output to contain 'test.txt', got %q", output)
	}
}

func TestRenderTreeSingleDirectory(t *testing.T) {
	// Test with single directory node
	root := model.NewDir("testdir")

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	RenderTree(root, "", true, "")

	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)

	output := buf.String()
	if !contains(output, "testdir") {
		t.Errorf("Expected output to contain 'testdir', got %q", output)
	}
}

func TestTreeLine(t *testing.T) {
	tests := []struct {
		prefix   string
		last     bool
		label    string
		expected string
	}{
		{"", true, "root", "root"},
		{"", false, "root", "root"},
		{"├── ", true, "file", "├── └── file"},
		{"├── ", false, "file", "├── ├── file"},
		{"    ", true, "child", "    └── child"},
		{"    ", false, "child", "    ├── child"},
	}

	for _, tt := range tests {
		t.Run(tt.prefix+tt.label, func(t *testing.T) {
			result := treeLine(tt.prefix, tt.last, tt.label)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestRenderNodeLabel(t *testing.T) {
	// Create a temporary file for testing
	tempDir := t.TempDir()
	existingFile := filepath.Join(tempDir, "existing.txt")
	if err := os.WriteFile(existingFile, []byte("test"), 0o644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name     string
		fullPath string
		ignore   bool
		expectIn string
	}{
		{"existing.txt", existingFile, false, "existing.txt"},            // Should be wrapped in Exists()
		{"new.txt", filepath.Join(tempDir, "new.txt"), false, "new.txt"}, // Should be wrapped in Added()
		{"ignored.txt", "ignored.txt", true, "ignored.txt"},              // Should be wrapped in Ignored()
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := model.NewFile(tt.name)
			node.Ignore = tt.ignore

			result := renderNodeLabel(node, tt.fullPath)

			if !contains(result, tt.expectIn) {
				t.Errorf("Expected result to contain %s, got %s", tt.expectIn, result)
			}
		})
	}
}

func TestRenderTreeWithChildren(t *testing.T) {
	// Test with directory containing children
	root := model.NewDir("project")
	root.AddChild(model.NewDir("src"))
	root.AddChild(model.NewFile("README.md"))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	RenderTree(root, "", false, "")

	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)

	output := buf.String()

	// Check that tree structure is rendered
	expectedElements := []string{
		"project",
		"├──",
		"src",
		"└──",
		"README.md",
	}

	for _, element := range expectedElements {
		if !contains(output, element) {
			t.Errorf("Expected output to contain %q, got %q", element, output)
		}
	}
}

func TestRenderTreeIgnoredChildren(t *testing.T) {
	// Test that ignored children are not rendered
	root := model.NewDir("project")
	root.AddChild(model.NewDir("src"))

	ignoredFile := model.NewFile("ignored.txt")
	ignoredFile.Ignore = true
	root.AddChild(ignoredFile)

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	RenderTree(root, "", false, "")

	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)

	output := buf.String()

	// Should contain src but not ignored.txt
	if !contains(output, "src") {
		t.Errorf("Expected output to contain 'src', got %q", output)
	}

	// The ignored file should be wrapped in Ignored() styling
	if !contains(output, "ignored.txt") {
		t.Errorf("Expected ignored.txt to be rendered with Ignored() styling, got %q", output)
	}

	if !contains(output, "ignored") {
		t.Errorf("Expected ignored.txt to be wrapped in Ignored() styling, got %q", output)
	}
}

func TestRenderTreePathConstruction(t *testing.T) {
	// Test that full paths are constructed correctly
	root := model.NewDir("root")
	child := model.NewDir("child")
	grandchild := model.NewFile("file.txt")

	child.AddChild(grandchild)
	root.AddChild(child)

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	RenderTree(root, "", false, "")

	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)

	output := buf.String()

	// Check that tree structure is rendered correctly
	if !contains(output, "root") {
		t.Errorf("Expected output to contain 'root', got %q", output)
	}

	if !contains(output, "child") {
		t.Errorf("Expected output to contain 'child', got %q", output)
	}

	if !contains(output, "file.txt") {
		t.Errorf("Expected output to contain 'file.txt', got %q", output)
	}
}

func TestRenderTreePrefixHandling(t *testing.T) {
	// Test that prefixes are handled correctly for nested structures
	root := model.NewDir("root")
	child1 := model.NewDir("child1")
	child2 := model.NewDir("child2")

	root.AddChild(child1)
	root.AddChild(child2)

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	RenderTree(root, "", false, "")

	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)

	output := buf.String()

	// Check that proper tree characters are used
	if !contains(output, "├──") {
		t.Errorf("Expected output to contain '├──' for non-last child, got %q", output)
	}

	if !contains(output, "└──") {
		t.Errorf("Expected output to contain '└──' for last child, got %q", output)
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && indexOf(s, substr) >= 0))
}

// Simple indexOf implementation for substring search
func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
