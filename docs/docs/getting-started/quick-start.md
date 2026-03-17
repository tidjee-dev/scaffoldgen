---
title: Quick Start
sidebar_position: 2
---

Get up and running with scaffoldgen in 5 minutes.

## 📦 Installation

### **Option 1: Go Install (Recommended)**

The easiest way to install scaffoldgen is with `go install`:

```bash
go install github.com/tidjee-dev/scaffoldgen/cmd/scaffoldgen@latest
```

This installs the latest version and keeps it updated automatically.

### **Option 2: Download Binary**

If you don't have Go installed, download pre-compiled binaries from [GitHub Releases](https://github.com/tidjee-dev/scaffoldgen/releases/latest).

### **Option 3: Build from Source**

```bash
git clone https://github.com/tidjee-dev/scaffoldgen.git
cd scaffoldgen
make build
```

## 1. Create a Structure File

Create `structure.md` with your project structure:

```markdown
# myapp

- src/
  - main.go
  - config.go

- tests/
  - main_test.go

- pkg/
  - logger/
    - logger.go
```

## 2. Preview the Structure

See what will be created:

```bash
scaffoldgen preview --in structure.md
```

You'll see a tree view with indicators:

- `✓` - Already exists
- `+` - Will be created
- `…` - Will be ignored

## 3. Generate Scripts

Generate the scaffold scripts:

```bash
scaffoldgen generate --in structure.md --shell both --out scripts
```

This creates:

- `scripts/scaffold.sh` (Linux/macOS)
- `scripts/scaffold.ps1` (Windows)

## 4. Review & Execute

Review the generated script:

```bash
cat scripts/scaffold.sh
# or on Windows:
cat .\scaffold.ps1
```

Then execute it:

```bash
./scripts/scaffold.sh
# or on Windows:
.\scripts\scaffold.ps1
```

## Supported Input Formats

Scaffoldgen accepts three input formats - choose what works for you:

- **Markdown** (`.md`) - Fast to write manually
- **YAML** (`.yml`) - Structured and readable
- **JSON** (`.json`) - Programmatic and flexible

See [Examples](../examples/examples.md) for each format.

## Key Commands

| Command      | Purpose                                               |
| ------------ | ----------------------------------------------------- |
| `preview`    | See what will be created                              |
| `generate`   | Create the scaffold scripts                           |
| `reverse`    | Scan a directory and generate a structure file        |
| `validate`   | Check structure against existing directory            |
| `version`    | Show version info (use `--verbose` for build details) |
| `completion` | Generate shell autocompletion scripts                 |

## Additional Examples

### Generate Shell Completion

```bash
# For bash
scaffoldgen completion bash > /etc/bash_completion.d/scaffoldgen

# For zsh
scaffoldgen completion zsh > ~/.zsh/completions/_scaffoldgen

# For fish
scaffoldgen completion fish > ~/.config/fish/completions/scaffoldgen.fish
```

### Show Version Information

```bash
# Basic version
scaffoldgen version

# Detailed build information
scaffoldgen version --verbose
```

### Validate Before Generation

```bash
# Check for conflicts before generating
scaffoldgen validate --in structure.md

# This will show any existing files that would conflict
```
