package app

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestVersionVariable(t *testing.T) {
	// Test that Version variable exists and has a value
	if Version == "" {
		t.Error("Expected Version variable to be set")
	}

	// Test that version follows semantic versioning format
	// Note: Version may be set during build, so we just validate it's not empty
	// and follows a reasonable format
	if len(Version) < 3 {
		t.Errorf("Version seems too short: %s", Version)
	}
}

func TestPrintVersion(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call PrintVersion
	PrintVersion()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check that output contains expected content
	expectedPrefix := "scaffoldgen "
	if !strings.HasPrefix(output, expectedPrefix) {
		t.Errorf("Expected output to start with %q, got %q", expectedPrefix, output)
	}

	// Check that it ends with newline
	if !strings.HasSuffix(output, "\n") {
		t.Errorf("Expected output to end with newline, got %q", output)
	}
}

func TestPrintVersionFormat(t *testing.T) {
	// Test that PrintVersion outputs in expected format
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintVersion()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check that output starts with "scaffoldgen"
	if len(output) < len("scaffoldgen") || output[:len("scaffoldgen")] != "scaffoldgen" {
		t.Errorf("Expected output to start with 'scaffoldgen', got %q", output)
	}

	// Check that output contains the version
	if !contains(output, Version) {
		t.Errorf("Expected output to contain version %s, got %q", Version, output)
	}
}

func TestVersionCanBeModified(t *testing.T) {
	// Test that Version variable can be modified (for build processes)
	originalVersion := Version
	defer func() {
		Version = originalVersion // Restore original value
	}()

	// Modify version
	testVersion := "v1.0.0-test"
	Version = testVersion

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintVersion()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expected := "scaffoldgen " + testVersion + "\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
