package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestReadVersionConfig(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	
	// Create test version file
	config := VersionConfig{
		Version: "1.0.0",
		Build: Build{
			Date:   "2023-12-25T10:30:00Z",
			Commit: "abc123",
			Dirty:  false,
		},
		Metadata: Meta{
			GoVersion: "1.21.0",
			Module:    "github.com/test/module",
		},
	}
	
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test config: %v", err)
	}
	
	versionFile := filepath.Join(tempDir, "version.json")
	if err := os.WriteFile(versionFile, data, 0644); err != nil {
		t.Fatalf("Failed to write test version file: %v", err)
	}
	
	// Test reading the config
	result, err := readVersionConfig(versionFile)
	if err != nil {
		t.Errorf("Expected no error reading config, got %v", err)
	}
	
	if result.Version != "1.0.0" {
		t.Errorf("Expected version 1.0.0, got %s", result.Version)
	}
	
	if result.Build.Commit != "abc123" {
		t.Errorf("Expected commit abc123, got %s", result.Build.Commit)
	}
	
	if result.Build.Dirty != false {
		t.Errorf("Expected dirty false, got %v", result.Build.Dirty)
	}
}

func TestReadVersionConfigNotFound(t *testing.T) {
	// Test reading non-existent file
	_, err := readVersionConfig("/non/existent/file.json")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}

func TestWriteVersionConfig(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	
	config := &VersionConfig{
		Version: "2.0.0",
		Build: Build{
			Date:   time.Now().Format(time.RFC3339),
			Commit: "def456",
			Dirty:  true,
		},
		Metadata: Meta{
			GoVersion: "1.22.0",
			Module:    "github.com/test/updated",
		},
	}
	
	versionFile := filepath.Join(tempDir, "version.json")
	
	// Test writing the config
	err := writeVersionConfig(versionFile, config)
	if err != nil {
		t.Errorf("Expected no error writing config, got %v", err)
	}
	
	// Verify the file was written correctly
	data, err := os.ReadFile(versionFile)
	if err != nil {
		t.Fatalf("Failed to read written config file: %v", err)
	}
	
	var writtenConfig VersionConfig
	if err := json.Unmarshal(data, &writtenConfig); err != nil {
		t.Fatalf("Failed to unmarshal written config: %v", err)
	}
	
	if writtenConfig.Version != "2.0.0" {
		t.Errorf("Expected written version 2.0.0, got %s", writtenConfig.Version)
	}
	
	if writtenConfig.Build.Dirty != true {
		t.Errorf("Expected written dirty true, got %v", writtenConfig.Build.Dirty)
	}
}

func TestUpdateGoVersionFile(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	
	versionGoFile := filepath.Join(tempDir, "version.go")
	version := "v1.2.3"
	
	// Test updating Go version file
	err := updateGoVersionFile(versionGoFile, version)
	if err != nil {
		t.Errorf("Expected no error updating Go version file, got %v", err)
	}
	
	// Verify the file was written correctly
	data, err := os.ReadFile(versionGoFile)
	if err != nil {
		t.Fatalf("Failed to read written Go version file: %v", err)
	}
	
	content := string(data)
	
	// Check that the file contains the expected template
	expectedParts := []string{
		"package app",
		"import \"fmt\"",
		"var Version = \"v1.2.3\"",
		"func PrintVersion()",
		"fmt.Println(\"scaffoldgen\", Version)",
	}
	
	for _, part := range expectedParts {
		if !contains(content, part) {
			t.Errorf("Expected Go version file to contain %q, got %q", part, content)
		}
	}
}

func TestGetGitInfo(t *testing.T) {
	// Test in a temporary directory (not a git repo)
	tempDir := t.TempDir()
	
	commit, dirty, err := getGitInfo(tempDir)
	
	// Should return an error since it's not a git repo
	if err == nil {
		t.Error("Expected error for non-git directory")
	}
	
	// Should return empty values on error
	if commit != "" {
		t.Errorf("Expected empty commit on error, got %s", commit)
	}
	
	if dirty != false {
		t.Errorf("Expected dirty false on error, got %v", dirty)
	}
}

func TestVersionConfigStruct(t *testing.T) {
	// Test that the struct can be marshaled/unmarshaled correctly
	config := VersionConfig{
		Version: "1.0.0",
		Build: Build{
			Date:   "2023-12-25T10:30:00Z",
			Commit: "abc123",
			Dirty:  false,
		},
		Metadata: Meta{
			GoVersion: "1.21.0",
			Module:    "github.com/test/module",
		},
	}
	
	// Marshal to JSON
	data, err := json.Marshal(config)
	if err != nil {
		t.Errorf("Expected no error marshaling config, got %v", err)
	}
	
	// Unmarshal from JSON
	var unmarshaled VersionConfig
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Errorf("Expected no error unmarshaling config, got %v", err)
	}
	
	// Verify all fields match
	if unmarshaled.Version != config.Version {
		t.Errorf("Expected version %s, got %s", config.Version, unmarshaled.Version)
	}
	
	if unmarshaled.Build.Date != config.Build.Date {
		t.Errorf("Expected build date %s, got %s", config.Build.Date, unmarshaled.Build.Date)
	}
	
	if unmarshaled.Metadata.GoVersion != config.Metadata.GoVersion {
		t.Errorf("Expected go version %s, got %s", config.Metadata.GoVersion, unmarshaled.Metadata.GoVersion)
	}
}

func TestBuildStruct(t *testing.T) {
	// Test Build struct
	build := Build{
		Date:   "2023-12-25T10:30:00Z",
		Commit: "abc123",
		Dirty:  true,
	}
	
	// Marshal to JSON
	data, err := json.Marshal(build)
	if err != nil {
		t.Errorf("Expected no error marshaling build, got %v", err)
	}
	
	// Unmarshal from JSON
	var unmarshaled Build
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Errorf("Expected no error unmarshaling build, got %v", err)
	}
	
	// Verify all fields match
	if unmarshaled.Date != build.Date {
		t.Errorf("Expected date %s, got %s", build.Date, unmarshaled.Date)
	}
	
	if unmarshaled.Commit != build.Commit {
		t.Errorf("Expected commit %s, got %s", build.Commit, unmarshaled.Commit)
	}
	
	if unmarshaled.Dirty != build.Dirty {
		t.Errorf("Expected dirty %v, got %v", build.Dirty, unmarshaled.Dirty)
	}
}

func TestMetaStruct(t *testing.T) {
	// Test Meta struct
	meta := Meta{
		GoVersion: "1.21.0",
		Module:    "github.com/test/module",
	}
	
	// Marshal to JSON
	data, err := json.Marshal(meta)
	if err != nil {
		t.Errorf("Expected no error marshaling meta, got %v", err)
	}
	
	// Unmarshal from JSON
	var unmarshaled Meta
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Errorf("Expected no error unmarshaling meta, got %v", err)
	}
	
	// Verify all fields match
	if unmarshaled.GoVersion != meta.GoVersion {
		t.Errorf("Expected go version %s, got %s", meta.GoVersion, unmarshaled.GoVersion)
	}
	
	if unmarshaled.Module != meta.Module {
		t.Errorf("Expected module %s, got %s", meta.Module, unmarshaled.Module)
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || 
		(len(s) > len(substr) && indexOf(s, substr) >= 0))
}

// Simple indexOf implementation for substring search
func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
