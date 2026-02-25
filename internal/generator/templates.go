package generator

import (
	"fmt"
	"strings"
)

// TemplateProvider generates initial content for a file based on its extension (legacy)
// Deprecated: Use TemplateRenderer with TemplateContext instead
type TemplateProvider func(pkgName string) string

var templates = map[string]TemplateProvider{
	".go":    goTemplate,
	".py":    pythonTemplate,
	".ts":    typescriptTemplate,
	".js":    javascriptTemplate,
	".rs":    rustTemplate,
	".java":  javaTemplate,
	".cpp":   cppTemplate,
	".c":     cTemplate,
	".h":     cHeaderTemplate,
	".hpp":   cppHeaderTemplate,
	".rb":    rubyTemplate,
	".php":   phpTemplate,
	".kt":    kotlinTemplate,
	".swift": swiftTemplate,
	".cs":    csharpTemplate,
	".sh":    shellTemplate,
	".lua":   luaTemplate,
}

var renderers = map[string]TemplateRenderer{
	".go":    goRenderer,
	".py":    pythonRenderer,
	".ts":    typescriptRenderer,
	".tsx":   typescriptRenderer,
	".js":    javascriptRenderer,
	".jsx":   javascriptRenderer,
	".rs":    rustRenderer,
	".java":  javaRenderer,
	".cpp":   cppRenderer,
	".cc":    cppRenderer,
	".cxx":   cppRenderer,
	".c":     cRenderer,
	".h":     cHeaderRenderer,
	".hpp":   cppHeaderRenderer,
	".hxx":   cppHeaderRenderer,
	".rb":    rubyRenderer,
	".php":   phpRenderer,
	".kt":    kotlinRenderer,
	".kts":   kotlinRenderer,
	".swift": swiftRenderer,
	".cs":    csharpRenderer,
	".sh":    shellRenderer,
	".lua":   luaRenderer,
}

// ================== Legacy Template Functions ==================

func goTemplate(pkgName string) string {
	return "package " + pkgName + "\n"
}

func pythonTemplate(pkgName string) string {
	return "\"\"\"" + pkgName + " module.\"\"\"\n"
}

func typescriptTemplate(pkgName string) string {
	return "// " + pkgName + " module\n"
}

func javascriptTemplate(pkgName string) string {
	return "// " + pkgName + " module\n"
}

func rustTemplate(pkgName string) string {
	return "// " + pkgName + " module\n"
}

func javaTemplate(pkgName string) string {
	className := toPascalCase(pkgName)
	return "public class " + className + " {\n}\n"
}

func cppTemplate(pkgName string) string {
	return "// " + pkgName + "\n#include <iostream>\n"
}

func cTemplate(pkgName string) string {
	return "// " + pkgName + "\n#include <stdio.h>\n"
}

func cHeaderTemplate(pkgName string) string {
	guard := strings.ToUpper(strings.ReplaceAll(pkgName, ".", "_")) + "_H"
	return "#ifndef " + guard + "\n#define " + guard + "\n\n#endif\n"
}

func cppHeaderTemplate(pkgName string) string {
	guard := strings.ToUpper(strings.ReplaceAll(pkgName, ".", "_")) + "_HPP"
	return "#ifndef " + guard + "\n#define " + guard + "\n\n#endif\n"
}

func rubyTemplate(pkgName string) string {
	return "# frozen_string_literal: true\n\n# " + pkgName + " module\n"
}

func phpTemplate(pkgName string) string {
	return "<?php\n\n/**\n * " + pkgName + " module\n */\n"
}

func kotlinTemplate(pkgName string) string {
	return "// " + pkgName + " module\n"
}

func swiftTemplate(pkgName string) string {
	return "// " + pkgName + " module\n"
}

func csharpTemplate(pkgName string) string {
	namespace := toPascalCase(pkgName)
	return "namespace " + namespace + "\n{\n}\n"
}

func shellTemplate(pkgName string) string {
	return "#!/usr/bin/env bash\n\n# " + pkgName + " script\n"
}

func luaTemplate(pkgName string) string {
	return "-- " + pkgName + " module\n"
}

// ================== Enhanced Renderer Functions ==================

func goRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("package %s\n", ctx.Package)
}

func pythonRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("\"\"\"Initialize %s module.\"\"\"\n", ctx.Package)
}

func typescriptRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("/**\n * %s module\n */\n\n", ctx.Package)
}

func javascriptRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("/**\n * %s module\n */\n\n", ctx.Package)
}

func rustRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("//! %s module\n\n", ctx.Package)
}

func javaRenderer(ctx TemplateContext) string {
	_, className, _, _ := ParseNameFormatting(ctx.Filename, "java")
	return fmt.Sprintf("/**\n * %s class\n */\npublic class %s {\n\n}\n", className, className)
}

func cppRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("// %s\n#include <iostream>\n\nnamespace %s {\n\n}\n", ctx.Package, strings.ToLower(ctx.Package))
}

func cRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("// %s\n#include <stdio.h>\n\n", ctx.Package)
}

func cHeaderRenderer(ctx TemplateContext) string {
	guard := strings.ToUpper(strings.ReplaceAll(ctx.Package, ".", "_")) + "_H"
	return fmt.Sprintf("#ifndef %s\n#define %s\n\n#endif // %s\n", guard, guard, guard)
}

func cppHeaderRenderer(ctx TemplateContext) string {
	guard := strings.ToUpper(strings.ReplaceAll(ctx.Package, ".", "_")) + "_HPP"
	return fmt.Sprintf("#ifndef %s\n#define %s\n\nnamespace %s {\n\n}\n\n#endif // %s\n", guard, guard, strings.ToLower(ctx.Package), guard)
}

func rubyRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("# frozen_string_literal: true\n\n# %s module\n\nmodule %s\nend\n", ctx.Package, toPascalCase(ctx.Package))
}

func phpRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("<?php\n\n/**\n * %s module\n */\n\nnamespace %s;\n", ctx.Package, toPascalCase(ctx.Package))
}

func kotlinRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("// %s module\n\npackage %s\n", ctx.Package, strings.ToLower(ctx.Package))
}

func swiftRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("//\n//  %s.swift\n//\n\nimport Foundation\n", ctx.Filename)
}

func csharpRenderer(ctx TemplateContext) string {
	namespace := toPascalCase(ctx.Package)
	_, className, _, _ := ParseNameFormatting(ctx.Filename, "csharp")
	return fmt.Sprintf("namespace %s\n{\n    public class %s\n    {\n    }\n}\n", namespace, className)
}

func shellRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("#!/usr/bin/env bash\n\n# %s script\n\nset -e\n\n", ctx.Package)
}

func luaRenderer(ctx TemplateContext) string {
	return fmt.Sprintf("-- %s module\n\nlocal %s = {}\n\nreturn %s\n", ctx.Package, strings.ToLower(ctx.Package), strings.ToLower(ctx.Package))
}

// ================== Template Lookup Functions ==================

// GetTemplate returns the legacy template provider for a file extension
func GetTemplate(ext string) TemplateProvider {
	if provider, ok := templates[strings.ToLower(ext)]; ok {
		return provider
	}
	return nil
}

// GetRenderer returns the enhanced template renderer for a file extension
func GetRenderer(ext string) TemplateRenderer {
	if renderer, ok := renderers[strings.ToLower(ext)]; ok {
		return renderer
	}
	return nil
}

// HasTemplate returns true if a template exists for the file extension
func HasTemplate(ext string) bool {
	_, ok := templates[strings.ToLower(ext)]
	return ok
}

// HasRenderer returns true if a renderer exists for the file extension
func HasRenderer(ext string) bool {
	_, ok := renderers[strings.ToLower(ext)]
	return ok
}
