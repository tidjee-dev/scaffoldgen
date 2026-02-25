package input

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/tidjee-dev/scaffoldgen/internal/model"
)

const indentSize = 2

func ParseMarkdown(r io.Reader) (*model.Node, error) {
	scanner := bufio.NewScanner(r)

	var root *model.Node
	stack := map[int]*model.Node{}
	lineNo := 0

	for scanner.Scan() {
		lineNo++
		raw := scanner.Text()

		if strings.Contains(raw, "\t") {
			return nil, fmt.Errorf("line %d: tabs are not allowed", lineNo)
		}

		line := strings.TrimRight(raw, " ")
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.HasPrefix(line, "# ") {
			if root != nil {
				return nil, fmt.Errorf("line %d: multiple root titles are not allowed", lineNo)
			}
			name := strings.TrimSuffix(strings.TrimSpace(strings.TrimPrefix(line, "# ")), "/")
			root = model.NewDir(name)
			stack[0] = root
			continue
		}

		if root == nil {
			return nil, fmt.Errorf("line %d: missing root title (# root)", lineNo)
		}

		leadingSpaces := countLeadingSpaces(line)
		if leadingSpaces%indentSize != 0 {
			return nil, fmt.Errorf("line %d: invalid indentation", lineNo)
		}

		depth := leadingSpaces/indentSize + 1

		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "- ") {
			return nil, fmt.Errorf("line %d: expected '- ' list item", lineNo)
		}

		name := strings.TrimPrefix(trimmed, "- ")
		name, meta := parseMeta(name)

		isDir := strings.HasSuffix(name, "/")
		name = strings.TrimSuffix(name, "/")

		var node *model.Node
		if isDir {
			node = model.NewDir(name)
		} else {
			node = model.NewFile(name)
		}
		node.Ignore = meta.Ignore
		node.Template = meta.Template

		parent, ok := stack[depth-1]
		if !ok {
			return nil, fmt.Errorf("line %d: invalid hierarchy depth", lineNo)
		}

		for _, c := range parent.Children {
			if c.Name == node.Name {
				return nil, fmt.Errorf("line %d: duplicate entry '%s'", lineNo, node.Name)
			}
		}

		parent.AddChild(node)
		stack[depth] = node
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if root == nil {
		return nil, fmt.Errorf("no root title found")
	}
	return root, nil
}

func countLeadingSpaces(s string) int {
	count := 0
	for _, r := range s {
		if r == ' ' {
			count++
			continue
		}
		break
	}
	return count
}

func parseMeta(s string) (string, Meta) {
	idx := strings.Index(s, "#")
	if idx == -1 {
		return strings.TrimSpace(s), Meta{}
	}

	name := strings.TrimSpace(s[:idx])
	metaStr := strings.ToLower(s[idx+1:])

	meta := Meta{
		Ignore: strings.Contains(metaStr, "ignore"),
	}

	// Extract template directive (e.g., #template:python or #py)
	if strings.Contains(metaStr, "template:") {
		parts := strings.Split(metaStr, "template:")
		if len(parts) > 1 {
			template := strings.TrimSpace(parts[1])
			// Extract only the first word
			if idx := strings.IndexAny(template, " \t"); idx != -1 {
				template = template[:idx]
			}
			meta.Template = template
		}
	}

	return name, meta
}

type Meta struct {
	Ignore   bool
	Template string
}
