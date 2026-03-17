package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tidjee-dev/scaffoldgen/internal/tui"
)

const (
	// GitHub API URL for releases
	githubAPI = "https://api.github.com/repos/tidjee-dev/scaffoldgen/releases/latest"
	// File to store last update check
	lastCheckFile = ".scaffoldgen_update_check"
	// Check interval (24 hours)
	checkInterval = 24 * time.Hour
)

// GitHubRelease represents a GitHub release
type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Body    string `json:"body"`
}

// checkForUpdates checks if a new version is available
func checkForUpdates() {
	// Don't check if we just checked recently
	if shouldSkipCheck() {
		return
	}

	// Get latest release from GitHub
	latest, err := getLatestRelease()
	if err != nil {
		// Silent fail - don't bother users with network errors
		return
	}

	// Compare versions
	current := Version
	if isNewerVersion(current, latest.TagName) {
		fmt.Fprintf(os.Stderr, "\n%s %s\n", tui.Info("🔄 Update available:"), latest.TagName)
		fmt.Fprintf(os.Stderr, "%s Run '%s' to update\n", tui.Info("💡"), "scaffoldgen update")
		fmt.Fprintf(os.Stderr, "%s %s\n", tui.Dim("Release notes:"), strings.Split(latest.Body, "\n")[0])
		fmt.Println()
	}

	// Update last check time
	updateLastCheck()
}

// shouldSkipCheck returns true if we should skip the update check
func shouldSkipCheck() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return true
	}

	checkFile := filepath.Join(homeDir, lastCheckFile)
	info, err := os.Stat(checkFile)
	if err != nil {
		return false // No check file, proceed with check
	}

	// Skip if last check was less than 24 hours ago
	return time.Since(info.ModTime()) < checkInterval
}

// updateLastCheck updates the timestamp of the last check
func updateLastCheck() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	checkFile := filepath.Join(homeDir, lastCheckFile)
	_ = os.WriteFile(checkFile, []byte(time.Now().Format(time.RFC3339)), 0644)
}

// getLatestRelease fetches the latest release from GitHub API
func getLatestRelease() (*GitHubRelease, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	
	resp, err := client.Get(githubAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &release, nil
}

// isNewerVersion compares current and latest version strings
func isNewerVersion(current, latest string) bool {
	// Remove 'v' prefix if present
	current = strings.TrimPrefix(current, "v")
	latest = strings.TrimPrefix(latest, "v")

	// Simple semantic version comparison
	currentParts := strings.Split(current, ".")
	latestParts := strings.Split(latest, ".")

	for i := 0; i < 3; i++ {
		if i >= len(latestParts) {
			return false // Current version has more parts
		}
		if i >= len(currentParts) {
			return true // Latest version has more parts
		}

		var currentNum, latestNum int
		_, err := fmt.Sscanf(currentParts[i], "%d", &currentNum)
		if err != nil {
			currentNum = 0
		}
		_, err = fmt.Sscanf(latestParts[i], "%d", &latestNum)
		if err != nil {
			latestNum = 0
		}

		if latestNum > currentNum {
			return true
		}
		if latestNum < currentNum {
			return false
		}
	}

	return false // Versions are equal
}
