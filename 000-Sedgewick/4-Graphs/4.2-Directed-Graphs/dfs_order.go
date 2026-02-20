package directed

import (
	"leetcode/lib"
)

// DepthFirstOrder performs DFS traversal and returns vertices in three different orderings.
// Returns pre-order (when first visited), post-order (when finished), and reverse post-order.
// Reverse post-order is equivalent to topological ordering for DAGs (Directed Acyclic Graphs).
// Time: O(V + E), Space: O(V) where V = vertices, E = edges
func DepthFirstOrder(g *Digraph) (pre *lib.IntQueue, post *lib.IntQueue, reverse *lib.IntStack) {
	// Initialize data structures for the three orderings
	pre = &lib.IntQueue{}     // Pre-order: vertices in order of first discovery
	post = &lib.IntQueue{}    // Post-order: vertices in order of completion
	reverse = lib.NewIntStack(g.V) // Reverse post-order: topological order for DAGs

	marked := make([]bool, g.V) // Track visited vertices

	// DFS traversal with ordering tracking
	var dfs func(v int)
	dfs = func(v int) {
		pre.Enqueue(v)   // Record vertex when first discovered (pre-order)
		marked[v] = true // Mark as visited to prevent revisiting

		// Recursively visit all unvisited adjacent vertices
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				dfs(w) // Explore unvisited neighbor
			}
		}

		// Record vertex when completely processed (post-order)
		post.Enqueue(v) // Post-order: added after all descendants processed
		reverse.Push(v) // Reverse post-order: stack reverses the post-order
	}

	// Process all vertices to handle disconnected components
	for v := range g.V {
		if !marked[v] {
			dfs(v) // Start DFS from unvisited vertex
		}
	}

	return // Named return values
}
