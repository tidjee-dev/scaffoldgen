# scaffoldgen

> Lightweight **Go CLI** that generates safe, reviewable scaffold scripts from **Markdown, YAML, or JSON** project structures.

**Input:** `structure.md | structure.yml | structure.json`
**Output:** `scaffold.sh` / `scaffold.ps1`

`scaffoldgen` follows a **review-first workflow**: instead of modifying your filesystem directly, it generates scripts you inspect before execution.

## ✨ Features

### Core

- Parse project structures from:
  - Minimal Markdown lists
  - YAML
  - JSON

- Cross-platform script generation:
  - `scaffold.sh` (Linux/macOS)
  - `scaffold.ps1` (Windows)

- Deterministic, readable scripts
- Existing files/directories safely skipped
- Scripts are **never executed automatically**
- Lightweight parsing (no full Markdown engine)

### Preview System

- `preview` command with tree rendering
- Filesystem detection:
  - `✓` already exists
  - `+` will be created
  - `…` will be ignored

- Styled terminal output via **lipgloss**

Example:

```plain
─── Preview ───

✓ backend
├── ✓ domain
│   ├── ✓ calculator
│   │   ├── ✓ engine.go
│   │   └── ✓ result.go
└── + infra
    └── + db
        ├── … sqlcgenerated (ignored)
        ├── + migrate.go
        └── + sqlite.go
```

## 🔒 Safety First

Always review generated scripts before running them.

- Scripts modify your filesystem.
- Never execute scripts from untrusted structure files.

Recommended workflow:

```bash
scaffoldgen generate --in structure.md --shell sh
cat scaffold.sh
```

## 🧰 Stack

- Go 1.25+
- cobra
- lipgloss

## 📦 Installation

```bash
go install github.com/tidjee-dev/scaffoldgen/cmd/scaffoldgen@latest
```

## 🚀 Usage

### Generate Scripts

```bash
scaffoldgen generate \
  --in <structure.{md|yml|json}> \
  --shell <sh|ps1|both> \
  [--out <dir>] \
  [--dry-run] \
  [--verbose]
```

### Preview Structure

```bash
scaffoldgen preview --in structure.yml
```

### Reverse Mode: Scan Directory

Generate structure files from existing directories:

```bash
scaffoldgen reverse --in ./src --format md --out structure.md
scaffoldgen reverse --in ./backend --format json > structure.json
scaffoldgen reverse --in ./lib --format yml
```

## ⚙️ Options

### Generate & Preview

| Flag        | Description                              |
| ----------- | ---------------------------------------- |
| `--in`      | Path to structure file                   |
| `--in`      | Path to structure file (preview command) |
| `--shell`   | `sh`, `ps1`, or `both`                   |
| `--out`     | Output directory                         |
| `--dry-run` | Print scripts without writing files      |
| `--verbose` | Generated scripts log runtime actions    |

### Reverse Mode

| Flag       | Description                        |
| ---------- | ---------------------------------- |
| `--in`     | Path to directory to scan          |
| `--format` | Output format: `md`, `json`, `yml` |
| `--out`    | Output file (default: stdout)      |

## 🧩 Supported Input Formats

All formats are converted into the same internal AST.

[View examples](./examples.md)

### Markdown Format (`.md`)

#### Supported Syntax

- Only **one H1**
- Nested unordered lists (`-`)
- Two spaces per indentation level
- Directories end with `/`
- Files do not end with `/`
- Inline comments allowed (`#ignore`)

Unsupported:

- Multiple headings
- Ordered lists
- Tables or full Markdown syntax

Example:

```md
# backend

- domain/
  - calculator/
    - engine.go
    - result.go
```

### YAML Format (`.yml` / `.yaml`)

- Nested keys represent directories
- Arrays represent files
- Directory names end with `/`
- To ignore a directory or file, use inline comment `#ignore`

```yaml
backend/:
  domain/:
    calculator/:
      - engine.go
      - result.go #ignore
```

### JSON Format (`.json`)

JSON has no comments.
Metadata like `ignore` must use object entries.

```json
{
  "backend/": {
    "infra/": {
      "db/": ["sqlite.go", { "name": "sqlcgenerated/", "ignore": true }]
    }
  }
}
```

Full verbose JSON trees are optional but not required.

## 🧠 Behaviour

### Existing Files and Directories

Scripts include safety guards.

**Shell:**

```sh
[ -d "backend/domain" ] || mkdir -p "backend/domain"
[ -f "backend/domain/engine.go" ] || touch "backend/domain/engine.go"
```

**PowerShell:**

```ps1
if (!(Test-Path "backend/domain")) {
  New-Item -ItemType Directory -Path "backend/domain" | Out-Null
}
```

### Language-Aware Templates

Scaffoldgen automatically generates appropriate file templates based on file extension:

- `.go` → `package <dir_name>` (Go)
- `.py` → `"""<dir_name> module."""` (Python)
- `.ts` → `// <dir_name> module` (TypeScript)
- `.js` → `// <dir_name> module` (JavaScript)
- `.rs` → `// <dir_name> module` (Rust)
- `.java` → `public class ClassName { }` (Java)
- `.cpp` → C++ with includes
- `.c` → C with includes
- `.h` / `.hpp` → Header guards

Templates are **only applied to new files**. Existing files are never modified.

### Template Directives

Override automatic templating with metadata:

**JSON:**

```json
{
  "src/": [
    { "name": "config.py", "template": "python" },
    { "name": "script.sh", "template": "none" }
  ]
}
```

**Markdown & YAML:**

```markdown
- config.py #template:python
- script.sh #template:none
```

Supported directives:

- `template:none` - No template (just create empty file)
- `template:<lang>` - Force specific language template

### Go File Template

When creating `.go` files:

```go
package <dir_name>
```

Only applied if the file does not exist.

### Ignore Rules Engine

Mark files/directories to skip during generation:

**JSON:**

```json
{
  "src/": [{ "name": "sqlcgenerated/", "ignore": true }, "main.rs"]
}
```

**Markdown & YAML:**

```markdown
- sqlcgenerated/ #ignore
- main.rs
```

Ignored items:

- Show as `…` in preview
- Skipped during script generation
- Useful for excluding generated/binary files

### Verbose Mode

`--verbose` affects only generated scripts:

- Logs created files/directories at runtime
- CLI output remains minimal

## 🪟 Generated Script Types

### Shell (`.sh`)

- `mkdir -p`
- `touch` / `printf`
- POSIX-safe guards

### PowerShell (`.ps1`)

- `New-Item`
- `Test-Path`
- Wide version compatibility

## 🏗️ Project Structure

```plain
internal/
├── model/
├── input/
├── generator/
├── tui/
└── app/
```

### Architecture

- Parser builds a neutral AST
- Walker emits structured events
- Generators render scripts
- Preview reads AST directly
- CLI orchestrates commands

## 📋 Format Rules Summary

| Format   | Ignore Support        |
| -------- | --------------------- |
| Markdown | `#ignore` comment     |
| YAML     | `ignore: true`        |
| JSON     | Object entry required |

## 🔎 Why Generate Scripts Instead of Files?

- Review before execution
- Full filesystem control
- CI/CD friendly
- Safer for large repositories

## 🧪 Example Workflow

1. Create structure file:

   ```bash
   structure.yml
   ```

2. Preview:

   ```bash
   scaffoldgen preview --in structure.yml
   ```

3. Generate:

   ```bash
   scaffoldgen generate --in structure.yml --shell both
   ```

4. Review:

   ```bash
   cat scaffold.sh
   ```

   ```powershell
   cat .\scaffold.ps1
   ```

5. Execute manually:

   ```bash
   ./scaffold.sh
   ```

   ```powershell
   .\scaffold.ps1
   ```

## 📌 Roadmap & Status

### ✅ Completed

- ✅ Language-aware file templates (Go, Python, TypeScript, Rust, Java, C/C++, etc.)
- ✅ Template directives for explicit control
- ✅ Ignore rules engine for excluding files/directories
- ✅ Reverse mode (scan directory → generate structure file)

### 🎯 Future Enhancements

- Advanced glob pattern support for ignore rules
- Custom file template providers
- .gitignore integration
- Parallel generation
- Configuration file support (.scaffoldgenrc)
- More language templates

## 📄 License

MIT
