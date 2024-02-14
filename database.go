package libgenders

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ryanmoran/libgenders/internal"
)

const DefaultGendersFilepath = "/etc/genders"

type Database struct {
	nodes []Node
	names map[string]int

	attrs    map[string]internal.Set
	attrvals map[string]internal.Set
	indices  internal.Set
}

func NewDatabase(path string) (Database, error) {
	file, err := os.Open(path)
	if err != nil {
		return Database{}, err
	}
	defer file.Close()

	database := Database{
		nodes:    []Node{},
		names:    make(map[string]int),
		attrs:    make(map[string]internal.Set),
		attrvals: make(map[string]internal.Set),
	}

	scanner := bufio.NewScanner(file)
	var parser internal.Parser
	for scanner.Scan() {
		line := scanner.Text()
		nodes, err := parser.Parse(line)
		if err != nil {
			panic(err)
		}

		for _, node := range nodes {
			if index, ok := database.names[node.Name]; ok {
				database.nodes[index].mergeAttributes(node.Attributes)
				continue
			}

			database.nodes = append(database.nodes, Node(node))
			index := len(database.nodes) - 1
			database.names[node.Name] = index
		}
	}

	database.indices = make(internal.Set, len(database.nodes))
	for index, node := range database.nodes {
		database.indices[index] = index
		for key, value := range node.Attributes {
			database.attrs[key] = append(database.attrs[key], index)

			if value != "" {
				keyval := fmt.Sprintf("%s=%s", key, value)
				database.attrvals[keyval] = append(database.attrvals[keyval], index)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return database, nil
}

func (d Database) GetNodes() []Node {
	return d.nodes
}

func (d Database) GetNodeAttr(name, attr string) (string, bool) {
	if index, ok := d.names[name]; ok {
		val, ok := d.nodes[index].Attributes[attr]
		return val, ok
	}

	return "", false
}

func (d Database) Query(query string) ([]Node, error) {
	tokens, err := internal.Tokenize(query)
	if err != nil {
		panic(err)
	}

	var nodes []Node
	for _, index := range internal.ParseQuery(tokens).Evaluate(d.attrs, d.attrvals, d.indices) {
		nodes = append(nodes, d.nodes[index])
	}

	return nodes, nil
}
