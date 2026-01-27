package p261

// validTree determines if n nodes and given edges form a valid tree.
// A valid tree must have exactly n-1 edges, be connected (1 component), and be acyclic.
// Time: O(V + E), Space: O(V) where V = n nodes, E = len(edges)
func validTree(n int, edges [][]int) bool {
	// Early check: valid tree must have exactly n-1 edges
	if len(edges) != n-1 {
		return false
	}

	// Build undirected adjacency list
	adj := make([][]int, n)
	for _, edge := range edges {
		// Add bidirectional edges since tree is undirected
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	numberOfComponents := 0
	hasCycle := false
	visited := make([]bool, n)

	// DFS to detect cycles in undirected graph
	// parent parameter prevents false cycle detection on the edge we came from
	var dfs func(v int, parent int)
	dfs = func(v int, parent int) {
		visited[v] = true

		// Explore all neighbors
		for _, w := range adj[v] {
			if !visited[w] {
				// Unvisited neighbor - continue DFS
				dfs(w, v) // v becomes parent of w
			} else if parent != w {
				// Visited neighbor that's not our parent = back edge = cycle
				// In undirected graph, we ignore the edge we came from (parent)
				hasCycle = true
			}
		}
	}

	// Check all nodes to count connected components
	for i := range n {
		if !visited[i] {
			dfs(i, -1) // Start DFS with no parent (-1 or any invalid node ID)
			numberOfComponents++
		}
	}

	// Valid tree conditions:
	// 1. No cycles (acyclic)
	// 2. Exactly one connected component (connected)
	// Note: Could also check len(edges) == n-1 as early optimization
	return !hasCycle && numberOfComponents == 1
}
