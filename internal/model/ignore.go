package model

import (
	"path/filepath"
	"strings"
)

// IgnoreRule represents a single ignore pattern
type IgnoreRule struct {
	Pattern    string // glob pattern like "*.tmp", "node_modules", "*.go"
	IsNegation bool   // ! prefix negates the rule
}

// IgnoreRules manages multiple ignore patterns
type IgnoreRules struct {
	rules []IgnoreRule
}

// NewIgnoreRules creates a new ignore rules collection
func NewIgnoreRules(patterns []string) *IgnoreRules {
	ir := &IgnoreRules{}
	for _, p := range patterns {
		ir.AddRule(p)
	}
	return ir
}

// AddRule adds a single ignore pattern
func (ir *IgnoreRules) AddRule(pattern string) {
	pattern = strings.TrimSpace(pattern)
	if pattern == "" || strings.HasPrefix(pattern, "#") {
		return
	}

	isNegation := strings.HasPrefix(pattern, "!")
	if isNegation {
		pattern = strings.TrimPrefix(pattern, "!")
		pattern = strings.TrimSpace(pattern)
	}

	ir.rules = append(ir.rules, IgnoreRule{
		Pattern:    pattern,
		IsNegation: isNegation,
	})
}

// ShouldIgnore checks if a path matches any ignore rules
// Returns true if the path should be ignored (considering negation rules)
func (ir *IgnoreRules) ShouldIgnore(path string, isDir bool) bool {
	if ir == nil || len(ir.rules) == 0 {
		return false
	}

	ignored := false
	for _, rule := range ir.rules {
		if matchesPattern(path, rule.Pattern, isDir) {
			if rule.IsNegation {
				ignored = false
			} else {
				ignored = true
			}
		}
	}

	return ignored
}

// matchesPattern checks if path matches a glob pattern
func matchesPattern(path, pattern string, isDir bool) bool {
	// Normalize path separators
	path = filepath.ToSlash(path)
	pattern = filepath.ToSlash(pattern)

	// Check exact match
	if path == pattern {
		return true
	}

	// Check basename match
	basename := filepath.Base(path)
	matched, err := filepath.Match(pattern, basename)
	if err == nil && matched {
		return true
	}

	// Check full path match
	matched, err = filepath.Match(pattern, path)
	if err == nil && matched {
		return true
	}

	// Check directory-specific match
	if isDir && !strings.HasSuffix(pattern, "/") {
		// Pattern "node_modules" should match "node_modules/" directories
		matched, err := filepath.Match(pattern, basename)
		if err == nil && matched {
			return true
		}
	}

	// Check if pattern contains path separator
	if strings.Contains(pattern, "/") && !strings.HasSuffix(path, "/") {
		// Add trailing slash for directory matching
		if isDir {
			path += "/"
			matched, err = filepath.Match(pattern, path)
			if err == nil && matched {
				return true
			}
		}
	}

	return false
}
