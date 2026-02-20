package directed


// KosarajuSharirSCC determines if two vertices are in the same strongly connected component (SCC).
//
// Algorithm Overview:
// A strongly connected component is a maximal set of vertices where every vertex is reachable
// from every other vertex via directed paths. The Kosaraju-Sharir algorithm finds SCCs in two phases:
//
// Phase 1: Compute finish times on the reversed graph
//   - Create G^R (transpose of G where all edges are reversed)
//   - Run DFS on G^R to compute reverse post-order (decreasing finish times)
//   - This ordering ensures we process SCCs in topological order
//
// Phase 2: Find SCCs on the original graph
//   - Process vertices from Phase 1 in reverse post-order
//   - For each unvisited vertex, run DFS on original graph G
//   - All vertices reached in one DFS call form one SCC
//   - Assign same component ID to vertices in the same SCC
//
// Key Insight: The reverse post-order from G^R gives us an ordering where if there's
// an edge between SCCs C1→C2, then all vertices in C1 are processed before C2.
// This prevents DFS from "leaking" between different SCCs.
//
// Example: Graph 0→1→2→0, 3→4→3
//   - SCCs: {0,1,2} and {3,4}
//   - Returns true for (0,1), (1,2), (3,4)
//   - Returns false for (0,3), (1,4), etc.
//
// Time: O(V + E), Space: O(V) where V = vertices, E = edges
func KosarajuSharirSCC(g *Digraph, v int, w int) bool {
	// Validate input vertices are within valid range
	if v < 0 || v >= g.V || w < 0 || w >= g.V {
		return false // Invalid vertices cannot be in same SCC
	}

	marked := make([]bool, g.V) // Track visited vertices in second DFS pass
	id := make([]int, g.V)      // SCC identifier for each vertex
	count := 0                  // Number of SCCs found so far

	// DFS to assign SCC identifiers to vertices
	var dfs func(v int)
	dfs = func(v int) {
		marked[v] = true // Mark vertex as visited
		id[v] = count    // Assign current SCC identifier

		// Explore all adjacent vertices in original graph
		for _, a := range g.Adjacent(v) {
			if !marked[a] {
				dfs(a) // Recursively process unvisited neighbors
			}
		}
	}

	// Phase 1: Get reverse post-order from reversed graph G^R
	// This provides vertices in decreasing finish time order, ensuring
	// we process SCCs in topological order (no edges between processed SCCs)
	_, _, reverse := DepthFirstOrder(g.Reverse())

	// Phase 2: Process vertices in reverse post-order on original graph G
	// Each DFS call finds exactly one complete SCC
	for !reverse.IsEmpty() {
		item, _ := reverse.Pop() // Get next vertex in decreasing finish time order
		if !marked[item] {
			// Found new SCC - run DFS to mark all vertices in this component
			dfs(item)
			count++ // Increment SCC counter for next component
		}
	}

	// Two vertices are in same SCC if they have same identifier
	return id[v] == id[w]
}
