package input

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

func ParseJSON(r io.Reader) (*model.Node, error) {
	var data map[string]interface{}
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	if len(data) != 1 {
		return nil, fmt.Errorf("JSON root must contain exactly one key")
	}

	var rootName string
	var rootData interface{}
	for k, v := range data {
		rootName = strings.TrimSuffix(k, "/")
		rootData = v
	}

	root := model.NewDir(rootName)
	if err := parseJSONNode(root, rootData); err != nil {
		return nil, err
	}

	return root, nil
}

func parseJSONNode(parent *model.Node, data interface{}) error {
	switch v := data.(type) {
	case map[string]interface{}:
		return parseJSONMap(parent, v)
	case []interface{}:
		return parseJSONArray(parent, v)
	default:
		return fmt.Errorf("unexpected type: %T", v)
	}
}

func parseJSONMap(parent *model.Node, m map[string]interface{}) error {
	for key, value := range m {
		name := strings.TrimSuffix(key, "/")
		isDir := strings.HasSuffix(key, "/")

		var node *model.Node
		if isDir {
			node = model.NewDir(name)
		} else {
			node = model.NewFile(name)
		}

		parent.AddChild(node)

		if isDir && value != nil {
			if err := parseJSONNode(node, value); err != nil {
				return err
			}
		}
	}
	return nil
}

func parseJSONArray(parent *model.Node, arr []interface{}) error {
	for _, item := range arr {
		switch v := item.(type) {
		case string:
			// Simple file entry
			node := model.NewFile(v)
			parent.AddChild(node)

		case map[string]interface{}:
			// Entry with metadata like {"name": "dir/", "ignore": true, "template": "python"}
			name, ok := v["name"].(string)
			if !ok {
				return fmt.Errorf("array item must have 'name' property")
			}

			ignore, _ := v["ignore"].(bool)
			template, _ := v["template"].(string)
			isDir := strings.HasSuffix(name, "/")
			name = strings.TrimSuffix(name, "/")

			var node *model.Node
			if isDir {
				node = model.NewDir(name)
			} else {
				node = model.NewFile(name)
			}
			node.Ignore = ignore
			node.Template = template

			parent.AddChild(node)

			// If it's a directory with nested structure
			if isDir {
				for k, val := range v {
					if k != "name" && k != "ignore" && k != "template" {
						if err := parseJSONNode(node, val); err != nil {
							return err
						}
					}
				}
			}

		default:
			return fmt.Errorf("array item must be a string or object, got %T", v)
		}
	}
	return nil
}
