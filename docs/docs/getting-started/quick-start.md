---
title: Quick Start
sidebar_position: 2
---


Get up and running with scaffoldgen in 5 minutes.

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

| Command | Purpose |
| ------- | ------- |
| `preview` | See what will be created |
| `generate` | Create the scaffold scripts |
| `reverse` | Scan a directory and generate a structure file |
| `validate` | Check structure against existing directory |
| `version` | Show version info |
