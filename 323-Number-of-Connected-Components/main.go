package p323

// countComponents returns the number of connected components in an undirected graph.
// Each connected component is a maximal set of nodes reachable from each other.
// Time: O(V + E), Space: O(V) where V = n nodes, E = len(edges)
func countComponents(n int, edges [][]int) int {
	// Build undirected adjacency list representation
	adj := make([][]int, n)
	for _, edge := range edges {
		// Add bidirectional edges for undirected graph
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	visited := make([]bool, n) // Track which nodes have been explored

	// DFS to mark all nodes in current connected component
	var dfs func(v int)
	dfs = func(v int) {
		visited[v] = true // Mark current node as visited

		// Recursively visit all unvisited neighbors
		for _, w := range adj[v] {
			if !visited[w] {
				dfs(w) // Explore neighbor's component
			}
		}
	}

	componentCount := 0

	// Process each node to find all connected components
	for v := range n {
		if !visited[v] {
			// Found new component - explore it completely
			dfs(v)           // Mark all nodes in this component as visited
			componentCount++ // Increment component counter
		}
	}

	return componentCount
}
