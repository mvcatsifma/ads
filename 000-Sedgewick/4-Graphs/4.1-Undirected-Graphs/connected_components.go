package undirected


// connectedComponents finds all connected components in an undirected graph using DFS.
// A connected component is a maximal set of vertices where every pair is connected by a path.
// Returns the number of components and a 2D slice where each inner slice contains vertices in one component.
//
// Parameters:
//   - g: Graph to analyze for connected components
//
// Returns:
//   - int: Number of connected components in the graph
//   - [][]int: Array of components, where each component is a slice of vertex IDs belonging to that component
//
// Example: For graph with edges [(0,1), (2,3)], returns (2, [[0,1], [2,3]])
//          Two separate components: {0,1} and {2,3}
//
// Time Complexity: O(V + E) where V is vertices and E is edges (DFS visits each vertex and edge once).
// Space Complexity: O(V) for marked array, id array, result storage, and recursion stack.
func connectedComponents(g *Graph) (int, [][]int) {
	// Input validation: handle nil graph and empty graph
	if g == nil || g.V == 0 {
		return 0, [][]int{}
	}

	marked := make([]bool, g.V) // Track visited vertices during DFS
	id := make([]int, g.V)      // Map each vertex to its component ID (0, 1, 2, ...)
	count := 0                  // Current component ID counter

	// DFS helper function to explore and mark all vertices in a connected component
	var dfs func(v int)
	dfs = func(v int) {
		marked[v] = true // Mark vertex as visited
		id[v] = count    // Assign current component ID to vertex

		// Recursively visit all unvisited adjacent vertices
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				dfs(w)
			}
		}
	}

	// Phase 1: Find all connected components using DFS
	for s := range g.V { // Go 1.22+ syntax: iterate 0 to g.V-1
		if !marked[s] {
			dfs(s)  // Explore entire component starting from vertex s
			count++ // Increment component counter for next component
		}
	}

	// Phase 2: Group vertices by their component ID
	result := make([][]int, count) // Create slice for each component
	for c := range count {         // Go 1.22+ syntax: iterate 0 to count-1
		result[c] = make([]int, 0) // Initialize empty slice for component c
	}

	// Assign each vertex to its corresponding component slice
	for vertex, componentID := range id {
		result[componentID] = append(result[componentID], vertex)
	}

	return count, result
}
