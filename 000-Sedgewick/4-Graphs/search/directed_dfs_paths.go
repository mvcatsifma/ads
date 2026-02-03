package search

import "leetcode/000-Sedgewick/4-Graphs/graphs"

// PathToVertexDirected finds a directed path from source vertex s to target vertex v.
// Returns the path as a slice of vertices, or nil if no path exists.
// Time: O(V + E), Space: O(V) where V = vertices, E = edges
func PathToVertexDirected(g *graphs.Digraph, s int, v int) []int {
	// Validate both source and target vertices are within valid range
	if s < 0 || s >= g.V || v < 0 || v >= g.V {
		return nil // Return nil for invalid vertices
	}

	// Track visited vertices and parent relationships for path reconstruction
	marked := make([]bool, g.V) // Tracks which vertices have been visited
	pathTo := make([]int, g.V)  // pathTo[w] = parent of w in DFS tree

	// Depth-first search to explore all reachable vertices from source
	var dfs func(v int)
	dfs = func(v int) {
		marked[v] = true // Mark current vertex as visited

		// Explore all adjacent vertices (following directed edges)
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				pathTo[w] = v // Record that we reached w from v
				dfs(w)        // Recursively explore from w
			}
		}
	}

	dfs(s) // Start DFS from source vertex

	// Check if target vertex is reachable from source
	if !marked[v] {
		return nil // No path exists from s to v
	}

	// Reconstruct path by following parent pointers backwards from target
	paths := make([]int, 0)
	for x := v; x != s; x = pathTo[x] {
		paths = append(paths, x) // Build path in reverse order (v -> ... -> s)
	}
	paths = append(paths, s) // Add source vertex

	// Reverse the path to get correct order (s -> ... -> v)
	for i, j := 0, len(paths)-1; i < j; {
		paths[i], paths[j] = paths[j], paths[i]
		i++
		j--
	}

	return paths // Return path from source to target
}
