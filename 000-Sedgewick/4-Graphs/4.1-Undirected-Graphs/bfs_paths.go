package undirected

import "leetcode/lib"


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
func ShortestPathToVertex(g *Graph, s int, t int) []int {
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
	queue := &lib.IntQueue{}
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