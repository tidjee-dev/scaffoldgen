package generator

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
	"gopkg.in/yaml.v3"
)

const indentSize = 2

// ExportMarkdown exports a node tree as Markdown
func ExportMarkdown(root *model.Node) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("# %s\n\n", root.Name))
	exportMarkdownNode(&b, root, 0)
	return b.String()
}

func exportMarkdownNode(b *strings.Builder, n *model.Node, depth int) {
	if n == nil || n.IsFile() {
		return
	}

	for _, child := range n.Children {
		indent := strings.Repeat(" ", depth*indentSize)
		name := child.Name
		if child.IsDir() {
			name += "/"
		}

		var meta string
		if child.Ignore {
			meta = " #ignore"
		}
		if child.Template != "" {
			meta += fmt.Sprintf(" #template:%s", child.Template)
		}

		fmt.Fprintf(b, "%s- %s%s\n", indent, name, meta)

		if child.IsDir() {
			exportMarkdownNode(b, child, depth+1)
		}
	}
}

// ExportJSON exports a node tree as JSON
func ExportJSON(root *model.Node) (string, error) {
	data := nodeToJSON(root)
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func nodeToJSON(n *model.Node) map[string]interface{} {
	result := make(map[string]interface{})

	for _, child := range n.Children {
		key := child.Name
		if child.IsDir() {
			key += "/"
		}

		if child.IsDir() && len(child.Children) > 0 {
			result[key] = nodeToJSON(child)
		} else if child.IsFile() {
			// Determine if we need metadata
			hasMetadata := child.Ignore || child.Template != ""

			if hasMetadata {
				// Create object with metadata
				obj := map[string]interface{}{
					"name": child.Name,
				}
				if child.Ignore {
					obj["ignore"] = true
				}
				if child.Template != "" {
					obj["template"] = child.Template
				}
				result[key] = obj
			} else {
				// Simple string entry
				result[key] = child.Name
			}
		}
	}

	return result
}

// ExportYAML exports a node tree as YAML
func ExportYAML(root *model.Node) (string, error) {
	data := nodeToYAML(root)
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func nodeToYAML(n *model.Node) map[string]interface{} {
	result := make(map[string]interface{})

	for _, child := range n.Children {
		key := child.Name
		if child.IsDir() {
			key += "/"
		}

		if child.IsDir() && len(child.Children) > 0 {
			result[key] = nodeToYAML(child)
		} else if child.IsFile() {
			// File entry
			name := child.Name
			if child.Ignore {
				name += " #ignore"
			}
			if child.Template != "" {
				name += fmt.Sprintf(" #template:%s", child.Template)
			}

			result[key] = []string{name}
		}
	}

	return result
}
