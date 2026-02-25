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

```bash
go install github.com/tidjee-dev/scaffoldgen/cmd/scaffoldgen@latest
```

## 🚀 Quick Start

```bash
# Generate scripts from a structure file
scaffoldgen generate --in structure.md --shell both

# Preview structure with filesystem detection
scaffoldgen preview --in structure.yml

# Scan an existing directory
scaffoldgen reverse --in ./src --format md
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
