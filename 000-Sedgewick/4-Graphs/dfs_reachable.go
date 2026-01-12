package p4Graphs

// ReachableVertices finds all vertices reachable from a source vertex using DFS.
// Returns a boolean array indicating reachability and the total count of reachable vertices.
// A vertex is reachable if there exists a path from the source to that vertex.
//
// Parameters:
//   - g: Graph to search in
//   - source: Starting vertex for reachability analysis
//
// Returns:
//   - []bool: Array where marked[i] is true if vertex i is reachable from source
//   - int: Total count of reachable vertices (including source itself)
//
// Example: ReachableVertices(graph, 0) might return ([true, true, false, true], 3)
//          meaning vertices 0, 1, and 3 are reachable from vertex 0
//
// Time Complexity: O(V + E) where V is vertices and E is edges (DFS traversal).
// Space Complexity: O(V) for marked array and recursion stack.
func ReachableVertices(g *Graph, source int) ([]bool, int) {
	// Input validation: check source vertex bounds
	if source < 0 || source >= g.V {
		return make([]bool, g.V), 0 // Return empty result for invalid input
	}

	marked := make([]bool, g.V) // Track visited/reachable vertices

	// DFS helper function to explore all reachable vertices
	var dfs func(v int)
	dfs = func(v int) {
		marked[v] = true

		// Explore all unvisited adjacent vertices
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				dfs(w)
			}
		}
	}

	// Start DFS from source vertex
	dfs(source)

	// Count total reachable vertices
	count := 0
	for _, isReachable := range marked {
		if isReachable {
			count++
		}
	}

	return marked, count
}
