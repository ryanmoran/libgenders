package libgenders

import (
	"fmt"
)

type QueryEngine struct {
	attrs    map[string]Set
	attrvals map[string]Set
	indices  Set
}

func NewQueryEngine(nodes []Node) QueryEngine {
	attrs := make(map[string]Set)
	attrvals := make(map[string]Set)
	indices := make(Set, len(nodes))

	for index, node := range nodes {
		indices[index] = index
		for key, value := range node.Attributes {
			attrs[key] = append(attrs[key], index)

			if value != "" {
				keyval := fmt.Sprintf("%s=%s", key, value)
				attrvals[keyval] = append(attrvals[keyval], index)
			}
		}
	}

	return QueryEngine{
		attrs:    attrs,
		attrvals: attrvals,
		indices:  indices,
	}
}

func (qe QueryEngine) Query(query string) Set {
	tokens, err := Tokenize(query)
	if err != nil {
		panic(err)
	}

	return ParseQuery(tokens).Evaluate(qe.attrs, qe.attrvals, qe.indices)
}
