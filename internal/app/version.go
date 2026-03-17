package app

import (
	"fmt"
)

// These will be set by build flags
var (
	Version     = "1.1.0" // Set during build
	BuildDate   = ""      // Set during build
	BuildCommit = ""      // Set during build
	GoVersion   = ""      // Set during build
)

type Build struct {
	Date   string
	Commit string
	Dirty  bool
}

type Meta struct {
	GoVersion string
	Module    string
}

type BuildInfo struct {
	Version  string
	Build    Build
	Metadata Meta
}

func PrintVersion() {
	fmt.Printf("scaffoldgen %s\n", Version)
}

// getEmbeddedBuildInfo returns build info from build flags
func getEmbeddedBuildInfo() *BuildInfo {
	return &BuildInfo{
		Version: Version,
		Build: Build{
			Date:   BuildDate,
			Commit: BuildCommit,
			Dirty:  false, // We could detect this during build if needed
		},
		Metadata: Meta{
			GoVersion: GoVersion,
			Module:    "github.com/tidjee-dev/scaffoldgen",
		},
	}
}
