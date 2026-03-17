---
title: Installation
sidebar_position: 1
---

Scaffoldgen is a lightweight Go CLI that generates safe, reviewable scaffold scripts from project structures.

## 🚀 Installation Methods

### **Option 1: Go Install (Recommended)**

```bash
go install github.com/tidjee-dev/scaffoldgen/cmd/scaffoldgen@latest
```

The standard Go way. Works for anyone with Go installed.

### **Option 2: Download Binary (No Go required)**

If you don't have Go installed, download pre-compiled binaries from [GitHub Releases](https://github.com/tidjee-dev/scaffoldgen/releases/latest):

```bash
# macOS (Apple Silicon)
curl -L -o scaffoldgen https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/scaffoldgen-darwin-arm64
chmod +x scaffoldgen
./scaffoldgen --version

# macOS (Intel)
curl -L -o scaffoldgen https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/scaffoldgen-darwin-amd64
chmod +x scaffoldgen
./scaffoldgen --version

# Linux (AMD64)
curl -L -o scaffoldgen https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/scaffoldgen-linux-amd64
chmod +x scaffoldgen
./scaffoldgen --version

# Linux (ARM64)
curl -L -o scaffoldgen https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/scaffoldgen-linux-arm64
chmod +x scaffoldgen
./scaffoldgen --version

# Windows (AMD64)
curl -L -o scaffoldgen.exe https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/scaffoldgen-windows-amd64.exe
./scaffoldgen.exe --version
```

### **Option 3: Build from Source**

For development or if you want the latest bleeding-edge version:

```bash
git clone https://github.com/tidjee-dev/scaffoldgen.git
cd scaffoldgen
make build
./scaffoldgen --version

# Or install locally for development:
go install ./cmd/scaffoldgen
```

## Prerequisites

- **Go 1.25+** (only required for Option 1 and 3)
- **macOS / Linux / Windows**

## Verify Installation

Check that the installation worked:

```bash
scaffoldgen version
```

You should see the version number printed to the console.

## Next Steps

- [Quick Start](./quick-start.md) - Generate your first scaffold
- [Examples](../examples/examples.md) - See real-world use cases
