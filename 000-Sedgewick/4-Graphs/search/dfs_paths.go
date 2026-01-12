package search

import "leetcode/000-Sedgewick/4-Graphs/graphs"

// PathToVertex finds a path from source vertex s to target vertex v using DFS.
// Returns the path as a slice of vertices from source to target, or nil if no path exists.
// The path represents one possible route (not necessarily the shortest) due to DFS traversal.
//
// Parameters:
//   - g: Graph to search in
//   - s: Source vertex (starting point)
//   - v: Target vertex (destination)
//
// Returns:
//   - []int: Path from s to v as vertex sequence, or nil if unreachable or invalid input
//
// Example: PathToVertex(graph, 0, 3) might return [0, 1, 2, 3] representing path 0→1→2→3
//
// Time Complexity: O(V + E) where V is vertices and E is edges (DFS traversal).
// Space Complexity: O(V) for marked array, edgeTo array, and recursion stack.
func PathToVertex(g *graphs.Graph, s int, v int) []int {
	// Input validation: check vertex bounds
	if s >= g.V || s < 0 || v >= g.V || v < 0 {
		return nil
	}

	// Early return: source equals target
	if s == v {
		return []int{s}
	}

	marked := make([]bool, g.V) // Track visited vertices
	edgeTo := make([]int, g.V)  // Track parent of each vertex in DFS tree

	// DFS helper function to explore graph and build parent tree
	var dfs func(vertex int)
	dfs = func(vertex int) {
		marked[vertex] = true

		// Explore all adjacent vertices
		for _, w := range g.Adjacent(vertex) {
			if !marked[w] {
				edgeTo[w] = vertex // Record parent relationship
				dfs(w)             // Recursively explore
			}
		}
	}

	// Start DFS from source vertex
	dfs(s)

	// Check if target vertex is reachable from source
	if !marked[v] {
		return nil // No path exists
	}

	// Reconstruct path by following parent pointers from target back to source
	var path []int
	for x := v; x != s; x = edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, s)

	// Reverse path to get source-to-target order
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
