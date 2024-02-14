package libgenders

import (
	"bufio"
	"os"
)

const DefaultGendersFilepath = "/etc/genders"

type Database struct {
	nodes []Node
	names map[string]int

	engine QueryEngine
}

func NewDatabase(path string) (Database, error) {
	file, err := os.Open(path)
	if err != nil {
		return Database{}, err
	}
	defer file.Close()

	database := Database{
		nodes: []Node{},
		names: make(map[string]int),
	}

	scanner := bufio.NewScanner(file)
	parser := NewParser()
	for scanner.Scan() {
		line := scanner.Text()
		nodes, err := parser.Parse(line)
		if err != nil {
			panic(err)
		}

		for _, node := range nodes {
			if index, ok := database.names[node.Name]; ok {
				database.nodes[index].MergeAttributes(node.Attributes)
				continue
			}

			database.nodes = append(database.nodes, node)
			index := len(database.nodes) - 1
			database.names[node.Name] = index
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	database.engine = NewQueryEngine(database.nodes)

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
	var nodes []Node
	indices := d.engine.Query(query)

	for _, index := range indices {
		nodes = append(nodes, d.nodes[index])
	}

	return nodes, nil
}
