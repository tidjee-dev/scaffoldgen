package input

import (
	"fmt"
	"io"
	"strings"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
	"gopkg.in/yaml.v3"
)

func ParseYAML(r io.Reader) (*model.Node, error) {
	var doc yaml.Node
	if err := yaml.NewDecoder(r).Decode(&doc); err != nil {
		return nil, fmt.Errorf("invalid YAML: %w", err)
	}

	// doc is a DocumentNode, doc.Content[0] is the actual root
	if len(doc.Content) == 0 {
		return nil, fmt.Errorf("empty YAML document")
	}

	root := doc.Content[0]
	if root.Kind != yaml.MappingNode {
		return nil, fmt.Errorf("YAML root must be a mapping")
	}

	if len(root.Content) < 2 {
		return nil, fmt.Errorf("YAML root must contain exactly one key")
	}

	rootKeyNode := root.Content[0]
	rootValueNode := root.Content[1]

	rootName := strings.TrimSuffix(rootKeyNode.Value, "/")
	nodeRoot := model.NewDir(rootName)

	if err := parseYAMLNodeValue(nodeRoot, rootValueNode); err != nil {
		return nil, err
	}

	return nodeRoot, nil
}

func parseYAMLNodeValue(parent *model.Node, node *yaml.Node) error {
	if node == nil {
		return nil
	}

	switch node.Kind {
	case yaml.MappingNode:
		return parseYAMLMapping(parent, node)
	case yaml.SequenceNode:
		return parseYAMLSequence(parent, node)
	default:
		return fmt.Errorf("unexpected YAML node kind: %v", node.Kind)
	}
}

func parseYAMLMapping(parent *model.Node, node *yaml.Node) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("expected mapping node")
	}

	// Process key-value pairs
	for i := 0; i < len(node.Content); i += 2 {
		keyNode := node.Content[i]
		valueNode := node.Content[i+1]

		key := strings.TrimSpace(keyNode.Value)
		name := strings.TrimSuffix(key, "/")
		isDir := strings.HasSuffix(key, "/")

		var child *model.Node
		if isDir {
			child = model.NewDir(name)
		} else {
			child = model.NewFile(name)
		}

		parent.AddChild(child)

		if isDir && valueNode != nil {
			if err := parseYAMLNodeValue(child, valueNode); err != nil {
				return err
			}
		}
	}
	return nil
}

func parseYAMLSequence(parent *model.Node, node *yaml.Node) error {
	if node.Kind != yaml.SequenceNode {
		return fmt.Errorf("expected sequence node")
	}

	for _, itemNode := range node.Content {
		var name string
		var metadata YAMLMeta
		var isDir bool

		switch itemNode.Kind {
		case yaml.ScalarNode:
			// Simple string entry - extract comment if present
			name = strings.TrimSpace(itemNode.Value)
			metadata = parseYAMLMetadata(name)

			// Also check the node comment
			if itemNode.LineComment != "" {
				commentMeta := parseYAMLMetadata(itemNode.LineComment)
				metadata.Ignore = metadata.Ignore || commentMeta.Ignore
				if commentMeta.Template != "" {
					metadata.Template = commentMeta.Template
				}
			}

			isDir = strings.HasSuffix(metadata.Name, "/")
			name = strings.TrimSuffix(metadata.Name, "/")

			var child *model.Node
			if isDir {
				child = model.NewDir(name)
			} else {
				child = model.NewFile(name)
			}
			child.Ignore = metadata.Ignore
			child.Template = metadata.Template
			parent.AddChild(child)

		case yaml.MappingNode:
			// Inline mapping with metadata
			metadata = YAMLMeta{}

			for i := 0; i < len(itemNode.Content); i += 2 {
				keyNode := itemNode.Content[i]
				valNode := itemNode.Content[i+1]

				key := keyNode.Value
				switch key {
				case "name":
					metadata.Name = valNode.Value
				case "ignore":
					// Handle boolean ignore flag
					metadata.Ignore = strings.ToLower(valNode.Value) == "true"
				case "template":
					metadata.Template = valNode.Value
				}
			}

			if metadata.Name == "" {
				return fmt.Errorf("mapping item must have 'name' property")
			}

			parsedMeta := parseYAMLMetadata(metadata.Name)
			metadata.Ignore = metadata.Ignore || parsedMeta.Ignore
			if parsedMeta.Template != "" {
				metadata.Template = parsedMeta.Template
			}

			isDir := strings.HasSuffix(metadata.Name, "/")
			name := strings.TrimSuffix(metadata.Name, "/")

			var child *model.Node
			if isDir {
				child = model.NewDir(name)
			} else {
				child = model.NewFile(name)
			}
			child.Ignore = metadata.Ignore
			child.Template = metadata.Template
			parent.AddChild(child)

		default:
			return fmt.Errorf("sequence item must be scalar or mapping, got %v", itemNode.Kind)
		}
	}
	return nil
}

func parseYAMLMetadata(s string) YAMLMeta {
	idx := strings.Index(s, "#")
	if idx == -1 {
		return YAMLMeta{Name: strings.TrimSpace(s)}
	}

	name := strings.TrimSpace(s[:idx])
	metaStr := strings.ToLower(s[idx+1:])

	meta := YAMLMeta{
		Name:   name,
		Ignore: strings.Contains(metaStr, "ignore"),
	}

	// Extract template directive
	if strings.Contains(metaStr, "template:") {
		parts := strings.Split(metaStr, "template:")
		if len(parts) > 1 {
			template := strings.TrimSpace(parts[1])
			if idx := strings.IndexAny(template, " \t"); idx != -1 {
				template = template[:idx]
			}
			meta.Template = template
		}
	}

	return meta
}

type YAMLMeta struct {
	Name     string
	Ignore   bool
	Template string
}

func parseYAMLMeta(s string) (string, bool) {
	idx := strings.Index(s, "#")
	if idx == -1 {
		return strings.TrimSpace(s), false
	}

	name := strings.TrimSpace(s[:idx])
	meta := strings.ToLower(s[idx+1:])
	return name, strings.Contains(meta, "ignore")
}
