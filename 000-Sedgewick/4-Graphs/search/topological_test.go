package search

import (
	"testing"

	"leetcode/000-Sedgewick/4-Graphs/graphs"

	"github.com/stretchr/testify/assert"
)

func TestTopological(t *testing.T) {

	t.Run("empty graph", func(t *testing.T) {
		g := graphs.NewDigraph(0)
		hasOrder, pre, post, reverse := Topological(g)

		assert.True(t, hasOrder, "Empty graph should have valid topological ordering")
		assert.Empty(t, pre, "Pre-order should be empty")
		assert.Empty(t, post, "Post-order should be empty")
		assert.Empty(t, reverse, "Reverse post-order should be empty")
	})

	t.Run("single vertex no edges", func(t *testing.T) {
		g := graphs.NewDigraph(1)
		hasOrder, pre, post, reverse := Topological(g)

		assert.True(t, hasOrder, "Single vertex should have valid topological ordering")
		assert.Equal(t, []int{0}, pre, "Pre-order should contain vertex 0")
		assert.Equal(t, []int{0}, post, "Post-order should contain vertex 0")
		assert.Equal(t, []int{0}, reverse, "Topological order should contain vertex 0")
	})

	t.Run("linear chain DAG", func(t *testing.T) {
		// Graph: 0→1→2→3 (valid DAG)
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)

		hasOrder, pre, post, reverse := Topological(g)

		assert.True(t, hasOrder, "Linear chain should have valid topological ordering")
		assert.Equal(t, []int{0, 1, 2, 3}, pre, "Pre-order should follow discovery order")
		assert.Equal(t, []int{3, 2, 1, 0}, post, "Post-order should follow completion order")
		assert.Equal(t, []int{0, 1, 2, 3}, reverse, "Topological order should respect dependencies")

		// Verify topological ordering constraints
		assert.True(t, indexOf(reverse, 0) < indexOf(reverse, 1), "0 should come before 1")
		assert.True(t, indexOf(reverse, 1) < indexOf(reverse, 2), "1 should come before 2")
		assert.True(t, indexOf(reverse, 2) < indexOf(reverse, 3), "2 should come before 3")
	})

	t.Run("simple cycle", func(t *testing.T) {
		// Graph: 0→1→2→0 (cycle)
		g := graphs.NewDigraph(3)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 0)

		hasOrder, pre, post, reverse := Topological(g)

		assert.False(t, hasOrder, "Cyclic graph should not have topological ordering")
		assert.Nil(t, pre, "Pre-order should be nil for cyclic graph")
		assert.Nil(t, post, "Post-order should be nil for cyclic graph")
		assert.Nil(t, reverse, "Topological order should be nil for cyclic graph")
	})

	t.Run("self loop", func(t *testing.T) {
		// Graph: 0→0 (self loop = cycle)
		g := graphs.NewDigraph(1)
		g.AddEdge(0, 0)

		hasOrder, pre, post, reverse := Topological(g)

		assert.False(t, hasOrder, "Self loop should prevent topological ordering")
		assert.Nil(t, pre, "Pre-order should be nil for cyclic graph")
		assert.Nil(t, post, "Post-order should be nil for cyclic graph")
		assert.Nil(t, reverse, "Topological order should be nil for cyclic graph")
	})

	t.Run("tree structure DAG", func(t *testing.T) {
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

		hasOrder, _, _, reverse := Topological(g)

		assert.True(t, hasOrder, "Tree structure should have valid topological ordering")
		assert.Len(t, reverse, 5, "Topological order should contain all 5 vertices")

		// Verify topological constraints
		assert.True(t, indexOf(reverse, 0) < indexOf(reverse, 1), "0 should come before 1")
		assert.True(t, indexOf(reverse, 0) < indexOf(reverse, 2), "0 should come before 2")
		assert.True(t, indexOf(reverse, 1) < indexOf(reverse, 3), "1 should come before 3")
		assert.True(t, indexOf(reverse, 1) < indexOf(reverse, 4), "1 should come before 4")

		// Root should be first in topological order
		assert.Equal(t, 0, reverse[0], "Root vertex 0 should be first in topological order")
	})

	t.Run("diamond DAG", func(t *testing.T) {
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

		hasOrder, _, _, reverse := Topological(g)

		assert.True(t, hasOrder, "Diamond DAG should have valid topological ordering")
		assert.Len(t, reverse, 4, "Topological order should contain all 4 vertices")

		// Verify topological constraints
		assert.True(t, indexOf(reverse, 0) < indexOf(reverse, 1), "0 should come before 1")
		assert.True(t, indexOf(reverse, 0) < indexOf(reverse, 2), "0 should come before 2")
		assert.True(t, indexOf(reverse, 1) < indexOf(reverse, 3), "1 should come before 3")
		assert.True(t, indexOf(reverse, 2) < indexOf(reverse, 3), "2 should come before 3")

		// Check start and end vertices
		assert.Equal(t, 0, reverse[0], "Vertex 0 should be first")
		assert.Equal(t, 3, reverse[len(reverse)-1], "Vertex 3 should be last")
	})

	t.Run("complex DAG", func(t *testing.T) {
		// Complex DAG: 0→1, 0→2, 1→3, 2→3, 3→4, 1→4
		g := graphs.NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(1, 4)

		hasOrder, _, _, reverse := Topological(g)

		assert.True(t, hasOrder, "Complex DAG should have valid topological ordering")
		assert.Len(t, reverse, 5, "Topological order should contain all 5 vertices")

		// Verify all dependency constraints
		assert.True(t, indexOf(reverse, 0) < indexOf(reverse, 1), "0 should come before 1")
		assert.True(t, indexOf(reverse, 0) < indexOf(reverse, 2), "0 should come before 2")
		assert.True(t, indexOf(reverse, 1) < indexOf(reverse, 3), "1 should come before 3")
		assert.True(t, indexOf(reverse, 2) < indexOf(reverse, 3), "2 should come before 3")
		assert.True(t, indexOf(reverse, 3) < indexOf(reverse, 4), "3 should come before 4")
		assert.True(t, indexOf(reverse, 1) < indexOf(reverse, 4), "1 should come before 4")
	})

	t.Run("disconnected DAG components", func(t *testing.T) {
		// Two separate DAG components: 0→1 and 2→3
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)

		hasOrder, _, _, reverse := Topological(g)

		assert.True(t, hasOrder, "Disconnected DAG components should have valid topological ordering")
		assert.Len(t, reverse, 4, "Topological order should contain all 4 vertices")
		assert.ElementsMatch(t, []int{0, 1, 2, 3}, reverse, "All vertices should be present")

		// Verify component constraints
		assert.True(t, indexOf(reverse, 0) < indexOf(reverse, 1), "0 should come before 1")
		assert.True(t, indexOf(reverse, 2) < indexOf(reverse, 3), "2 should come before 3")
	})

	t.Run("disconnected with cycle", func(t *testing.T) {
		// One DAG component (0→1) and one cycle (2→3→2)
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)
		g.AddEdge(3, 2)

		hasOrder, pre, post, reverse := Topological(g)

		assert.False(t, hasOrder, "Graph with any cycle should not have topological ordering")
		assert.Nil(t, pre, "Pre-order should be nil when cycle exists")
		assert.Nil(t, post, "Post-order should be nil when cycle exists")
		assert.Nil(t, reverse, "Topological order should be nil when cycle exists")
	})

	t.Run("isolated vertices", func(t *testing.T) {
		// Graph: 0, 1, 2 (no edges)
		g := graphs.NewDigraph(3)

		hasOrder, _, _, reverse := Topological(g)

		assert.True(t, hasOrder, "Isolated vertices should have valid topological ordering")
		assert.Equal(t, []int{0, 1, 2}, reverse, "Isolated vertices should be in natural order")
		assert.Len(t, reverse, 3, "All vertices should be present")
	})

	t.Run("course scheduling example", func(t *testing.T) {
		// Course dependencies:
		// Course 0 (Intro) → Course 1 (Intermediate)
		// Course 0 (Intro) → Course 2 (Math)
		// Course 1 (Intermediate) → Course 3 (Advanced)
		// Course 2 (Math) → Course 3 (Advanced)
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1) // Intro → Intermediate
		g.AddEdge(0, 2) // Intro → Math
		g.AddEdge(1, 3) // Intermediate → Advanced
		g.AddEdge(2, 3) // Math → Advanced

		hasOrder, _, _, courseOrder := Topological(g)

		assert.True(t, hasOrder, "Course schedule should be possible")
		assert.Equal(t, 0, courseOrder[0], "Should start with Intro course")
		assert.Equal(t, 3, courseOrder[len(courseOrder)-1], "Should end with Advanced course")

		// Verify prerequisites are satisfied
		assert.True(t, indexOf(courseOrder, 0) < indexOf(courseOrder, 1), "Intro before Intermediate")
		assert.True(t, indexOf(courseOrder, 0) < indexOf(courseOrder, 2), "Intro before Math")
		assert.True(t, indexOf(courseOrder, 1) < indexOf(courseOrder, 3), "Intermediate before Advanced")
		assert.True(t, indexOf(courseOrder, 2) < indexOf(courseOrder, 3), "Math before Advanced")
	})
}

// Helper function to find index of element in slice
func indexOf(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}
