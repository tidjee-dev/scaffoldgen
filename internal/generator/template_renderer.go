package generator

import (
	"regexp"
	"strings"
	"time"
)

// RenderTemplate uses the enhanced renderer if available, falls back to legacy provider
func RenderTemplate(extension string, filename, pkgName, language, fullPath string) string {
	ctx := TemplateContext{
		Package:  pkgName,
		Filename: filename,
		Language: language,
		Path:     fullPath,
		Date:     time.Now(),
		Author:   "",
		Metadata: make(map[string]interface{}),
	}

	// Try modern renderer first
	if renderer := GetRenderer(extension); renderer != nil {
		return renderer(ctx)
	}

	// Fall back to legacy provider
	if provider := GetTemplate(extension); provider != nil {
		return provider(pkgName)
	}

	return ""
}

// RenderTemplateWithContext uses context directly for maximum control
func RenderTemplateWithContext(extension string, ctx TemplateContext) string {
	if renderer := GetRenderer(extension); renderer != nil {
		return renderer(ctx)
	}

	if provider := GetTemplate(extension); provider != nil {
		return provider(ctx.Package)
	}

	return ""
}

// ApplyVariableSubstitution replaces template variables in content
// Supported variables: {{package}}, {{filename}}, {{date}}, {{author}}, {{year}}
func ApplyVariableSubstitution(content string, ctx TemplateContext) string {
	replacer := strings.NewReplacer(
		"{{package}}", ctx.Package,
		"{{filename}}", ctx.Filename,
		"{{path}}", ctx.Path,
		"{{language}}", ctx.Language,
		"{{date}}", ctx.Date.Format("2006-01-02"),
		"{{datetime}}", ctx.Date.Format("2006-01-02 15:04:05"),
		"{{year}}", ctx.Date.Format("2006"),
		"{{author}}", ctx.Author,
	)
	return replacer.Replace(content)
}

// RenderTemplateWithSubstitution combines template rendering and variable substitution
func RenderTemplateWithSubstitution(extension string, ctx TemplateContext) string {
	content := RenderTemplateWithContext(extension, ctx)
	return ApplyVariableSubstitution(content, ctx)
}

// ValidateTemplate checks if template output is syntactically valid (basic checks)
func ValidateTemplate(extension string, content string) error {
	// Remove comments for validation
	cleaned := stripComments(content, extension)

	switch extension {
	case ".go":
		return validateGoTemplate(cleaned)
	case ".py":
		return validatePythonTemplate(cleaned)
	case ".java":
		return validateJavaTemplate(cleaned)
	case ".js", ".ts", ".tsx":
		return validateJavaScriptTemplate(cleaned)
	case ".h", ".hpp":
		return validateHeaderTemplate(cleaned)
	}

	return nil
}

func validateGoTemplate(content string) error {
	// Check for package declaration
	if !strings.Contains(content, "package") {
		return NewTemplateSyntaxError("go", "missing 'package' keyword")
	}
	return nil
}

func validatePythonTemplate(content string) error {
	// Check for balanced quotes in docstring
	if strings.Count(content, "\"\"\"")%2 != 0 {
		return NewTemplateSyntaxError("python", "unbalanced triple quotes")
	}
	return nil
}

func validateJavaTemplate(content string) error {
	// Check for class declaration or valid Java syntax
	if !(strings.Contains(content, "class ") || strings.Contains(content, "interface ")) {
		// Could be just package declaration, which is fine
		if !strings.Contains(content, "package") {
			return NewTemplateSyntaxError("java", "missing class/interface or package declaration")
		}
	}
	return nil
}

func validateJavaScriptTemplate(content string) error {
	// Basic JavaScript validation - check for balanced braces
	if strings.Count(content, "{") != strings.Count(content, "}") {
		return NewTemplateSyntaxError("javascript", "unbalanced braces")
	}
	return nil
}

func validateHeaderTemplate(content string) error {
	// Check for header guard
	if !strings.Contains(content, "#ifndef") || !strings.Contains(content, "#define") {
		return NewTemplateSyntaxError("header", "missing header guards (#ifndef/#define)")
	}
	return nil
}

func stripComments(content string, extension string) string {
	switch extension {
	case ".go", ".java", ".js", ".ts", ".cpp", ".c", ".cs":
		// Remove single-line comments
		re := regexp.MustCompile(`//.*`)
		content = re.ReplaceAllString(content, "")
		// Remove multi-line comments
		re = regexp.MustCompile(`/\*.*?\*/`)
		content = re.ReplaceAllString(content, "")
	case ".py":
		// Remove Python comments
		re := regexp.MustCompile(`#.*`)
		content = re.ReplaceAllString(content, "")
	}
	return content
}

// TemplateSyntaxError represents a template validation error
type TemplateSyntaxError struct {
	Language string
	Message  string
}

func (e *TemplateSyntaxError) Error() string {
	return "template syntax error (" + e.Language + "): " + e.Message
}

// NewTemplateSyntaxError creates a new template syntax error
func NewTemplateSyntaxError(language, message string) error {
	return &TemplateSyntaxError{
		Language: language,
		Message:  message,
	}
}
