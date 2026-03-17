# `scaffoldgen`

[![Go Report Card](https://goreportcard.com/badge/github.com/tidjee-dev/scaffoldgen)](https://goreportcard.com/report/github.com/tidjee-dev/scaffoldgen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> Lightweight **Go CLI** that generates safe, reviewable scaffold scripts from **Markdown, YAML, or JSON** project structures.

Instead of directly modifying your filesystem, scaffoldgen generates shell scripts (`scaffold.sh`) or PowerShell scripts (`scaffold.ps1`) that you review before execution.

## ✨ Quick Features

- 📝 Parse structures from **Markdown, YAML, or JSON**
- 🔄 Cross-platform script generation (Linux/macOS/Windows)
- 👀 Preview system with filesystem detection (`✓` / `+` / `…`)
- 🤖 Language-aware file templates (Go, Python, TypeScript, Rust, Java, C/C++, etc.)
- 🚫 Ignore rules engine for excluding generated/binary files
- 🔄 Reverse mode: scan directories → generate structure files
- 🔒 Safe by default: existing files are never overwritten

## 📦 Installation

### **Option 1: Go Install (Recommended)**

```bash
go install github.com/tidjee-dev/scaffoldgen/cmd/scaffoldgen@latest
```

**Auto-updates**: Scaffoldgen automatically checks for updates every 24 hours and will notify you when a new version is available.

**Manual update**: Run `scaffoldgen update` to update to the latest version.

### **Option 2: Download Binary (No Go required)**

Download from [GitHub Releases](https://github.com/tidjee-dev/scaffoldgen/releases/latest):

```bash
# macOS
curl -L -o scaffoldgen https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/scaffoldgen-darwin-arm64
chmod +x scaffoldgen
./scaffoldgen --version

# Linux
curl -L -o scaffoldgen https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/scaffoldgen-linux-amd64
chmod +x scaffoldgen
./scaffoldgen --version

# Windows
curl -L -o scaffoldgen.exe https://github.com/tidjee-dev/scaffoldgen/releases/latest/download/scaffoldgen-windows-amd64.exe
./scaffoldgen.exe --version
```

### **Option 3: Build from Source**

```bash
git clone https://github.com/tidjee-dev/scaffoldgen.git
cd scaffoldgen
make build
./scaffoldgen --version
```

## 🚀 Quick Start

```bash
# Generate scripts from a structure file
scaffoldgen generate --in structure.md --shell both

# Preview structure with filesystem detection
scaffoldgen preview --in structure.yml

# Scan an existing directory
scaffoldgen reverse --in ./src --format md

# Update to the latest version
scaffoldgen update
```

## 📚 Documentation

Full documentation is available at [scaffoldgen docs](https://tidjee-dev.github.io/scaffoldgen):

- [Getting Started](https://tidjee-dev.github.io/scaffoldgen/docs/getting-started/installation)
- [Format Guide](https://tidjee-dev.github.io/scaffoldgen/docs/guide/templating) (Markdown, YAML, JSON)
- [Usage Examples](https://tidjee-dev.github.io/scaffoldgen/docs/examples/examples)
- [Templating System](https://tidjee-dev.github.io/scaffoldgen/docs/guide/templating)

## 🔒 Safety First

The review-first workflow ensures you control what gets created:

```bash
scaffoldgen generate --in structure.md --shell sh
cat scaffold.sh        # 👀 Review the script
./scaffold.sh          # ✅ Execute when ready
```

Scripts include safety guards—existing files are never overwritten.

## 📄 License

MIT
