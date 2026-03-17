package generator

import (
	"strings"
	"testing"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

func TestExportMarkdown(t *testing.T) {
	// Create a test node structure
	root := model.NewDir("testproject")

	// Add files and directories
	srcDir := model.NewDir("src")
	srcDir.AddChild(model.NewFile("main.go"))
	srcDir.AddChild(model.NewFile("config.go"))

	pkgDir := model.NewDir("pkg")
	loggerDir := model.NewDir("logger")
	loggerDir.AddChild(model.NewFile("logger.go"))
	pkgDir.AddChild(loggerDir)

	root.AddChild(srcDir)
	root.AddChild(pkgDir)

	// Test markdown export
	result := ExportMarkdown(root)

	// Check basic structure
	if !strings.Contains(result, "# testproject") {
		t.Error("Expected markdown to contain root title")
	}

	if !strings.Contains(result, "src/") {
		t.Error("Expected markdown to contain src directory")
	}

	if !strings.Contains(result, "pkg/") {
		t.Error("Expected markdown to contain pkg directory")
	}

	if !strings.Contains(result, "- main.go") {
		t.Error("Expected markdown to contain main.go file")
	}

	if !strings.Contains(result, "- logger.go") {
		t.Error("Expected markdown to contain logger.go file")
	}
}

func TestExportMarkdownEmpty(t *testing.T) {
	// Test with empty directory
	root := model.NewDir("empty")
	result := ExportMarkdown(root)

	expected := "# empty\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestExportMarkdownSingleFile(t *testing.T) {
	// Test with single file
	root := model.NewDir("single")
	root.AddChild(model.NewFile("test.go"))
	result := ExportMarkdown(root)

	if !strings.Contains(result, "# single") {
		t.Error("Expected markdown to contain root title")
	}

	if !strings.Contains(result, "- test.go") {
		t.Error("Expected markdown to contain test.go file")
	}
}

func TestExportJSON(t *testing.T) {
	// Create a test node structure
	root := model.NewDir("testproject")
	srcDir := model.NewDir("src")
	srcDir.AddChild(model.NewFile("main.go"))
	root.AddChild(srcDir)

	// Test JSON export
	result, err := ExportJSON(root)
	if err != nil {
		t.Fatalf("ExportJSON failed: %v", err)
	}

	// Check that it's valid JSON and contains expected content
	// Note: ExportJSON doesn't include the root name, only children
	if !strings.Contains(result, "src") {
		t.Error("Expected JSON to contain src directory")
	}

	if !strings.Contains(result, "main.go") {
		t.Error("Expected JSON to contain main.go file")
	}
}

func TestExportYAML(t *testing.T) {
	// Create a test node structure
	root := model.NewDir("testproject")
	srcDir := model.NewDir("src")
	srcDir.AddChild(model.NewFile("main.go"))
	root.AddChild(srcDir)

	// Test YAML export
	result, err := ExportYAML(root)
	if err != nil {
		t.Fatalf("ExportYAML failed: %v", err)
	}

	// Check that it contains expected content
	// Note: ExportYAML doesn't include the root name, only children
	if !strings.Contains(result, "src") {
		t.Error("Expected YAML to contain src directory")
	}

	if !strings.Contains(result, "main.go") {
		t.Error("Expected YAML to contain main.go file")
	}
}

func TestExportMarkdownNestedStructure(t *testing.T) {
	// Test deeply nested structure
	root := model.NewDir("deep")
	level1 := model.NewDir("level1")
	level2 := model.NewDir("level2")
	level3 := model.NewDir("level3")

	level3.AddChild(model.NewFile("deep.go"))
	level2.AddChild(level3)
	level1.AddChild(level2)
	root.AddChild(level1)

	result := ExportMarkdown(root)

	// Check indentation and structure
	lines := strings.Split(result, "\n")

	var foundLevel1, foundLevel2, foundLevel3 bool
	for _, line := range lines {
		if strings.Contains(line, "level1/") {
			foundLevel1 = true
		}
		if strings.Contains(line, "level2/") {
			foundLevel2 = true
		}
		if strings.Contains(line, "level3/") {
			foundLevel3 = true
		}
	}

	if !foundLevel1 {
		t.Error("Expected to find level1 directory")
	}
	if !foundLevel2 {
		t.Error("Expected to find level2 directory")
	}
	if !foundLevel3 {
		t.Error("Expected to find level3 directory")
	}
}
