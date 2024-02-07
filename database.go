package libgenders

import (
	"bufio"
	"os"
)

const DefaultGendersFilepath = "/etc/genders"

type Database struct {
	nodes []Node
	names map[string]int
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
			database.names[node.Name] = len(database.nodes) - 1
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
