package undirected


// hasCycle detects whether an undirected graph contains a cycle
// Uses depth-first search to identify back edges that form cycles
func hasCycle(g *Graph) bool {
	// Track which vertices have been visited during DFS traversal
	marked := make([]bool, g.V)
	hasCycle := false

	// DFS helper function that explores from vertex v
	// parent parameter tracks the vertex we came from to avoid false cycle detection
	var dfs func(g *Graph, v int, parent int)
	dfs = func(g *Graph, v int, parent int) {
		// Mark current vertex as visited
		marked[v] = true

		// Examine all adjacent vertices
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				// Unvisited vertex: continue DFS with v as parent of w
				dfs(g, w, v)
			} else if w != parent {
				// Visited vertex that's not our immediate parent = back edge = cycle!
				// This means we've found an alternative path back to a visited vertex
				hasCycle = true
			}
			// Note: if w == parent, we ignore it (that's the edge we just came from)
		}
	}

	// Check all connected components in the graph
	// A graph may be disconnected, so we need to start DFS from each unvisited vertex
	for v := 0; v < g.V; v++ {
		if !marked[v] {
			// Start DFS from unvisited vertex v
			// Initially, parent is set to v itself (arbitrary choice for root)
			dfs(g, v, v)
		}
	}

	return hasCycle
}
