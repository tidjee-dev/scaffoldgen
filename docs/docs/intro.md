---
sidebar_position: 0
title: Scaffoldgen
---

**Safe, reviewable scaffold generation from structured formats.**

Generate production-ready project structures from simple YAML, JSON, or Markdown files. Review the generated scripts before execution — because your filesystem should never be modified without your approval.

## ✨ Key Features

**📋 Multiple Input Formats** - Write structures in Markdown (fastest), YAML (readable), or JSON (programmatic)

**🔒 Safety First** - Generate shell scripts that you review and execute manually, never automatic modifications

**🚀 Cross-Platform** - Target Linux/macOS with Bash or Windows with PowerShell from one structure file

**🧠 Language-Aware** - Automatic appropriate boilerplate for 20+ languages (Go, Python, TypeScript, Rust, Java, etc.)

**👁️ Preview System** - Visualize exactly what will be created before generating scripts

**🔄 Reverse Mode** - Scan existing directories and generate structure files from them

## Why Scaffoldgen?

Traditional project generators modify your filesystem directly. Scaffoldgen generates **readable scripts** that you can:

- ✅ Review line-by-line
- ✅ Modify before running
- ✅ Version control
- ✅ Integrate with CI/CD
- ✅ Share safely with team members

## Get Started in Seconds

### 1️⃣ Create a Structure File

```yaml
# structure.yml
backend/:
  cmd/:
    api/:
      - main.go
  internal/:
    domain/:
      user/:
        - entity.go
        - repository.go
    infra/:
      db/:
        - migrate.go
```

### 2️⃣ Preview the Structure

```bash
scaffoldgen preview --in structure.yml
```

### 3️⃣ Generate the Scripts

```bash
scaffoldgen generate --in structure.yml --shell both
```

### 4️⃣ Review

```bash
cat scaffold.sh  # Review the script
```

### 5️⃣ Execute

```bash
./scaffold.sh  # Run the script
```

### 6️⃣ Other tools

- `validate` - Validate structure files

  ```bash
  scaffoldgen validate --in structure.yml
  ```

- `reverse` - Generate structure files from existing directories

  ```bash
  scaffoldgen reverse --in ./my-project --out structure.yml
  ```

## Living Documentation

**The structure file IS the documentation.** No separate diagrams or outdated READMEs.

- New team members understand the project layout immediately
- Changes to structure are tracked in git
- Templates enforce consistency across your codebase
- No drift between documentation and reality

## Perfect For

- **Creating new projects** - Bootstrap monorepos, microservices, full-stack apps
- **Onboarding templates** - Share project structures with your team
- **Codegen workflows** - Part of larger build pipelines
- **Documentation** - Keep architecture docs in sync with reality
- **CI/CD integration** - Generate scaffolds as part of automation

## Built With

- **Go 1.25+** - Fast, reliable, single binary
- **Minimal Dependencies** - Only `cobra` and `lipgloss`
- **Well-Tested** - Comprehensive test coverage
- **Open Source** - MIT licensed
