package main

// Pair represents (a, b) value.
type Pair struct {
	first, second interface{}
}

// PriorityQueue is used in Dijkstra's algorithm.
type PriorityQueue struct {
	items []*Node
}

// CreatePriorityQueue is a constructor for the priority queue.
func CreatePriorityQueue() *PriorityQueue {
	return &PriorityQueue{make([]*Node, 0)}
}

// Enqueue adds a node to the queue.
func (queue *PriorityQueue) Enqueue(item *Node) {
	queue.items = append(queue.items, item)
}

// Dequeue removes the smallest distance node from the queue.
func (queue *PriorityQueue) Dequeue(distances map[*Node]int) (item *Node) {
	// Find the smallest node.
	smallestIndex := 0

	for i, node := range queue.items {
		if distances[node] < distances[queue.items[smallestIndex]] {
			smallestIndex = i
		}
	}

	// Remove the smalles node from the queue.
	queue.items[len(queue.items)-1], queue.items[smallestIndex] = queue.items[smallestIndex], queue.items[len(queue.items)-1]
	item = queue.items[len(queue.items)-1]
	queue.items = queue.items[:len(queue.items)-1]

	return item
}
