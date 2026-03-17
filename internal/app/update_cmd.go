package app

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tidjee-dev/scaffoldgen/internal/tui"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update scaffoldgen to the latest version",
	Long: `Update scaffoldgen to the latest version from GitHub releases.

This command will:
1. Check for the latest available version
2. Download the appropriate binary for your platform
3. Replace the current binary
4. Verify the update`,
	RunE: runUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func runUpdate(cmd *cobra.Command, args []string) error {
	fmt.Println(tui.Header("🔄 Checking for updates..."))

	// Get current version
	fmt.Printf("Current version: %s\n", Version)

	// Get latest release
	latest, err := getLatestRelease()
	if err != nil {
		return fmt.Errorf("failed to check for updates: %w", err)
	}

	fmt.Printf("Latest version: %s\n", latest.TagName)

	// Check if update is needed
	if !isNewerVersion(Version, latest.TagName) {
		fmt.Println(tui.Success("✅ You already have the latest version!"))
		return nil
	}

	fmt.Println(tui.Info("📦 Updating to latest version..."))

	// Determine binary URL and path
	binaryURL, binaryPath, err := getBinaryInfo()
	if err != nil {
		return fmt.Errorf("failed to determine binary info: %w", err)
	}

	// Download latest binary
	if err := downloadBinary(binaryURL, binaryPath); err != nil {
		return fmt.Errorf("failed to download binary: %w", err)
	}

	// Make binary executable
	if runtime.GOOS != "windows" {
		if err := os.Chmod(binaryPath, 0755); err != nil {
			return fmt.Errorf("failed to make binary executable: %w", err)
		}
	}

	// Verify update
	if err := verifyUpdate(binaryPath); err != nil {
		return fmt.Errorf("update verification failed: %w", err)
	}

	fmt.Println(tui.Success("✅ Update completed successfully!"))
	fmt.Printf("Updated from %s to %s\n", Version, latest.TagName)
	fmt.Println(tui.Info("💡 Please restart scaffoldgen to use the new version"))

	return nil
}

// getBinaryInfo returns the download URL and local binary path
func getBinaryInfo() (string, string, error) {
	// Determine platform
	goos := runtime.GOOS
	arch := runtime.GOARCH

	// Map architecture names
	switch arch {
	case "amd64":
		arch = "amd64"
	case "arm64":
		arch = "arm64"
	default:
		return "", "", fmt.Errorf("unsupported architecture: %s", arch)
	}

	// Determine binary name and URL
	var binaryName, binaryURL string
	switch goos {
	case "darwin":
		binaryName = "scaffoldgen-darwin-" + arch
		binaryURL = fmt.Sprintf("https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/%s", binaryName)
	case "linux":
		binaryName = "scaffoldgen-linux-" + arch
		binaryURL = fmt.Sprintf("https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/%s", binaryName)
	case "windows":
		binaryName = "scaffoldgen-windows-" + arch + ".exe"
		binaryURL = fmt.Sprintf("https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/%s", binaryName)
	default:
		return "", "", fmt.Errorf("unsupported operating system: %s", goos)
	}

	// Get current binary path
	exePath, err := os.Executable()
	if err != nil {
		return "", "", fmt.Errorf("failed to get current binary path: %w", err)
	}

	return binaryURL, exePath, nil
}

// downloadBinary downloads the binary from the given URL
func downloadBinary(url, destPath string) error {
	fmt.Printf("Downloading from: %s\n", url)

	// Use curl for downloading (more reliable than Go HTTP client for binaries)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-Command", "Invoke-WebRequest", "-Uri", url, "-OutFile", destPath)
	} else {
		cmd = exec.Command("curl", "-L", "-o", destPath, url)
	}

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("download failed: %v\nOutput: %s", err, string(output))
	}

	return nil
}

// verifyUpdate verifies the updated binary
func verifyUpdate(binaryPath string) error {
	// Try to run the updated binary with --version
	cmd := exec.Command(binaryPath, "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("updated binary failed to run: %v\nOutput: %s", err, string(output))
	}

	// Check that output contains version information
	if !strings.Contains(string(output), "scaffoldgen") {
		return fmt.Errorf("updated binary output is invalid: %s", string(output))
	}

	return nil
}
