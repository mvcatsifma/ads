package search

import "leetcode/000-Sedgewick/4-Graphs/graphs"

// ShortestPathToVertex finds the shortest path from source vertex s to target vertex t using BFS.
// In an unweighted graph, BFS guarantees the shortest path (minimum number of edges).
// Returns the path as a slice of vertices from source to target, or nil if no path exists.
//
// Parameters:
//   - g: Graph to search in
//   - s: Source vertex (starting point)
//   - t: Target vertex (destination)
//
// Returns:
//   - []int: Shortest path from s to t as vertex sequence, or nil if unreachable or invalid input
//
// Example: ShortestPathToVertex(graph, 0, 3) returns [0, 1, 3] if that's the shortest route
//
// Time Complexity: O(V + E) where V is vertices and E is edges (BFS traversal).
// Space Complexity: O(V) for marked array, edgeTo array, and queue storage.
func ShortestPathToVertex(g *graphs.Graph, s int, t int) []int {
	// Input validation: check vertex bounds
	if s < 0 || s >= g.V || t < 0 || t >= g.V {
		return nil
	}

	// Early return: source equals target
	if s == t {
		return []int{s}
	}

	marked := make([]bool, g.V) // Track visited vertices
	edgeTo := make([]int, g.V)  // Track parent of each vertex in BFS tree

	// BFS using queue to explore vertices level by level
	queue := &IntQueue{}
	queue.Enqueue(s)
	marked[s] = true

	// BFS traversal to build shortest path tree
	for !queue.IsEmpty() {
		v := queue.Dequeue()

		// Explore all adjacent vertices
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				edgeTo[w] = v    // Record parent relationship
				marked[w] = true // Mark as visited
				queue.Enqueue(w) // Add to queue for next level
			}
		}
	}

	// Check if target vertex is reachable from source
	if !marked[t] {
		return nil // No path exists
	}

	// Reconstruct path by following parent pointers from target back to source
	var path []int
	for x := t; x != s; x = edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, s)

	// Reverse path to get source-to-target order
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// IntQueue implements a simple FIFO queue for BFS traversal using head pointer optimization.
// Provides efficient O(1) enqueue and dequeue operations.
type IntQueue struct {
	items []int // Underlying slice storing queue elements
	head  int   // Index of first element (front of queue)
}

// Enqueue adds an item to the back of the queue.
// Time Complexity: O(1) amortized.
func (q *IntQueue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue.
// Assumes queue is not empty (caller should check IsEmpty() first).
// Time Complexity: O(1).
func (q *IntQueue) Dequeue() int {
	if q.IsEmpty() {
		panic("dequeue from empty queue")
	}
	item := q.items[q.head]
	q.head++
	return item
}

// IsEmpty returns true if the queue contains no elements.
// Time Complexity: O(1).
func (q *IntQueue) IsEmpty() bool {
	return q.head == len(q.items)
}
