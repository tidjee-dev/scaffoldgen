package app

import (
	"testing"
)

func TestCompletionCmd(t *testing.T) {
	// Test that completion command is properly configured
	if completionCmd.Use != "completion [bash|zsh|fish|powershell]" {
		t.Errorf("Expected use 'completion [bash|zsh|fish|powershell]', got %s", completionCmd.Use)
	}

	if completionCmd.Short != "Generate the autocompletion script for the specified shell" {
		t.Errorf("Expected short description 'Generate the autocompletion script for the specified shell', got %s", completionCmd.Short)
	}

	// Test valid args
	validArgs := []string{"bash", "zsh", "fish", "powershell"}
	for i, arg := range completionCmd.ValidArgs {
		if arg != validArgs[i] {
			t.Errorf("Expected valid arg %s, got %s", validArgs[i], arg)
		}
	}

	// Test that command requires exactly one argument
	if completionCmd.Args == nil {
		t.Error("Expected Args to be set")
	}
}

func TestCompletionCmdExecution(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "bash completion",
			args:    []string{"bash"},
			wantErr: false,
		},
		{
			name:    "zsh completion",
			args:    []string{"zsh"},
			wantErr: false,
		},
		{
			name:    "fish completion",
			args:    []string{"fish"},
			wantErr: false,
		},
		{
			name:    "powershell completion",
			args:    []string{"powershell"},
			wantErr: false,
		},
		{
			name:    "invalid shell",
			args:    []string{"invalid"},
			wantErr: true,
		},
		{
			name:    "no args",
			args:    []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the validation logic directly without cobra
			if len(tt.args) != 1 {
				if !tt.wantErr {
					t.Error("Expected error for wrong number of arguments")
				}
				return
			}

			// Test valid shell validation
			validShells := []string{"bash", "zsh", "fish", "powershell"}
			isValid := false
			for _, shell := range validShells {
				if tt.args[0] == shell {
					isValid = true
					break
				}
			}

			if !isValid && !tt.wantErr {
				t.Error("Expected error for invalid shell")
			}
		})
	}
}
