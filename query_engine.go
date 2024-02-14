package libgenders

import (
	"fmt"

	"github.com/ryanmoran/libgenders/search"
)

type QueryEngine struct {
	attrs    map[string]search.Set
	attrvals map[string]search.Set
	indices  search.Set
}

func NewQueryEngine(nodes []Node) QueryEngine {
	attrs := make(map[string]search.Set)
	attrvals := make(map[string]search.Set)
	indices := make(search.Set, len(nodes))

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

func (qe QueryEngine) Query(query string) []int {
	tokens, err := search.Tokenize(query)
	if err != nil {
		panic(err)
	}

	return search.ParseQuery(tokens).Evaluate(qe.attrs, qe.attrvals, qe.indices)
}
