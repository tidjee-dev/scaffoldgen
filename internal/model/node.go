package model

type NodeType int

const (
	TypeDir NodeType = iota
	TypeFile
)

func (t NodeType) String() string {
	switch t {
	case TypeDir:
		return "DIR"
	case TypeFile:
		return "FILE"
	default:
		return "UNKNOWN"
	}
}

type Node struct {
	Name     string
	Type     NodeType
	Children []*Node

	Ignore   bool
	Template string // Template directive: empty (default), "none" (no template), or language name
}

func NewDir(name string) *Node {
	return &Node{Name: name, Type: TypeDir}
}

func NewFile(name string) *Node {
	return &Node{Name: name, Type: TypeFile}
}

func (n *Node) AddChild(child *Node) {
	if n == nil {
		panic("model.Node.AddChild called on nil receiver")
	}
	if child == nil {
		return
	}

	for _, c := range n.Children {
		if c.Name == child.Name && c.Type == child.Type {
			return
		}
	}

	n.Children = append(n.Children, child)
}

func (n *Node) IsDir() bool  { return n.Type == TypeDir }
func (n *Node) IsFile() bool { return n.Type == TypeFile }

func (n *Node) HasChildren() bool {
	return len(n.Children) > 0
}
