package main

import (
	"fmt"
)

func main() {
	//Create graph.
	G := CreateGraph("G")

	A := &Node{"A", 1}
	B := &Node{"B", 2}
	C := &Node{"C", 3}
	D := &Node{"D", 4}
	E := &Node{"E", 5}
	F := &Node{"F", 6}
	H := &Node{"H", 7}
	I := &Node{"I", 8}

	G.AddNode(A)
	G.AddNode(B)
	G.AddNode(C)
	G.AddNode(D)
	G.AddNode(E)
	G.AddNode(F)
	G.AddNode(H)
	G.AddNode(I)

	G.AddNeighbor(A, B, 1)
	G.AddNeighbor(A, C, 2)
	G.AddNeighbor(B, D, 3)
	G.AddNeighbor(D, A, 4)
	G.AddNeighbor(D, B, 5)
	G.AddNeighbor(D, H, 4)
	G.AddNeighbor(D, F, 2)
	G.AddNeighbor(F, H, 1)
	G.AddNeighbor(A, H, 10)

	fmt.Println(G)

	// Calculate the shortest path from A to H.
	shortestPath, distanceAH := FindShortestPath(G, A, H)
	fmt.Printf("Shortest path from A to H is:\n%s\n", Path{shortestPath})
	fmt.Printf("Shortest distance from A to H is: %d\n", distanceAH)
}
