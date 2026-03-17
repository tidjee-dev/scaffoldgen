package model

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrintTree(t *testing.T) {
	// Test with nil node
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	PrintTree(nil, 0)
	
	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)
	
	// Should not print anything for nil node
	if buf.Len() > 0 {
		t.Errorf("Expected no output for nil node, got %q", buf.String())
	}
}

func TestPrintTreeSingleFile(t *testing.T) {
	// Test with single file node
	root := NewFile("test.txt")
	
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	PrintTree(root, 0)
	
	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)
	
	output := buf.String()
	expected := "- test.txt (FILE)\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPrintTreeSingleDirectory(t *testing.T) {
	// Test with single directory node
	root := NewDir("testdir")
	
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	PrintTree(root, 0)
	
	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)
	
	output := buf.String()
	expected := "- testdir (DIR)\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestPrintTreeWithChildren(t *testing.T) {
	// Test with directory containing children
	root := NewDir("project")
	root.AddChild(NewDir("src"))
	root.AddChild(NewFile("README.md"))
	
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	PrintTree(root, 0)
	
	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)
	
	output := buf.String()
	
	// Check that all nodes are printed
	expectedLines := []string{
		"- project (DIR)",
		"  - src (DIR)",
		"  - README.md (FILE)",
	}
	
	for _, line := range expectedLines {
		if !contains(output, line) {
			t.Errorf("Expected output to contain %q, got %q", line, output)
		}
	}
}

func TestPrintTreeDepth(t *testing.T) {
	// Test that depth parameter affects indentation
	root := NewDir("root")
	child := NewDir("child")
	child.AddChild(NewFile("grandchild.txt"))
	root.AddChild(child)
	
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	PrintTree(root, 1) // Start with depth 1
	
	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)
	
	output := buf.String()
	
	// Check that root has one level of indentation
	if !contains(output, "  - root (DIR)") {
		t.Errorf("Expected root to be indented, got %q", output)
	}
	
	// Check that child has two levels of indentation
	if !contains(output, "    - child (DIR)") {
		t.Errorf("Expected child to be more indented, got %q", output)
	}
}

func TestPrintTreeNestedStructure(t *testing.T) {
	// Test with deeply nested structure
	root := NewDir("root")
	level1 := NewDir("level1")
	level2 := NewDir("level2")
	level3 := NewFile("deep.txt")
	
	level2.AddChild(level3)
	level1.AddChild(level2)
	root.AddChild(level1)
	
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	PrintTree(root, 0)
	
	w.Close()
	os.Stdout = oldStdout
	io.Copy(&buf, r)
	
	output := buf.String()
	
	// Check that all levels are present with correct indentation
	expectedLines := []string{
		"- root (DIR)",
		"  - level1 (DIR)",
		"    - level2 (DIR)",
		"      - deep.txt (FILE)",
	}
	
	for _, line := range expectedLines {
		if !contains(output, line) {
			t.Errorf("Expected output to contain %q, got %q", line, output)
		}
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
