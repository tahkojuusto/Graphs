package main

import (
	"fmt"
)

// Node is a vertex in a graph.
type Node struct {
	name string
	data int
}

// Path is a slice of connected nodes.
type Path struct {
	nodes []*Node
}

func (node Node) String() string {
	return fmt.Sprintf("Node %s: %d", node.name, node.data)
}

func (path Path) String() string {
	description := ""
	for i := len(path.nodes) - 1; i >= 0; i-- {
		description += "(" + path.nodes[i].String() + ")"

		if i > 0 {
			description += " -> "
		}
	}

	return description
}
