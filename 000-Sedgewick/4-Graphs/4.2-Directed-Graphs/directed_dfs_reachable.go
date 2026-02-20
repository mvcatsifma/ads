package directed


// ReachableVerticesDirected finds all vertices reachable from a source vertex in a directed graph.
// Returns a boolean array indicating reachable vertices and the total count.
// Time: O(V + E), Space: O(V) where V = vertices, E = edges
func ReachableVerticesDirected(g *Digraph, source int) ([]bool, int) {
	// Validate source vertex is within valid range
	if source < 0 || source >= g.V {
		return make([]bool, g.V), 0 // Return all false if invalid source
	}

	// Track which vertices have been visited/reached
	marked := make([]bool, g.V) // Initially all false (unreachable)

	// Depth-first search to mark all reachable vertices
	var dfs func(v int)
	dfs = func(v int) {
		marked[v] = true // Mark current vertex as reachable

		// Explore all adjacent vertices (outgoing edges)
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				dfs(w) // Recursively visit unvisited neighbors
			}
		}
	}

	dfs(source) // Start DFS from the source vertex

	// Count total number of reachable vertices
	count := 0
	for _, m := range marked {
		if m {
			count++
		}
	}

	return marked, count // Return reachability array and count
}
