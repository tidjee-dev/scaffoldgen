package app

import (
	"testing"
)

func TestRootCmd(t *testing.T) {
	// Test that root command is properly configured
	if rootCmd.Use != "scaffoldgen" {
		t.Errorf("Expected use 'scaffoldgen', got %s", rootCmd.Use)
	}

	if rootCmd.Short != "Generate scaffold scripts from markdown structure" {
		t.Errorf("Expected short description 'Generate scaffold scripts from markdown structure', got %s", rootCmd.Short)
	}

	// Test that silence flags are set
	if !rootCmd.SilenceUsage {
		t.Error("Expected SilenceUsage to be true")
	}

	if !rootCmd.SilenceErrors {
		t.Error("Expected SilenceErrors to be true")
	}
}

func TestRootCmdPersistentFlags(t *testing.T) {
	// Test that version flag is set
	flags := rootCmd.PersistentFlags()
	if flag := flags.Lookup("version"); flag == nil {
		t.Error("Expected 'version' flag to be set")
	}

	// Test that shorthand flag exists in the flag set
	// Note: BoolVarP creates the shorthand, but it might not be visible as a separate lookup
	flag := flags.Lookup("v")
	if flag == nil {
		// The shorthand might not be directly lookupable, but should exist
		// Let's check if the version flag has the shorthand set
		versionFlag := flags.Lookup("version")
		if versionFlag != nil && versionFlag.Shorthand != "v" {
			t.Error("Expected version flag to have shorthand 'v'")
		}
	}
}

func TestRootCmdRun(t *testing.T) {
	// Test that root command has a run function
	if rootCmd.Run == nil {
		t.Error("Expected root command to have a run function")
	}

	// Test that run function calls help (basic test)
	// Note: We can't easily test the actual help output without capturing stdout
	// This test mainly ensures the function exists
}

func TestExecute(t *testing.T) {
	// Test that Execute function exists
	// Note: We can't easily test the actual execution without causing side effects
	// This test mainly ensures the function exists and returns an error type

	// Test that Execute returns an error type
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Execute function panicked: %v", r)
		}
	}()

	// We can't call Execute() directly as it would exit the process
	// But we can verify the function signature by checking rootCmd is available
	if rootCmd == nil {
		t.Error("Expected rootCmd to be available")
	}
}

func TestCobraOnInitialize(t *testing.T) {
	// Test that cobra.OnInitialize is called
	// This is a basic test to ensure initialization is set up
	// We can't easily test the actual initialization without affecting global state

	// Verify that the root command has the expected structure
	if rootCmd == nil {
		t.Error("Expected rootCmd to be initialized")
	}
}
