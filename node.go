package libgenders

type Node struct {
	Name       string
	Attributes map[string]string
}

func NewNode(name string, attributes map[string]string) Node {
	return Node{
		Name:       name,
		Attributes: attributes,
	}
}

func (n Node) MergeAttributes(attributes map[string]string) {
	for key, value := range attributes {
		n.Attributes[key] = value
	}
}
