package input

import (
	"strings"
	"testing"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

func TestParseMarkdownBasic(t *testing.T) {
	markdown := `# project
- src/
  - main.go
  - utils.go
- README.md`

	r := strings.NewReader(markdown)
	root, err := ParseMarkdown(r)

	if err != nil {
		t.Errorf("ParseMarkdown should not error: %v", err)
	}

	if root == nil {
		t.Fatal("Root node should not be nil")
	}

	if root.Name != "project" {
		t.Errorf("Root name should be 'project', got %s", root.Name)
	}

	if len(root.Children) != 2 {
		t.Errorf("Root should have 2 children, got %d", len(root.Children))
	}
}

func TestParseMarkdownMissingRoot(t *testing.T) {
	markdown := `- src/
  - main.go`

	r := strings.NewReader(markdown)
	_, err := ParseMarkdown(r)

	if err == nil {
		t.Error("ParseMarkdown should error when root is missing")
	}
}

func TestParseMarkdownMultipleRoots(t *testing.T) {
	markdown := `# project1
- src/

# project2
- src/`

	r := strings.NewReader(markdown)
	_, err := ParseMarkdown(r)

	if err == nil {
		t.Error("ParseMarkdown should error with multiple roots")
	}
}

func TestParseMarkdownWithIndentation(t *testing.T) {
	markdown := `# backend
- cmd/
  - api/
    - main.go
  - worker/
    - main.go`

	r := strings.NewReader(markdown)
	root, err := ParseMarkdown(r)

	if err != nil {
		t.Errorf("ParseMarkdown should not error: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}

	// Check nested structure
	if len(root.Children) != 1 {
		t.Fatalf("Expected 1 child at root, got %d", len(root.Children))
	}

	cmd := root.Children[0]
	if cmd.Name != "cmd" {
		t.Errorf("Expected 'cmd' dir, got %s", cmd.Name)
	}

	if len(cmd.Children) != 2 {
		t.Errorf("Expected 2 children in cmd, got %d", len(cmd.Children))
	}
}

func TestParseMarkdownWithIgnore(t *testing.T) {
	markdown := `# project
- src/
- vendor/ #ignore`

	r := strings.NewReader(markdown)
	root, err := ParseMarkdown(r)

	if err != nil {
		t.Errorf("ParseMarkdown should not error: %v", err)
	}

	// Check for ignore flag
	vendor := findNode(root, "vendor")
	if vendor == nil {
		t.Error("vendor node should exist")
	} else if !vendor.Ignore {
		t.Error("vendor should be marked as ignored")
	}
}

func TestParseMarkdownWithTemplate(t *testing.T) {
	markdown := `# project
- src/
  - handler.go #template:go
  - utils.py #template:python`

	r := strings.NewReader(markdown)
	root, err := ParseMarkdown(r)

	if err != nil {
		t.Errorf("ParseMarkdown should not error: %v", err)
	}

	// Check for template directive
	src := findNode(root, "src")
	if src != nil {
		handler := findNode(src, "handler.go")
		if handler == nil {
			t.Error("handler.go should exist")
		} else if handler.Template != "go" {
			t.Errorf("handler.go template should be 'go', got %s", handler.Template)
		}
	}
}

func TestParseMarkdownEmptyLines(t *testing.T) {
	markdown := `# project

- src/

  - main.go

- README.md`

	r := strings.NewReader(markdown)
	root, err := ParseMarkdown(r)

	if err != nil {
		t.Errorf("ParseMarkdown should handle empty lines: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}
}

func TestParseMarkdownInvalidIndentation(t *testing.T) {
	markdown := `# project
- src/
 - main.go`

	r := strings.NewReader(markdown)
	_, err := ParseMarkdown(r)

	if err == nil {
		t.Error("ParseMarkdown should error on invalid indentation")
	}
}

func TestParseMarkdownTabs(t *testing.T) {
	markdown := `# project
- src/
	- main.go` // Tab instead of spaces

	r := strings.NewReader(markdown)
	_, err := ParseMarkdown(r)

	if err == nil {
		t.Error("ParseMarkdown should error when tabs are used")
	}
}

func TestParseMarkdownRemovesTrailingSlash(t *testing.T) {
	markdown := `# project
- src/`

	r := strings.NewReader(markdown)
	root, err := ParseMarkdown(r)

	if err != nil {
		t.Errorf("ParseMarkdown should not error: %v", err)
	}

	if len(root.Children) > 0 {
		child := root.Children[0]
		if strings.HasSuffix(child.Name, "/") {
			t.Errorf("Trailing slash should be removed, got %s", child.Name)
		}
	}
}

func TestParseMarkdownNodeTypes(t *testing.T) {
	markdown := `# project
- src/
  - main.go
- README.md`

	r := strings.NewReader(markdown)
	root, err := ParseMarkdown(r)

	if err != nil {
		t.Fatalf("ParseMarkdown should not error: %v", err)
	}

	src := findNode(root, "src")
	if src == nil {
		t.Fatal("src not found")
	}
	if !src.IsDir() {
		t.Error("src should be a directory")
	}

	mainGo := findNode(src, "main.go")
	if mainGo == nil {
		t.Fatal("main.go not found")
	}
	if !mainGo.IsFile() {
		t.Error("main.go should be a file")
	}
}

func findNode(root *model.Node, name string) *model.Node {
	if root == nil {
		return nil
	}
	for _, child := range root.Children {
		if child.Name == name {
			return child
		}
		if result := findNode(child, name); result != nil {
			return result
		}
	}
	return nil
}
