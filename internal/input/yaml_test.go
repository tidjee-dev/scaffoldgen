package input

import (
	"strings"
	"testing"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

func TestParseYAMLBasic(t *testing.T) {
	yaml := `backend:
  src:`
	r := strings.NewReader(yaml)
	root, err := ParseYAML(r)

	if err != nil {
		t.Errorf("ParseYAML should not error: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}

	if root.Name != "backend" {
		t.Errorf("Root name should be 'backend', got %s", root.Name)
	}
}

func TestParseYAMLSimpleStructure(t *testing.T) {
	yaml := `backend:
  src:
  tests:`

	r := strings.NewReader(yaml)
	root, err := ParseYAML(r)

	if err != nil {
		t.Errorf("ParseYAML should not error: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}

	if root.Name != "backend" {
		t.Errorf("Expected 'backend', got %s", root.Name)
	}
}

func TestParseYAMLWithFileList(t *testing.T) {
	yaml := `project:
  src:
    - main.go
    - utils.go`

	r := strings.NewReader(yaml)
	root, err := ParseYAML(r)

	if err != nil {
		t.Errorf("ParseYAML should not error: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}
}

func TestParseYAMLNestedStructure(t *testing.T) {
	yaml := `backend:
  internal:
    domain:`

	r := strings.NewReader(yaml)
	root, err := ParseYAML(r)

	if err != nil {
		t.Errorf("ParseYAML should not error: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}

	if root.Name != "backend" {
		t.Errorf("Expected 'backend', got %s", root.Name)
	}
}

func TestParseYAMLRootNode(t *testing.T) {
	yaml := `myapp:
  component1:
  component2:`

	r := strings.NewReader(yaml)
	root, err := ParseYAML(r)

	if err != nil {
		t.Errorf("ParseYAML should not error: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}

	if root.Name != "myapp" {
		t.Errorf("Root should be 'myapp', got %s", root.Name)
	}
}

func findChild(root *model.Node, name string) *model.Node {
	if root == nil {
		return nil
	}
	for _, child := range root.Children {
		if child.Name == name {
			return child
		}
	}
	return nil
}
