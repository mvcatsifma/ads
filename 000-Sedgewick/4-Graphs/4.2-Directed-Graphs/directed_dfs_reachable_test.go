package directed

import (
	"testing"


	"github.com/stretchr/testify/assert"
)

func TestReachableVerticesDirected(t *testing.T) {
	t.Run("single vertex graph", func(t *testing.T) {
		g := NewDigraph(1)

		marked, count := ReachableVerticesDirected(g, 0)

		assert.Equal(t, []bool{true}, marked, "Single vertex should be reachable from itself")
		assert.Equal(t, 1, count, "Count should be 1")
	})

	t.Run("disconnected graph", func(t *testing.T) {
		// Graph: 0  1-2  3
		//        (isolated vertices)
		g := NewDigraph(4)
		g.AddEdge(1, 2)

		marked, count := ReachableVerticesDirected(g, 0)

		expected := []bool{true, false, false, false}
		assert.Equal(t, expected, marked, "Only vertex 0 should be reachable from 0")
		assert.Equal(t, 1, count, "Count should be 1")
	})

	t.Run("connected component", func(t *testing.T) {
		// Graph: 0-1-2
		//          |
		//          3
		g := NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)

		marked, count := ReachableVerticesDirected(g, 0)

		expected := []bool{true, true, true, true}
		assert.Equal(t, expected, marked, "All vertices should be reachable from 0")
		assert.Equal(t, 4, count, "Count should be 4")
	})

	t.Run("partially connected graph", func(t *testing.T) {
		// Graph: 0-1  2-3-4
		//        (two separate components)
		g := NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)

		marked, count := ReachableVerticesDirected(g, 0)

		expected := []bool{true, true, false, false, false}
		assert.Equal(t, expected, marked, "Only vertices 0,1 should be reachable from 0")
		assert.Equal(t, 2, count, "Count should be 2")

		// Test from other component
		marked2, count2 := ReachableVerticesDirected(g, 2)
		expected2 := []bool{false, false, true, true, true}
		assert.Equal(t, expected2, marked2, "Vertices 2,3,4 should be reachable from 2")
		assert.Equal(t, 3, count2, "Count should be 3")
	})

	t.Run("cyclic graph", func(t *testing.T) {
		// Graph: 0-1
		//        |/|
		//        2-3 (cycle)
		g := NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3)

		marked, count := ReachableVerticesDirected(g, 0)

		expected := []bool{true, true, true, true}
		assert.Equal(t, expected, marked, "All vertices should be reachable in connected graph with cycle")
		assert.Equal(t, 4, count, "Count should be 4")
	})

	t.Run("linear chain", func(t *testing.T) {
		// Graph: 0→1→2→3→4 (directed linear chain)
		g := NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)

		// From start: should reach all vertices in the chain
		marked, count := ReachableVerticesDirected(g, 0)
		expected := []bool{true, true, true, true, true}
		assert.Equal(t, expected, marked, "All vertices should be reachable from chain start")
		assert.Equal(t, 5, count, "Should reach all 5 vertices")

		// From middle: should reach remaining vertices in chain
		marked2, count2 := ReachableVerticesDirected(g, 2)
		expected2 := []bool{false, false, true, true, true}
		assert.Equal(t, expected2, marked2, "Vertices 2,3,4 should be reachable from vertex 2")
		assert.Equal(t, 3, count2, "Count should be 3") // Fixed: 3, not 5
	})

	t.Run("star graph", func(t *testing.T) {
		// Graph: 1-0-2 (0 is center, connected to all others)
		//          |
		//          3
		g := NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)

		marked, count := ReachableVerticesDirected(g, 0)

		expected := []bool{true, true, true, true}
		assert.Equal(t, expected, marked, "All vertices should be reachable from center")
		assert.Equal(t, 4, count, "Count should be 4")

		// Test from leaf
		marked2, count2 := ReachableVerticesDirected(g, 1)
		expected2 := []bool{false, true, false, false}
		assert.Equal(t, expected2, marked2, "Vertex 1 should be reachable from leaf in undirected graph")
		assert.Equal(t, 1, count2, "Count should be 1")
	})

	t.Run("invalid source vertex - negative", func(t *testing.T) {
		g := NewDigraph(3)
		g.AddEdge(0, 1)

		marked, count := ReachableVerticesDirected(g, -1)

		expected := []bool{false, false, false}
		assert.Equal(t, expected, marked, "Should return all false for invalid negative source")
		assert.Equal(t, 0, count, "Count should be 0 for invalid source")
	})

	t.Run("invalid source vertex - out of bounds", func(t *testing.T) {
		g := NewDigraph(3)
		g.AddEdge(0, 1)

		marked, count := ReachableVerticesDirected(g, 5)

		expected := []bool{false, false, false}
		assert.Equal(t, expected, marked, "Should return all false for out of bounds source")
		assert.Equal(t, 0, count, "Count should be 0 for invalid source")
	})

	t.Run("empty graph", func(t *testing.T) {
		g := NewDigraph(0)

		marked, count := ReachableVerticesDirected(g, 0)

		assert.Empty(t, marked, "Should return empty array for empty graph")
		assert.Equal(t, 0, count, "Count should be 0 for empty graph")
	})

	t.Run("large connected graph", func(t *testing.T) {
		// Create a larger graph to test performance
		const size = 100
		g := NewDigraph(size)

		// Create a connected graph (chain)
		for i := 0; i < size-1; i++ {
			g.AddEdge(i, i+1)
		}

		marked, count := ReachableVerticesDirected(g, 0)

		// All vertices should be reachable
		for i := 0; i < size; i++ {
			assert.True(t, marked[i], "Vertex %d should be reachable", i)
		}
		assert.Equal(t, size, count, "Count should equal graph size")
	})
}
