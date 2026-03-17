---
sidebar_position: 5
title: Version Command
---

The `version` command displays version information for scaffoldgen.

## Basic Usage

Show the current version:

```bash
scaffoldgen version
```

Output:

```
scaffoldgen 1.0.2
```

## Verbose Mode

Use the `--verbose` flag to see detailed build information including commit hash, build date, and Go runtime details:

```bash
scaffoldgen version --verbose
```

Example output:

```
scaffoldgen 1.0.2

Build Information:
  Commit: 7d8e8d8f (dirty)
  Built:  2026-03-17 11:54:01

Runtime:
  Go:     1.25.6
  Module: github.com/tidjee-dev/scaffoldgen
```

## Version Information Fields

- **Version**: The semantic version number (e.g., 1.0.2)
- **Commit**: Git commit hash (truncated to 8 characters)
- **Dirty**: Indicates if the build was made from a working directory with uncommitted changes
- **Built**: Timestamp when the binary was built
- **Go**: Go version used to build the binary
- **Module**: Go module path

## Short Flag

You can use the short flag `-V` instead of `--verbose`:

```bash
scaffoldgen version -V
```

## Global Version Flag

You can also use the global `--version` flag for a quick version check:

```bash
scaffoldgen --version
```

This is equivalent to `scaffoldgen version` but doesn't support verbose output.
