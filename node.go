package main

import (
	"fmt"
)

// Node is a vertex in a graph.
type Node struct {
	name string
	data int
}

func (node Node) String() string {
	return fmt.Sprintf("Node %s: %d", node.name, node.data)
}
