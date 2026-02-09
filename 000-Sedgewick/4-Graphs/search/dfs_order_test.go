package search

import (
	"testing"

	"leetcode/000-Sedgewick/4-Graphs/graphs"

	"github.com/stretchr/testify/assert"
)

func TestDepthFirstOrder(t *testing.T) {

	t.Run("empty graph", func(t *testing.T) {
		g := graphs.NewDigraph(0)
		pre, post, reverse := DepthFirstOrder(g)

		assert.True(t, pre.IsEmpty(), "Pre-order should be empty")
		assert.True(t, post.IsEmpty(), "Post-order should be empty")
		assert.True(t, reverse.IsEmpty(), "Reverse post-order should be empty")

		// Test that popping from empty stack returns error
		_, err := reverse.Pop()
		assert.Error(t, err, "Popping from empty stack should return error")
	})

	t.Run("single vertex no edges", func(t *testing.T) {
		g := graphs.NewDigraph(1)
		pre, post, reverse := DepthFirstOrder(g)

		// All orderings should contain just vertex 0
		assert.Equal(t, 0, pre.Dequeue(), "Pre-order should contain vertex 0")
		assert.True(t, pre.IsEmpty(), "Pre-order should have only one element")

		assert.Equal(t, 0, post.Dequeue(), "Post-order should contain vertex 0")
		assert.True(t, post.IsEmpty(), "Post-order should have only one element")

		val, err := reverse.Pop()
		assert.NoError(t, err, "Pop should not return error")
		assert.Equal(t, 0, val, "Reverse should contain vertex 0")
		assert.True(t, reverse.IsEmpty(), "Reverse should have only one element")
	})

	t.Run("linear chain", func(t *testing.T) {
		// Graph: 0→1→2→3
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)

		pre, post, reverse := DepthFirstOrder(g)

		// Pre-order: order of discovery (0,1,2,3)
		expectedPre := []int{0, 1, 2, 3}
		actualPre := queueToSlice(pre)
		assert.Equal(t, expectedPre, actualPre, "Pre-order should follow discovery order")

		// Post-order: order of completion (3,2,1,0)
		expectedPost := []int{3, 2, 1, 0}
		actualPost := queueToSlice(post)
		assert.Equal(t, expectedPost, actualPost, "Post-order should follow completion order")

		// Reverse post-order: topological order (0,1,2,3)
		expectedReverse := []int{0, 1, 2, 3}
		actualReverse := stackToSlice(reverse)
		assert.Equal(t, expectedReverse, actualReverse, "Reverse post-order should be topological order")
	})

	t.Run("tree structure", func(t *testing.T) {
		// Graph: 0→1, 0→2, 1→3, 1→4
		//        0
		//       / \
		//      1   2
		//     / \
		//    3   4
		g := graphs.NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(1, 4)

		pre, post, reverse := DepthFirstOrder(g)

		actualPre := queueToSlice(pre)
		actualPost := queueToSlice(post)
		actualReverse := stackToSlice(reverse)

		// Pre-order should start with root
		assert.Equal(t, 0, actualPre[0], "Pre-order should start with root vertex 0")

		// Post-order should end with root
		assert.Equal(t, 0, actualPost[len(actualPost)-1], "Post-order should end with root vertex 0")

		// Reverse post-order should start with root
		assert.Equal(t, 0, actualReverse[0], "Reverse post-order should start with root vertex 0")

		// All vertices should be present
		assert.Len(t, actualPre, 5, "Pre-order should contain all 5 vertices")
		assert.Len(t, actualPost, 5, "Post-order should contain all 5 vertices")
		assert.Len(t, actualReverse, 5, "Reverse post-order should contain all 5 vertices")
	})

	t.Run("disconnected components", func(t *testing.T) {
		// Two separate components: 0→1 and 2→3
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)

		pre, post, reverse := DepthFirstOrder(g)

		actualPre := queueToSlice(pre)
		actualPost := queueToSlice(post)
		actualReverse := stackToSlice(reverse)

		// All vertices should be visited
		assert.Len(t, actualPre, 4, "All vertices should be in pre-order")
		assert.Len(t, actualPost, 4, "All vertices should be in post-order")
		assert.Len(t, actualReverse, 4, "All vertices should be in reverse post-order")

		// Check that all vertices 0-3 are present
		assert.ElementsMatch(t, []int{0, 1, 2, 3}, actualPre, "Pre-order should contain all vertices")
		assert.ElementsMatch(t, []int{0, 1, 2, 3}, actualPost, "Post-order should contain all vertices")
		assert.ElementsMatch(t, []int{0, 1, 2, 3}, actualReverse, "Reverse post-order should contain all vertices")
	})

	t.Run("complex DAG", func(t *testing.T) {
		// DAG: 0→1, 0→2, 1→3, 2→3, 3→4
		g := graphs.NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)

		_, _, reverse := DepthFirstOrder(g)

		actualReverse := stackToSlice(reverse)

		// Reverse post-order should be a valid topological ordering
		// Vertex 0 should come before 1 and 2
		pos0 := findIndex(actualReverse, 0)
		pos1 := findIndex(actualReverse, 1)
		pos2 := findIndex(actualReverse, 2)
		pos3 := findIndex(actualReverse, 3)
		pos4 := findIndex(actualReverse, 4)

		assert.True(t, pos0 < pos1, "Vertex 0 should come before vertex 1 in topological order")
		assert.True(t, pos0 < pos2, "Vertex 0 should come before vertex 2 in topological order")
		assert.True(t, pos1 < pos3, "Vertex 1 should come before vertex 3 in topological order")
		assert.True(t, pos2 < pos3, "Vertex 2 should come before vertex 3 in topological order")
		assert.True(t, pos3 < pos4, "Vertex 3 should come before vertex 4 in topological order")
	})

	t.Run("single vertex with self loop", func(t *testing.T) {
		// Graph: 0→0 (self loop)
		g := graphs.NewDigraph(1)
		g.AddEdge(0, 0)

		pre, post, reverse := DepthFirstOrder(g)

		// Should visit vertex 0 only once
		actualPre := queueToSlice(pre)
		actualPost := queueToSlice(post)
		actualReverse := stackToSlice(reverse)

		assert.Equal(t, []int{0}, actualPre, "Pre-order should contain vertex 0 once")
		assert.Equal(t, []int{0}, actualPost, "Post-order should contain vertex 0 once")
		assert.Equal(t, []int{0}, actualReverse, "Reverse post-order should contain vertex 0 once")
	})

	t.Run("multiple isolated vertices", func(t *testing.T) {
		// Graph: 0, 1, 2 (no edges)
		g := graphs.NewDigraph(3)

		pre, post, reverse := DepthFirstOrder(g)

		actualPre := queueToSlice(pre)
		actualPost := queueToSlice(post)
		actualReverse := stackToSlice(reverse)

		// All vertices should be visited in order 0, 1, 2
		assert.Equal(t, []int{0, 1, 2}, actualPre, "Isolated vertices should be visited in order")
		assert.Equal(t, []int{0, 1, 2}, actualPost, "Isolated vertices should be completed in order")
		assert.Equal(t, []int{2, 1, 0}, actualReverse, "Reverse should be opposite order")
	})

	t.Run("diamond pattern", func(t *testing.T) {
		// Diamond: 0→1, 0→2, 1→3, 2→3
		//    0
		//   / \
		//  1   2
		//   \ /
		//    3
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3)

		pre, post, reverse := DepthFirstOrder(g)

		actualPre := queueToSlice(pre)
		actualPost := queueToSlice(post)
		actualReverse := stackToSlice(reverse)

		// Verify all vertices are present
		assert.ElementsMatch(t, []int{0, 1, 2, 3}, actualPre, "All vertices should be in pre-order")
		assert.ElementsMatch(t, []int{0, 1, 2, 3}, actualPost, "All vertices should be in post-order")
		assert.ElementsMatch(t, []int{0, 1, 2, 3}, actualReverse, "All vertices should be in reverse post-order")

		// Topological constraints
		pos0 := findIndex(actualReverse, 0)
		pos1 := findIndex(actualReverse, 1)
		pos2 := findIndex(actualReverse, 2)
		pos3 := findIndex(actualReverse, 3)

		assert.True(t, pos0 < pos1, "0 should come before 1")
		assert.True(t, pos0 < pos2, "0 should come before 2")
		assert.True(t, pos1 < pos3, "1 should come before 3")
		assert.True(t, pos2 < pos3, "2 should come before 3")
	})
}

// Helper function to convert queue to slice for testing
func queueToSlice(q *IntQueue) []int {
	var result []int
	for !q.IsEmpty() {
		result = append(result, q.Dequeue())
	}
	return result
}

// Helper function to convert stack to slice for testing
func stackToSlice(s *IntStack) []int {
	var result []int
	for !s.IsEmpty() {
		val, err := s.Pop()
		if err != nil {
			break // Stop if error (shouldn't happen if IsEmpty() works correctly)
		}
		result = append(result, val)
	}
	return result
}

// Helper function to find index of element in slice
func findIndex(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}
