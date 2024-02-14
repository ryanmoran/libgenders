# Libgenders

An implementation of [libgenders](https://github.com/chaos/genders/tree/master/src/libgenders) in Go.

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/ryanmoran/libgenders"
)

func main() {
    database, err := libgenders.NewDatabase(libgenders.DefaultGendersFilepath)
	if err != nil {
		log.Fatal(err)
	}

	value, ok := database.GetNodeAttr("node1", "attr2")
	if ok {
		fmt.Println(value)
	}

	nodes, err := database.Query("~(attr1 -- ((attr1 && attr3) || (attr1 && attr5)))")
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {
		fmt.Printf("name: %s, attributes: %v\n", node.Name, node.Attributes)
	}
}
```
