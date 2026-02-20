package directed

// ShortestPathToVertexDirected finds the shortest path from source s to target t in a directed graph.
// Uses BFS to guarantee the shortest path in unweighted graphs (minimum number of edges).
// Returns the path as a slice of vertices, or nil if no path exists.
// Time: O(V + E), Space: O(V) where V = vertices, E = edges
func ShortestPathToVertexDirected(g *Digraph, s int, t int) []int {
	// Validate both source and target vertices are within valid range
	if s < 0 || s >= g.V || t < 0 || t >= g.V {
		return nil // Return nil for invalid vertices
	}

	// BFS data structures
	queue := &IntQueue{}        // Queue for BFS traversal
	marked := make([]bool, g.V) // Track visited vertices
	pathTo := make([]int, g.V)  // Parent pointers for path reconstruction

	// Initialize BFS from source vertex
	queue.Enqueue(s)
	marked[s] = true // Mark source as visited when enqueued

	// BFS traversal to find shortest paths
	for !queue.IsEmpty() {
		v := queue.Dequeue() // Process next vertex in queue

		// Explore all adjacent vertices (following directed edges)
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				pathTo[w] = v    // Record parent for path reconstruction
				marked[w] = true // Mark as visited to prevent re-queuing
				queue.Enqueue(w) // Add to queue for future processing
			}
		}
	}

	// Check if target is reachable from source
	if !marked[t] {
		return nil // No path exists from s to t
	}

	// Reconstruct shortest path by following parent pointers backwards
	var path []int
	for x := t; x != s; x = pathTo[x] {
		path = append(path, x) // Build path in reverse order (t -> ... -> s)
	}
	path = append(path, s) // Add source vertex

	// Reverse path to get correct order (s -> ... -> t)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path // Return shortest path from source to target
}
