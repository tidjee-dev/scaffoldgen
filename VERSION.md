# Version Management

This project uses an automated version management system to keep version information consistent across all files.

## Files Managed

- `Makefile` - Contains the VERSION variable used for builds
- `internal/app/version.go` - Go version variable used by the CLI (updated during build)

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

## Current Version

The current version is defined in the Makefile:

```makefile
VERSION ?= 1.1.2
```

## Build Process

Always build after version updates:

```bash
make build-with-version
```

This ensures the binary contains the latest version information by passing the version via ldflags to the Go binary.

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

# Check Makefile version
grep "VERSION" Makefile
```
