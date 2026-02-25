# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-02-26

### Added

- ✨ **Core Features**
  - `generate` command: Create Bash (Linux/macOS) and PowerShell (Windows) scaffold scripts
  - `preview` command: Visual tree preview with status indicators (✓/+/…)
  - `reverse` command: Scan existing directories and generate structure files
  - `validate` command: Check for filesystem conflicts before generation
  - `version` command: Display scaffoldgen version

- 📋 **Multiple Input Formats**
  - Markdown (`.md`) - Fast to write, lightweight parsing
  - YAML (`.yml`/`.yaml`) - Structured, readable format
  - JSON (`.json`) - Programmatic, metadata support

- 🌍 **Cross-Platform Support**
  - Linux (x86_64, ARM64)
  - macOS (Intel x86_64, Apple Silicon ARM64)
  - Windows (x86_64) with PowerShell scripts

- 🧠 **Language Templates** (20+ languages)
  - **Compiled:** Go, Java, C, C++, C#, Rust, Kotlin, Swift
  - **Scripting:** Python, JavaScript/TypeScript, Ruby, PHP, Lua
  - **Shell:** Bash, PowerShell

- 🛡️ **Safety Features**
  - Review-first workflow (scripts generated, never auto-executed)
  - Filesystem conflict detection
  - Safe guards in generated scripts (existing files/dirs untouched)
  - Deterministic, human-readable script output

- 📝 **Advanced Features**
  - Variable substitution in templates ({{package}}, {{date}}, {{author}}, {{year}})
  - Ignore rules for skipping files/directories
  - Template directives for per-file customization
  - Verbose logging for debugging
  - Dry-run mode for testing
  - Custom output directory support

- 📚 **Documentation**
  - Comprehensive README with examples
  - Docusaurus documentation site
  - Installation instructions (binary & source)
  - Quick-start guide
  - Templating system guide
  - Real-world examples in all three formats

### Technical Excellence

- ✅ Clean Go architecture (1.25+)
- ✅ Minimal dependencies (3: Cobra, lipgloss, YAML)
- ✅ Comprehensive error handling with context
- ✅ Unit tests for template system
- ✅ Cross-platform script generation with proper escaping
- ✅ Modular design (input → model → generator → export)

### Security

- ✅ No code injection vulnerabilities
- ✅ Safe shell escaping (Bash & PowerShell)
- ✅ Conflict detection prevents overwrites
- ✅ Review-first workflow prevents accidents
- ✅ No automation that bypasses safety checks

## [Unreleased]

### Planned for 1.1.0

- [ ] Integration tests for script execution
- [ ] Additional language template support (more variants)
- [ ] GitHub Action for CI/CD integration
- [ ] Package manager distributions (Homebrew, APT, Scoop)
- [ ] Web UI for structure builder

### Planned for 2.0.0

- [ ] Plugin system for custom templates
- [ ] Custom DSL for advanced structures
- [ ] Git integration (auto-commit generated files)
- [ ] Docker support (containerize scaffolded projects)
- [ ] Cloud storage integration

---

## Installation

### From Binary (Easiest)

Download pre-built binary from [GitHub Releases](https://github.com/tidjee-dev/scaffoldgen/releases)

### From Source

```bash
go install github.com/tidjee-dev/scaffoldgen/cmd/scaffoldgen@latest
```

Verify installation:

```bash
scaffoldgen version
```

---

## Quick Start

### 1. Create a Structure File

```yaml
# structure.yml
myapp/: src/
  main.go
  config.go
  tests/
  main_test.go
```

### 2. Preview

```bash
scaffoldgen preview --in structure.yml
```

### 3. Generate Scripts

```bash
scaffoldgen generate --in structure.yml --shell both --out scripts/
```

### 4. Review & Execute

```bash
cat scripts/scaffold.sh  # Review before running!
./scripts/scaffold.sh    # Then execute
```

---

## Known Limitations

- [ ] No built-in git integration (use scripts with git commands)
- [ ] No validation of Go module/package structure
- [ ] Template validation is basic (not comprehensive)
- [ ] No support for circular dependencies (by design - folders are acyclic)

---

## Contributing

Contributions welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) (if applicable)

---

## License

This project is licensed under the [License Type TBD] - see LICENSE file for details

---

## Credits & Acknowledgments

- Built with [Cobra](https://cobra.dev/) for CLI framework
- Styled with [lipgloss](https://github.com/charmbracelet/lipgloss) for terminal output
- Documented with [Docusaurus](https://docusaurus.io/)
- Go ecosystem for excellent standard library

---

## Support

- 📖 [Documentation](https://tidjee-dev.github.io/scaffoldgen/)
- 🐛 [Issue Tracker](https://github.com/tidjee-dev/scaffoldgen/issues)
- 💬 [Discussions](https://github.com/tidjee-dev/scaffoldgen/discussions)
