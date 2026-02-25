---
title: Installation
sidebar_position: 1
---

Scaffoldgen is a lightweight Go CLI that generates safe, reviewable scaffold scripts from project structures.

## Prerequisites

- **Go 1.25+** (for building from source)
- **macOS / Linux / Windows**

## Install from Binary

The easiest way to get started is to download a pre-built binary from the [GitHub Releases](https://github.com/tidjee-dev/scaffoldgen/releases) page.

## Build from Source

If you have Go installed:

```bash
go install github.com/tidjee-dev/scaffoldgen/cmd/scaffoldgen@latest
```

This installs the `scaffoldgen` binary to your `$GOPATH/bin` directory.

## Verify Installation

Check that the installation worked:

```bash
scaffoldgen version
```

You should see the version number printed to the console.

## Next Steps

- [Quick Start](./quick-start.md) - Generate your first scaffold
- [Examples](../examples/examples.md) - See real-world use cases
