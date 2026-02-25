---
title: Templating System
sidebar_position: 1
---

## Overview

The templating system has been significantly improved to provide:

- **Rich Template Context** - More information available to templates
- **Multiple Language Support** - 20+ programming languages
- **Variable Substitution** - Use `{{variable}}` placeholders in files
- **Better Boilerplate** - More realistic and complete templates
- **Backward Compatibility** - Old system continues to work
- **Template Validation** - Basic syntax checking for generated templates

## Key Improvements

### 1. Enhanced Template Context

Templates now receive rich context including:

```go
type TemplateContext struct {
    Package   string                    // Directory/module name
    Filename  string                    // Just the filename
    Language  string                    // Language identifier
    Path      string                    // Full relative path
    Date      time.Time                 // Current date/time
    Author    string                    // Optional author name
    Metadata  map[string]interface{}    // Custom metadata
}
```

### 2. Supported Languages

The system now includes templates for 20+ languages:

**Compiled Languages:**

- Go (`.go`) - Package declarations
- Java (`.java`) - Class structures with Javadoc
- C (`.c`, `.h`) - Includes with header guards
- C++ (`.cpp`, `.cc`, `.cxx`, `.hpp`, `.hxx`) - Namespaces and includes
- Rust (`.rs`) - Module documentation
- Kotlin (`.kt`, `.kts`) - Package and imports
- Swift (`.swift`) - Import statements
- C# (`.cs`) - Namespace structures

**Interpreted Languages:**

- Python (`.py`) - Module docstrings
- JavaScript (`.js`, `.jsx`) - JSDoc comments
- TypeScript (`.ts`, `.tsx`) - TSDoc comments and types
- Ruby (`.rb`) - Frozen string literals and modules
- PHP (`.php`) - Opening PHP tags with namespace
- Lua (`.lua`) - Lua modules

**Scripting:**

- Shell/Bash (`.sh`) - Shebang with error handling
- PowerShell (`.ps1`) - PowerShell specific headers

### 3. Template Rendering

Two rendering strategies:

#### Legacy Renderer (Backward Compatible)

```go
template := generator.GetTemplate(".go")
content := template("mypackage")
```

#### Enhanced Renderer (New)

```go
ctx := generator.TemplateContext{
    Package:  "mypackage",
    Filename: "handler.go",
    Language: "go",
    Path:     "cmd/api/handler.go",
}
content := generator.RenderTemplateWithContext(".go", ctx)
```

### 4. Variable Substitution

Use variables in structure metadata to customize templates:

Supported variables:

- `{{package}}` - Package/module name
- `{{filename}}` - Filename with extension
- `{{path}}` - Full file path
- `{{language}}` - Language identifier
- `{{date}}` - Current date (YYYY-MM-DD)
- `{{datetime}}` - Full datetime (YYYY-MM-DD HH:MM:SS)
- `{{year}}` - Current year
- `{{author}}` - Author name (if set)

### 5. Template Validation

Basic syntax checking validates generated templates:

```go
if err := generator.ValidateTemplate(".go", content); err != nil {
    log.Printf("Template error: %v", err)
}
```

Checks for:

- Missing required keywords (package, class, etc.)
- Balanced braces/brackets
- Header guards in C/C++
- Proper quotes in Python

### 6. Better Named Templates

Helper functions for name formatting:

```go
// Convert snake_case to different formats
base := "user_handler"
camelCase := generator.ToCamelCase(base)      // "userHandler"
pascalCase := generator.FormatFilename(base, "java")   // "UserHandler"
snakeCase := generator.ToSnakeCase(base)      // "user_handler"
```

## Migration Guide

### Before (Legacy)

```go
template := generator.GetTemplate(".py")
content := template("models")  // Returns minimal docstring
```

### After (Enhanced, Fully Compatible)

```go
// Still works with legacy templates
template := generator.GetTemplate(".py")
content := template("models")

// Or use enhanced renderer
ctx := generator.NewTemplateContext("models.py", "models", "python", "app/models.py")
content := generator.RenderTemplateWithContext(".py", ctx)

// With variable substitution
content = generator.RenderTemplateWithSubstitution(".py", ctx)
```

## Usage Examples

### Creating A Go Project

```bash
scaffoldgen preview --in structure.yml
scaffoldgen generate --in structure.yml --shell sh
```

With enhanced templates, generated files will have:

```go
package myservice

// Additional context available for IDE integration
```

### Creating A Full-Stack Project

```json
{
  "myapp/": {
    "backend/": {
      "main.go": {},
      "handler.go": {}
    },
    "frontend/": {
      "App.tsx": {},
      "main.ts": {}
    }
  }
}
```

Generated files will include appropriate headers for each language.

## Architecture

### File Organization

```plain
internal/generator/
├── template_context.go      # Context and utilities
├── templates.go             # Template definitions (legacy + new)
├── template_renderer.go     # Rendering and substitution
├── walker.go                # AST walking (updated)
├── sh.go                    # Shell script generation (updated)
└── ps1.go                   # PowerShell generation (updated)
```

### Key Components

1. **TemplateContext** - Rich metadata about the file being created
2. **TemplateRenderer** - Function type for rendering templates
3. **GetRenderer()** - Lookup enhanced renderers by extension
4. **GetTemplate()** - Lookup legacy templates (backward compatible)
5. **RenderTemplate()** - Smart routing to right renderer

## Future Enhancements

Possible improvements:

1. **Custom template files** - Load templates from disk
2. **Template inheritance** - Base templates with overrides
3. **Conditional sections** - `{{#if condition}}...{{/if}}`
4. **Filters** - `{{package | upper}}`, `{{date | long}}`
5. **Inline template blocks** - Direct template content in structure files
6. **Plugin system** - Custom renderer registration
