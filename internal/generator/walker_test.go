package generator

import (
	"testing"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

func TestEventLabel(t *testing.T) {
	tests := []struct {
		event    Event
		expected string
	}{
		{
			event:    Event{Kind: EventDir, Path: "src"},
			expected: "DIR src",
		},
		{
			event:    Event{Kind: EventFile, Path: "main.go"},
			expected: "FILE main.go",
		},
	}

	for _, tt := range tests {
		result := tt.event.Label()
		if result != tt.expected {
			t.Errorf("Event.Label() = %s, want %s", result, tt.expected)
		}
	}
}

func TestWalkerBasic(t *testing.T) {
	root := model.NewDir("project")
	root.AddChild(model.NewDir("src"))
	root.AddChild(model.NewFile("README.md"))

	eventCount := 0
	walker := func(e Event) {
		eventCount++
	}

	Walk(root, ".", walker)

	// Should walk root, src dir, and README.md
	if eventCount != 3 {
		t.Errorf("Expected 3 events, got %d", eventCount)
	}
}

func TestWalkerFilePath(t *testing.T) {
	root := model.NewDir("project")
	src := model.NewDir("src")
	root.AddChild(src)
	src.AddChild(model.NewFile("main.go"))

	events := []Event{}
	walker := func(e Event) {
		events = append(events, e)
	}

	Walk(root, "./project", walker)

	// Check that paths are correctly constructed
	if len(events) > 0 {
		hasGoFile := false
		for _, e := range events {
			if e.Kind == EventFile && e.Node.Name == "main.go" {
				hasGoFile = true
			}
		}
		if !hasGoFile {
			t.Error("Expected to find main.go in events")
		}
	}
}

func TestWalkerWithRules(t *testing.T) {
	root := model.NewDir("project")
	root.AddChild(model.NewDir("src"))
	root.AddChild(model.NewDir("vendor"))
	root.AddChild(model.NewFile("main.go"))

	rules := model.NewIgnoreRules([]string{"vendor"})

	eventCount := 0
	walker := func(e Event) {
		eventCount++
	}

	WalkWithRules(root, ".", walker, rules)

	// Should walk root, src, and main.go (vendor should be skipped)
	if eventCount != 3 {
		t.Errorf("Expected 3 events (vendor excluded), got %d", eventCount)
	}
}

func TestGetFileTemplate(t *testing.T) {
	tests := []struct {
		node        *model.Node
		shouldBeNil bool
		description string
	}{
		{
			node:        model.NewFile("main.go"),
			shouldBeNil: false,
			description: "Should find Go template",
		},
		{
			node:        model.NewFile("utils.py"),
			shouldBeNil: false,
			description: "Should find Python template",
		},
		{
			node:        &model.Node{Name: "README", Type: model.TypeFile, Template: "none"},
			shouldBeNil: true,
			description: "Should return nil for template:none",
		},
		{
			node:        &model.Node{Name: "unknown.xyz", Type: model.TypeFile},
			shouldBeNil: true,
			description: "Should return nil for unknown extension",
		},
	}

	for _, tt := range tests {
		result := getFileTemplate(tt.node)
		isNil := result == nil
		if isNil != tt.shouldBeNil {
			t.Errorf("%s: expected template nil=%v, got nil=%v", tt.description, tt.shouldBeNil, isNil)
		}
	}
}

func TestWalkerNesting(t *testing.T) {
	root := model.NewDir("project")

	src := model.NewDir("src")
	root.AddChild(src)

	handlers := model.NewDir("handlers")
	src.AddChild(handlers)

	handlers.AddChild(model.NewFile("user.go"))
	handlers.AddChild(model.NewFile("admin.go"))

	eventCount := 0
	walker := func(e Event) {
		eventCount++
	}

	Walk(root, ".", walker)

	// Should walk: root, src, handlers, user.go, admin.go = 5 events
	if eventCount != 5 {
		t.Errorf("Expected 5 events for nested structure, got %d", eventCount)
	}
}

func TestWalkerEventKind(t *testing.T) {
	root := model.NewDir("project")
	root.AddChild(model.NewDir("src"))
	root.AddChild(model.NewFile("main.go"))

	dirCount := 0
	fileCount := 0
	walker := func(e Event) {
		if e.Kind == EventDir {
			dirCount++
		} else if e.Kind == EventFile {
			fileCount++
		}
	}

	Walk(root, ".", walker)

	if dirCount != 2 {
		t.Errorf("Expected 2 directories, got %d", dirCount)
	}
	if fileCount != 1 {
		t.Errorf("Expected 1 file, got %d", fileCount)
	}
}
