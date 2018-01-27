package main

// MaxDistance describes that there is no connection.
var MaxDistance = 9999

// FindShortestPath utilizes Dijkstra's algorithm for finding the shortest path
// from the source node to the destination node. Return both the path and the distance.
func FindShortestPath(graph *Graph, source *Node, destination *Node) (nodes []*Node, distance int) {
	// Initialize queue and other data structures.
	queue := CreatePriorityQueue()
	distances := make(map[*Node]int)
	previous := make(map[*Node]*Node)

	// Set all distances to maximum.
	for node := range graph.edges {
		distances[node] = MaxDistance
		previous[node] = nil
		queue.Enqueue(node)
	}

	// Start position distance is zero.
	distances[source] = 0

	// While there are nodes in the queue,
	// take the smallest node from the priority queue
	// and update distances if possible.
	for len(queue.items) > 0 {
		node := queue.Dequeue(distances)

		for _, neighbor := range graph.edges[node] {
			newDistance := distances[node] + neighbor.second.(int)
			if newDistance < distances[neighbor.first.(*Node)] {
				distances[neighbor.first.(*Node)] = newDistance
				previous[neighbor.first.(*Node)] = node
			}
		}
	}

	// Create a slice describing the shortest path.
	path := make([]*Node, 0)
	node := destination
	for node != nil {
		path = append(path, node)
		node = previous[node]
	}

	return path, distances[destination]
}
