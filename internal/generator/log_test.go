package generator

import (
	"testing"
)

func TestLogVerbose(t *testing.T) {
	// Test logVerbose with verbose disabled (no output expected)
	logVerbose(false, "DIR", "src")
	// Should not panic
}

func TestLogVerboseEnabled(t *testing.T) {
	// Test logVerbose with verbose enabled
	// Output will be printed but test doesn't verify it directly
	logVerbose(true, "DIR", "src")
	logVerbose(true, "FILE", "main.go")
}
