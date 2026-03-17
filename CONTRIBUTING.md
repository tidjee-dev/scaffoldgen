# Contributing to ScaffoldGen

Thank you for your interest in contributing to ScaffoldGen! This document provides guidelines for contributors.

## Getting Started

### Prerequisites

- Go 1.25.6 or later
- Git
- Make (optional, for using Makefile commands)

### Development Setup

1. **Fork and clone the repository**

   ```bash
   git clone https://github.com/your-username/scaffoldgen.git
   cd scaffoldgen
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Build the project**

   ```bash
   make build
   # or directly
   go build -o scaffoldgen ./cmd/scaffoldgen
   ```

4. **Run tests**
   ```bash
   make test
   # or directly
   go test ./...
   ```

## Development Workflow

### Making Changes

1. **Create a feature branch**

   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes**
   - Follow existing code style and patterns
   - Add tests for new functionality
   - Update documentation if needed

3. **Test your changes**

   ```bash
   make test
   make lint  # if available
   ```

4. **Commit your changes**

   ```bash
   git add .
   git commit -m "feat: add your feature description"
   ```

5. **Push and create a pull request**
   ```bash
   git push origin feature/your-feature-name
   ```

## Code Style

### Go Code

- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for exported functions and complex logic
- Keep functions small and focused

### Commit Messages

Use conventional commit format:

- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `refactor:` for code refactoring
- `test:` for test-related changes
- `chore:` for maintenance tasks

Example:

```
feat: add support for TOML input format
fix: resolve template variable substitution edge case
docs: update installation instructions
```

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific package tests
go test ./internal/generator
```

### Writing Tests

- Write unit tests for new functionality
- Test edge cases and error conditions
- Use table-driven tests for multiple scenarios
- Aim for good test coverage

## Project Structure

```
scaffoldgen/
├── cmd/scaffoldgen/     # CLI entry point
├── internal/
│   ├── app/            # Application logic
│   ├── generator/      # Code generation
│   ├── input/          # Input format parsers
│   └── model/          # Data models
├── docs/               # Docusaurus documentation
├── scripts/            # Build and utility scripts
├── examples/           # Example structures
├── .github/            # GitHub Actions workflows
└── dist/               # Build artifacts (gitignored)
```

## Areas for Contribution

### High Priority

- [ ] Additional language templates
- [ ] Integration tests for script execution
- [ ] Performance improvements
- [ ] Error handling enhancements
- [ ] GitHub Action for CI/CD integration

### Medium Priority

- [ ] Documentation improvements
- [ ] Code examples and tutorials
- [ ] Plugin system architecture
- [ ] Web UI prototype
- [ ] Package manager distributions (Homebrew, APT, Scoop)

### Low Priority

- [ ] Additional input formats
- [ ] Advanced template features
- [ ] Cloud integrations

## Submitting Changes

### Pull Request Process

1. **Update documentation** if your changes affect user-facing behavior
2. **Add tests** for new functionality
3. **Ensure all tests pass** before submitting
4. **Update CHANGELOG.md** for significant changes
5. **Create a descriptive PR title** and description

### PR Template

```markdown
## Description

Brief description of changes made.

## Type of Change

- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing

- [ ] Tests pass locally
- [ ] Added new tests for new functionality
- [ ] Manual testing completed

## Checklist

- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Documentation updated if needed
- [ ] CHANGELOG.md updated for significant changes
```

## Getting Help

- Create an issue for bugs or feature requests
- Start a discussion for questions or ideas
- Check existing issues and discussions first

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to ScaffoldGen! 🎉
