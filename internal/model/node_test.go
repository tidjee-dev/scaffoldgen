package model

import (
	"testing"
)

func TestNewDir(t *testing.T) {
	node := NewDir("testdir")
	if node.Name != "testdir" {
		t.Errorf("Expected name 'testdir', got %s", node.Name)
	}
	if node.Type != TypeDir {
		t.Errorf("Expected type TypeDir, got %v", node.Type)
	}
	if node.IsDir() == false {
		t.Error("Expected IsDir() to return true")
	}
	if node.IsFile() == true {
		t.Error("Expected IsFile() to return false")
	}
}

func TestNewFile(t *testing.T) {
	node := NewFile("test.go")
	if node.Name != "test.go" {
		t.Errorf("Expected name 'test.go', got %s", node.Name)
	}
	if node.Type != TypeFile {
		t.Errorf("Expected type TypeFile, got %v", node.Type)
	}
	if node.IsFile() == false {
		t.Error("Expected IsFile() to return true")
	}
	if node.IsDir() == true {
		t.Error("Expected IsDir() to return false")
	}
}

func TestAddChild(t *testing.T) {
	parent := NewDir("parent")
	child := NewFile("child.go")

	parent.AddChild(child)

	if len(parent.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(parent.Children))
	}
	if parent.Children[0].Name != "child.go" {
		t.Errorf("Expected child name 'child.go', got %s", parent.Children[0].Name)
	}
}

func TestAddChildDuplicate(t *testing.T) {
	parent := NewDir("parent")
	child1 := NewFile("child.go")
	child2 := NewFile("child.go")

	parent.AddChild(child1)
	parent.AddChild(child2)

	if len(parent.Children) != 1 {
		t.Errorf("Expected 1 child (duplicates not added), got %d", len(parent.Children))
	}
}

func TestAddChildNil(t *testing.T) {
	parent := NewDir("parent")
	parent.AddChild(nil)

	if len(parent.Children) != 0 {
		t.Errorf("Expected 0 children when adding nil, got %d", len(parent.Children))
	}
}

func TestHasChildren(t *testing.T) {
	parent := NewDir("parent")
	if parent.HasChildren() {
		t.Error("Expected HasChildren() to return false for empty parent")
	}

	parent.AddChild(NewFile("child.go"))
	if !parent.HasChildren() {
		t.Error("Expected HasChildren() to return true after adding child")
	}
}

func TestNodeTypeString(t *testing.T) {
	tests := []struct {
		nodeType NodeType
		expected string
	}{
		{TypeDir, "DIR"},
		{TypeFile, "FILE"},
		{NodeType(99), "UNKNOWN"},
	}

	for _, tt := range tests {
		if tt.nodeType.String() != tt.expected {
			t.Errorf("NodeType.String() = %s, want %s", tt.nodeType.String(), tt.expected)
		}
	}
}

func TestNodeWithTemplateDirective(t *testing.T) {
	node := NewFile("handler.go")
	node.Template = "go"

	if node.Template != "go" {
		t.Errorf("Expected template 'go', got %s", node.Template)
	}
}

func TestNodeIgnore(t *testing.T) {
	node := NewDir("generated")
	node.Ignore = true

	if !node.Ignore {
		t.Error("Expected Ignore to be true")
	}
}
