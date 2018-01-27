package main

import (
	"fmt"
)

// Graph consists of nodes and edges.
type Graph struct {
	name  string
	edges map[*Node][]*Node
}

// CreateGraph creates a new graph (constructor).
func CreateGraph(name string) *Graph {
	return &Graph{name, make(map[*Node][]*Node)}
}

// AddNode pushes a node to the graph.
func (graph *Graph) AddNode(node *Node) {
	graph.edges[node] = make([]*Node, 0)
}

// AddNeighbor connects a source node to a destination node.
func (graph *Graph) AddNeighbor(node *Node, neighbor *Node) {
	graph.edges[node] = append(graph.edges[node], neighbor)
}

func (graph Graph) String() string {
	var description = fmt.Sprintf("Graph %s:\n", graph.name)
	for source, destinations := range graph.edges {
		description += fmt.Sprintf("(%s) -> ", source)
		for _, destination := range destinations {
			description += fmt.Sprintf("(%s) ", destination)
		}
		description += "\n"
	}

	return description
}
