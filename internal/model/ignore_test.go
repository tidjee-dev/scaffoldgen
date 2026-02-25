package model

import (
	"testing"
)

func TestNewIgnoreRules(t *testing.T) {
	patterns := []string{"*.tmp", "node_modules", "dist"}
	rules := NewIgnoreRules(patterns)

	if len(rules.rules) != 3 {
		t.Errorf("Expected 3 rules, got %d", len(rules.rules))
	}
}

func TestNewIgnoreRulesWithEmpty(t *testing.T) {
	patterns := []string{"*.tmp", "", "  ", "node_modules"}
	rules := NewIgnoreRules(patterns)

	if len(rules.rules) != 2 {
		t.Errorf("Expected 2 rules (empty strings skipped), got %d", len(rules.rules))
	}
}

func TestNewIgnoreRulesWithComments(t *testing.T) {
	patterns := []string{"*.tmp", "# comment", "node_modules"}
	rules := NewIgnoreRules(patterns)

	if len(rules.rules) != 2 {
		t.Errorf("Expected 2 rules (comments skipped), got %d", len(rules.rules))
	}
}

func TestNewIgnoreRulesWithNegation(t *testing.T) {
	patterns := []string{"*.tmp", "!important.tmp"}
	rules := NewIgnoreRules(patterns)

	if len(rules.rules) != 2 {
		t.Errorf("Expected 2 rules, got %d", len(rules.rules))
	}

	if !rules.rules[1].IsNegation {
		t.Error("Expected second rule to be negation")
	}
}

func TestAddRule(t *testing.T) {
	rules := &IgnoreRules{}
	rules.AddRule("*.tmp")
	rules.AddRule("node_modules")

	if len(rules.rules) != 2 {
		t.Errorf("Expected 2 rules, got %d", len(rules.rules))
	}
}

func TestShouldIgnore(t *testing.T) {
	rules := NewIgnoreRules([]string{"*.tmp", "node_modules", "dist/"})

	tests := []struct {
		path     string
		isDir    bool
		expected bool
	}{
		{"test.tmp", false, true},
		{"test.go", false, false},
		{"node_modules", true, true},
		{"src", true, false},
		{"dist", true, true},
	}

	for _, tt := range tests {
		result := rules.ShouldIgnore(tt.path, tt.isDir)
		if result != tt.expected {
			t.Errorf("ShouldIgnore(%s, %v) = %v, want %v", tt.path, tt.isDir, result, tt.expected)
		}
	}
}

func TestShouldIgnoreWithNegation(t *testing.T) {
	rules := NewIgnoreRules([]string{"*.tmp", "!important.tmp"})

	result := rules.ShouldIgnore("important.tmp", false)
	if result {
		t.Error("Expected important.tmp to not be ignored due to negation rule")
	}
}

func TestShouldIgnoreNilRules(t *testing.T) {
	var rules *IgnoreRules
	result := rules.ShouldIgnore("test.tmp", false)
	if result {
		t.Error("Expected nil rules to return false")
	}
}

func TestShouldIgnoreEmptyRules(t *testing.T) {
	rules := &IgnoreRules{}
	result := rules.ShouldIgnore("test.tmp", false)
	if result {
		t.Error("Expected empty rules to return false")
	}
}

func TestIgnoreRulePattern(t *testing.T) {
	patterns := []string{
		"*.go",
		"vendor/*",
		".git",
		"build/",
		"*.tmp",
	}
	rules := NewIgnoreRules(patterns)

	for i, pattern := range patterns {
		if rules.rules[i].Pattern != pattern {
			t.Errorf("Expected pattern %s, got %s", pattern, rules.rules[i].Pattern)
		}
	}
}
