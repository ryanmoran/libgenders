package libgenders

type Node struct {
	Name       string
	Attributes map[string]string
}

func (n Node) mergeAttributes(attributes map[string]string) {
	for key, value := range attributes {
		n.Attributes[key] = value
	}
}
