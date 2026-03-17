# Version Management

This project uses an automated version management system to keep version information consistent across all files.

## Files Managed

- `version.json` - Master version configuration with build metadata
- `internal/app/version.go` - Go version variable used by the CLI

## Version Update Methods

### 1. Manual Version Update

```bash
# Set specific version
make version VERSION=1.2.3

# Or use the script directly
go run scripts/update-version.go 1.2.3
```

### 2. Semantic Version Bumping

```bash
# Patch version (1.0.0 -> 1.0.1)
make bump-patch

# Minor version (1.0.0 -> 1.1.0)  
make bump-minor

# Major version (1.0.0 -> 2.0.0)
make bump-major
```

### 3. Automatic Updates

The version is automatically updated when:
- Creating a Git tag with `v*` (via GitHub Actions)
- Running release workflows

## Version Information

The `version.json` file contains:

```json
{
  "version": "1.0.2",
  "build": {
    "date": "2026-03-17T11:51:56+01:00",
    "commit": "7d8e8d8f0009e1509863ffb6b3f6a1d9d01d8209",
    "dirty": true
  },
  "metadata": {
    "go_version": "1.25.6",
    "module": "github.com/tidjee-dev/scaffoldgen"
  }
}
```

## Build Process

Always build after version updates:

```bash
make build-with-version
```

This ensures the binary contains the latest version information.

## Release Process

1. Update version: `make bump-patch` (or minor/major)
2. Review changes: `git diff`
3. Commit and tag: The bump script does this automatically
4. Push: `git push && git push --tags`
5. GitHub Actions will build and create releases automatically

## Verification

```bash
# Check current version
./scaffoldgen --version

# Check version config
cat version.json
```
