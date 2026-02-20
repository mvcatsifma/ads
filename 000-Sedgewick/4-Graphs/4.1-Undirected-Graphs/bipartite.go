package undirected


// isBipartite determines if an undirected graph is bipartite using 2-coloring
// A graph is bipartite if its vertices can be divided into two disjoint sets
// such that no two vertices within the same set are adjacent
func isBipartite(g *Graph) bool {
	isBipartite := true
	// Track which vertices have been visited during DFS traversal
	marked := make([]bool, g.V)
	// Assign colors (true/false) to vertices - adjacent vertices must have different colors
	color := make([]bool, g.V)

	// DFS helper function that attempts to 2-color the graph starting from vertex v
	var dfs func(g *Graph, v int)
	dfs = func(g *Graph, v int) {
		// Mark current vertex as visited
		marked[v] = true

		// Examine all adjacent vertices
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				// Unvisited vertex: assign opposite color to maintain bipartite property
				color[w] = !color[v]
				// Continue DFS from the newly colored vertex
				dfs(g, w)
			} else if color[v] == color[w] {
				// Conflict detected: adjacent vertices have same color
				// This means the graph cannot be properly 2-colored, so it's not bipartite
				// Example: odd-length cycles create this conflict
				isBipartite = false
			}
			// Note: if marked[w] && color[v] != color[w], no conflict - continue
		}
	}

	// Check all connected components in the graph
	// The graph is bipartite only if ALL components are bipartite
	// If any component is non-bipartite, the entire graph is non-bipartite
	for i := range g.V {
		if !marked[i] {
			// Start DFS from unvisited vertex
			// Initial color assignment (false) is arbitrary for the root of each component
			dfs(g, i)
		}
	}

	return isBipartite
}
