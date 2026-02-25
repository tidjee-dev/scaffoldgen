package generator

import (
	"strings"
	"time"
)

// TemplateContext provides rich context for template rendering
type TemplateContext struct {
	// Package/directory name (used for boilerplate naming)
	Package string

	// Filename including extension
	Filename string

	// Language identifier (go, python, typescript, etc.)
	Language string

	// Full relative path from root
	Path string

	// Current date/time for copyright headers
	Date time.Time

	// Optional author name for headers
	Author string

	// Custom metadata from structure file
	Metadata map[string]interface{}
}

// TemplateRenderer generates file content from template and context
type TemplateRenderer func(ctx TemplateContext) string

// NewTemplateContext creates a template context from file information
func NewTemplateContext(filename, pkgName, language, fullPath string) TemplateContext {
	return TemplateContext{
		Package:  pkgName,
		Filename: filename,
		Language: language,
		Path:     fullPath,
		Date:     time.Now(),
		Author:   "",
		Metadata: make(map[string]interface{}),
	}
}

// GetFileExtension extracts the extension from a filename
func GetFileExtension(filename string) string {
	idx := strings.LastIndexByte(filename, '.')
	if idx == -1 {
		return ""
	}
	return filename[idx:]
}

// GetLanguageFromExtension determines language from file extension
func GetLanguageFromExtension(filename string) string {
	ext := strings.ToLower(GetFileExtension(filename))
	switch ext {
	case ".go":
		return "go"
	case ".py":
		return "python"
	case ".ts", ".tsx":
		return "typescript"
	case ".js", ".jsx":
		return "javascript"
	case ".rs":
		return "rust"
	case ".java":
		return "java"
	case ".cpp", ".cc", ".cxx":
		return "cpp"
	case ".c":
		return "c"
	case ".h":
		return "c-header"
	case ".hpp", ".hxx":
		return "cpp-header"
	case ".rb":
		return "ruby"
	case ".php":
		return "php"
	case ".kt", ".kts":
		return "kotlin"
	case ".swift":
		return "swift"
	case ".cs":
		return "csharp"
	case ".sh":
		return "shell"
	case ".ps1":
		return "powershell"
	case ".lua":
		return "lua"
	case ".go.mod":
		return "go-mod"
	default:
		return "text"
	}
}

// FormatFilesize returns a nicely formatted filename (converts snake_case to camelCase for classes if needed)
func FormatFilename(name string, language string) string {
	base := strings.TrimRight(name, strings.Join([]string{".", GetFileExtension(name)}, ""))

	switch language {
	case "java", "csharp":
		// Convert snake_case to PascalCase for class-based languages
		return toPascalCase(base)
	default:
		return base
	}
}

// toPascalCase converts snake_case to PascalCase
func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	var result []string
	for _, part := range parts {
		if len(part) > 0 {
			result = append(result, strings.ToUpper(string(part[0]))+strings.ToLower(part[1:]))
		}
	}
	return strings.Join(result, "")
}

// ToCamelCase converts snake_case to camelCase
func ToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	if len(parts) == 0 {
		return s
	}

	var result []string
	for i, part := range parts {
		if i == 0 {
			result = append(result, strings.ToLower(part))
		} else if len(part) > 0 {
			result = append(result, strings.ToUpper(string(part[0]))+strings.ToLower(part[1:]))
		}
	}
	return strings.Join(result, "")
}

// ToSnakeCase converts camelCase or PascalCase to snake_case
func ToSnakeCase(s string) string {
	var result []byte
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, byte(strings.ToLower(string(r))[0]))
	}
	return string(result)
}

// ParseNameFormatting extracts base name and determines formatting
func ParseNameFormatting(name, language string) (base, className, funcName, varName string) {
	base = strings.TrimSuffix(name, GetFileExtension(name))

	// Handle different naming conventions by language
	switch language {
	case "java", "csharp":
		className = FormatFilename(name, language)
		funcName = ToCamelCase(base)
		varName = ToCamelCase(base)
	default:
		className = toPascalCase(base)
		funcName = ToCamelCase(base)
		varName = ToSnakeCase(base)
	}

	return base, className, funcName, varName
}
