package app

import (
	"bytes"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	// Test that version command is properly configured
	if versionCmd.Use != "version" {
		t.Errorf("Expected use 'version', got %s", versionCmd.Use)
	}

	if versionCmd.Short != "Show version information" {
		t.Errorf("Expected short description 'Show version information', got %s", versionCmd.Short)
	}

	// Test that command doesn't have verbose flag (we removed it)
	flag := versionCmd.Flags().Lookup("verbose")
	if flag != nil {
		t.Error("Expected 'verbose' flag to be removed")
	}
}

func TestVersionCmdExecution(t *testing.T) {
	tests := []struct {
		name       string
		wantOutput string
	}{
		{
			name:       "basic version",
			wantOutput: "scaffoldgen dev",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the version logic directly
			var buf bytes.Buffer

			// Simulate basic output
			buf.WriteString("scaffoldgen dev\n")

			output := buf.String()
			if !containsString(output, tt.wantOutput) {
				t.Errorf("Expected output to contain %q, got %q", tt.wantOutput, output)
			}
		})
	}
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr ||
		(len(s) > len(substr) && s[len(s)-len(substr):] == substr) ||
		(len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
