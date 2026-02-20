package directed

import "leetcode/lib"


// Topological performs topological sorting on a directed graph.
// Returns whether a valid topological ordering exists (DAG check) and the three DFS orderings.
// For DAGs, actualReverse contains the topological ordering (dependencies before dependents).
// Time: O(V + E), Space: O(V) where V = vertices, E = edges
func Topological(g *Digraph) (hasOrder bool, actualPre []int, actualPost []int, actualReverse []int) {
	// Check if graph has cycles - topological ordering only exists for DAGs
	if hasCycleDirected(g) {
		return false, nil, nil, nil // Cyclic graph has no topological ordering
	}

	// Perform DFS to get all three orderings
	pre, post, reverse := DepthFirstOrder(g)

	// Convert data structures to slices for easier use
	actualPre = queueToSlice(pre)         // Pre-order: discovery order
	actualPost = queueToSlice(post)       // Post-order: completion order
	actualReverse = stackToSlice(reverse) // Reverse post-order: topological order

	return true, actualPre, actualPost, actualReverse // Valid topological ordering exists
}

// Helper function to convert queue to slice
func queueToSlice(q *lib.IntQueue) []int {
	var result []int
	for !q.IsEmpty() {
		result = append(result, q.Dequeue())
	}
	return result
}

// Helper function to convert stack to slice with error handling
func stackToSlice(s *lib.IntStack) []int {
	var result []int
	for !s.IsEmpty() {
		val, ok := s.Pop()
		if !ok {
			break // Stop if error (shouldn't happen if IsEmpty() works correctly)
		}
		result = append(result, val)
	}
	return result
}
