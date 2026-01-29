package p684

// findRedundantConnection finds the last edge that creates a cycle in the graph.
// Uses Union-Find with path compression and union by size.
// Time: O(n * α(n)), Space: O(n) where α is inverse Ackermann function
func findRedundantConnection(edges [][]int) []int {
	// Number of nodes (edges connect nodes 1 to n)
	n := len(edges)

	// Parent array for Union-Find: par[i] points to parent of node i
	par := make([]int, n+1)
	for i := range par {
		par[i] = i // Each node is its own parent initially (disjoint sets)
	}

	// Size array: tracks number of nodes in each connected component
	size := make([]int, n+1)
	for i := range size {
		size[i] = 1 // Each component starts with size 1 (single node)
	}

	// Find operation with path compression
	// Returns root of the set containing node n
	var find func(n int) int
	find = func(n int) int {
		if n != par[n] {
			par[n] = find(par[n]) // Path compression: point directly to root
		}
		return par[n]
	}

	// Union operation with union by size
	// Returns false if nodes are already connected (cycle detected)
	union := func(n1 int, n2 int) bool {
		p1, p2 := find(n1), find(n2) // Find roots of both components
		if p1 == p2 {
			return false // Same component = adding edge creates cycle
		}

		// Union by size: attach smaller component to larger one for balance
		if size[p1] > size[p2] {
			par[p2] = p1         // Make p1 the new root
			size[p1] += size[p2] // Update size of merged component
		} else {
			par[p1] = p2         // Make p2 the new root
			size[p2] += size[p1] // Update size of merged component
		}

		return true // Successfully merged components
	}

	// Process edges in order - first edge that creates cycle is redundant
	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge // Found redundant edge that creates cycle
		}
	}

	return nil // Should never reach here for valid input (tree + 1 edge)
}
