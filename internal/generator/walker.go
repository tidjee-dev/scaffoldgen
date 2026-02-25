package generator

import (
	"path/filepath"
	"strings"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

type EventKind int

const (
	EventDir EventKind = iota
	EventFile
)

type Event struct {
	Kind     EventKind
	Path     string
	IsGo     bool // Deprecated: use Template instead
	Template TemplateProvider
	Node     *model.Node
	// Enhanced template context for new rendering system
	TemplateContext *TemplateContext
}

func (e Event) Label() string {
	if e.Kind == EventDir {
		return "DIR " + e.Path
	}
	return "FILE " + e.Path
}

type EventHandler func(e Event)

func Walk(root *model.Node, base string, fn EventHandler) {
	walk(root, base, fn, nil)
}

// WalkWithRules walks the tree and respects ignore rules
func WalkWithRules(root *model.Node, base string, fn EventHandler, ignoreRules *model.IgnoreRules) {
	walk(root, base, fn, ignoreRules)
}

func getFileTemplate(n *model.Node) TemplateProvider {
	// If directive is "none", skip templating
	if strings.ToLower(n.Template) == "none" {
		return nil
	}

	// If explicit directive exists, use it
	if n.Template != "" {
		if provider := GetTemplate("." + strings.ToLower(n.Template)); provider != nil {
			return provider
		}
	}

	// Fall back to file extension detection
	return GetTemplate(GetFileExtension(n.Name))
}

// createTemplateContext creates an enhanced template context for a file
func createTemplateContext(n *model.Node, fullPath string) *TemplateContext {
	if !n.IsFile() {
		return nil
	}

	ext := GetFileExtension(n.Name)
	if ext == "" {
		return nil
	}

	// Get language from extension
	language := GetLanguageFromExtension(n.Name)

	// Create context
	ctx := NewTemplateContext(n.Name, "", language, fullPath)

	// Determine package name from parent directory
	if n != nil && n.Name != "" {
		// Use the filename without extension as default
		ctx.Package = strings.TrimSuffix(n.Name, ext)
	}

	return &ctx
}

func walk(n *model.Node, base string, fn EventHandler, ignoreRules *model.IgnoreRules) {
	if n == nil || n.Ignore {
		return
	}

	// Check if this node matches ignore rules
	if ignoreRules != nil {
		if ignoreRules.ShouldIgnore(n.Name, n.IsDir()) {
			return
		}
	}

	current := n.Name
	if base != "" {
		current = filepath.Join(base, n.Name)
	}

	if n.IsDir() {
		fn(Event{Kind: EventDir, Path: current, Node: n})
		for _, c := range n.Children {
			walk(c, current, fn, ignoreRules)
		}
		return
	}

	fn(Event{
		Kind:            EventFile,
		Path:            current,
		IsGo:            strings.HasSuffix(n.Name, ".go"),
		Template:        getFileTemplate(n),
		Node:            n,
		TemplateContext: createTemplateContext(n, current),
	})
}
