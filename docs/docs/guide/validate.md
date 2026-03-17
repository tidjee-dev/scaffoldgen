---
sidebar_position: 6
title: Validate Command
---

The `validate` command checks a structure file against the current filesystem to detect potential conflicts before generation.

## Usage

```bash
scaffoldgen validate --in <structure-file>
```

## Examples

```bash
# Validate a markdown structure file
scaffoldgen validate --in structure.md

# Validate a YAML structure file
scaffoldgen validate --in structure.yml

# Validate a JSON structure file
scaffoldgen validate --in structure.json
```

## What It Checks

The validate command scans your structure and compares it with existing files and directories:

### ✅ **No Conflicts**

- Files/directories don't exist yet
- Safe to proceed with generation

### ⚠️ **Warnings**

- Directory already exists (not a conflict, but noted)
- Files would be created in existing directories

### ❌ **Conflicts**

- **File exists as directory**: Structure expects a file but a directory exists
- **Directory exists as file**: Structure expects a directory but a file exists
- **File already exists**: Structure would overwrite an existing file

## Output Examples

### Successful Validation

```bash
$ scaffoldgen validate --in structure.md
✅ Validation passed - no conflicts detected
```

### Conflicts Detected

```bash
$ scaffoldgen validate --in structure.md
❌ Validation failed:
 - src/main.go file already exists
 - config exists as DIRECTORY but expected FILE
 - docs/README.md file already exists
```

## When to Use

Use `validate` before running `generate` to:

1. **Prevent accidental overwrites** - Check if you're about to overwrite existing files
2. **Verify structure** - Ensure your structure matches expectations
3. **CI/CD pipelines** - Fail fast if there are conflicts
4. **Team collaboration** - Check if someone else created files you're expecting to generate

## Workflow Integration

### Typical Safe Workflow

```bash
# 1. Preview the structure
scaffoldgen preview --in structure.md

# 2. Validate for conflicts
scaffoldgen validate --in structure.md

# 3. Generate scripts (only if validation passes)
scaffoldgen generate --in structure.md --shell both

# 4. Review and execute
cat scaffold.sh
./scaffold.sh
```

### CI/CD Integration

```bash
#!/bin/bash
# ci-check.sh

echo "Validating project structure..."
if scaffoldgen validate --in structure.md; then
    echo "✅ Structure validation passed"
    exit 0
else
    echo "❌ Structure validation failed"
    exit 1
fi
```

## Exit Codes

- **0**: Success - no conflicts detected
- **1**: Error - conflicts detected or file not found

## Flags

- `--in`: Path to structure file (required)
- `--help`: Show help information

Supported input formats: `.md`, `.yml`, `.yaml`, `.json`

## Tips

- Run `validate` before every `generate` to ensure safety
- Use in automated scripts to prevent accidental overwrites
- Combine with `preview` for complete visibility
- The command only checks for conflicts, it doesn't modify anything
