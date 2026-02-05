package search

import "leetcode/000-Sedgewick/4-Graphs/graphs"

// hasCycleDirected detects if a directed graph contains any cycles using DFS.
// Uses white/gray/black vertex coloring: white=unvisited, gray=on stack, black=finished.
// A cycle exists if we find a back edge (edge to a vertex currently on the DFS stack).
// Time: O(V + E), Space: O(V) where V = vertices, E = edges
func hasCycleDirected(g *graphs.Digraph) bool {
	marked := make([]bool, g.V)  // Tracks visited vertices (white → gray/black)
	onStack := make([]bool, g.V) // Tracks vertices on current DFS recursion stack (gray)

	// DFS with recursion stack tracking for cycle detection
	var dfs func(v int) bool
	dfs = func(v int) bool {
		onStack[v] = true // Mark vertex as gray (on current DFS path)
		marked[v] = true  // Mark vertex as visited

		// Explore all outgoing edges from current vertex
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				// White vertex: unvisited neighbor, recurse
				if dfs(w) {
					return true // Cycle found in subtree
				}
			} else if onStack[w] {
				// Gray vertex: back edge detected = cycle found
				return true
			}
			// Black vertex: cross/forward edge, safe to ignore
		}

		onStack[v] = false // Mark vertex as black (finished processing)
		return false       // No cycle found from this vertex
	}

	// Check all disconnected components in the graph
	for v := range g.V {
		if !marked[v] && dfs(v) {
			return true // Cycle detected in this component
		}
	}

	return false // No cycles found in entire graph
}
