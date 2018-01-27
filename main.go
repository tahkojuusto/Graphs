package main

import (
	"fmt"
)

func main() {
	G := CreateGraph("G")

	A := &Node{"A", 1}
	B := &Node{"B", 2}
	C := &Node{"C", 3}
	D := &Node{"D", 4}
	E := &Node{"E", 5}

	G.AddNode(A)
	G.AddNode(B)
	G.AddNode(C)
	G.AddNode(D)
	G.AddNode(E)

	G.AddNeighbor(A, B)
	G.AddNeighbor(A, C)
	G.AddNeighbor(B, D)
	G.AddNeighbor(D, A)
	G.AddNeighbor(D, B)

	fmt.Println(G)
}
