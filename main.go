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

	G.AddNeighbor(A, B, 1)
	G.AddNeighbor(A, C, 2)
	G.AddNeighbor(B, D, 3)
	G.AddNeighbor(D, A, 4)
	G.AddNeighbor(D, B, 5)

	fmt.Println(G)
}
