package directed

import (
	"testing"


	"github.com/stretchr/testify/assert"
)

func TestHasCycleDirected(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		g := NewDigraph(0)

		result := hasCycleDirected(g)

		assert.False(t, result, "Empty graph should have no cycles")
	})

	t.Run("single vertex", func(t *testing.T) {
		g := NewDigraph(1)

		result := hasCycleDirected(g)

		assert.False(t, result, "Single vertex with no edges should have no cycles")
	})

	t.Run("two vertices with edge", func(t *testing.T) {
		// Graph: 0-1
		g := NewDigraph(2)
		g.AddEdge(0, 1)

		result := hasCycleDirected(g)

		assert.False(t, result, "Simple edge between two vertices should not form cycle")
	})

	t.Run("triangle cycle", func(t *testing.T) {
		// Graph: 0-1-2-0 (triangle)
		g := NewDigraph(3)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 0)

		result := hasCycleDirected(g)

		assert.True(t, result, "Triangle should form a cycle")
	})

	t.Run("tree structure", func(t *testing.T) {
		// Graph: 0-1-2, 1-3 (tree)
		g := NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)

		result := hasCycleDirected(g)

		assert.False(t, result, "Tree structure should have no cycles")
	})

	t.Run("tree with additional edge creating cycle", func(t *testing.T) {
		// Graph: 0-1-2, 1-3, 2-3 (tree + cycle)
		g := NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3) // creates cycle 1-2-3-1

		result := hasCycleDirected(g)

		assert.False(t, result, "Adding edge to tree should create cycle")
	})

	t.Run("disconnected components no cycles", func(t *testing.T) {
		// Graph: 0-1, 2-3, 4-5 (three separate edges)
		g := NewDigraph(6)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)
		g.AddEdge(4, 5)

		result := hasCycleDirected(g)

		assert.False(t, result, "Disconnected components without cycles should return false")
	})

	t.Run("disconnected with one component having cycle", func(t *testing.T) {
		// Graph: 0-1 (no cycle), 2-3-4-2 (cycle)
		g := NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 2)

		result := hasCycleDirected(g)

		assert.True(t, result, "Should detect cycle even in disconnected components")
	})

	t.Run("large cycle", func(t *testing.T) {
		// Graph: 0-1-2-3-4-0 (5-vertex cycle)
		g := NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 0)

		result := hasCycleDirected(g)

		assert.True(t, result, "Large cycle should be detected")
	})

	t.Run("complex graph w/o cycle", func(t *testing.T) {
		// Graph: 0-1, 0-2, 1-3, 2-3, 3-4, 4-5 (cycle at 1-3-2-0-1)
		g := NewDigraph(6)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 5)

		result := hasCycleDirected(g)

		assert.False(t, result, "Complex graph with embedded cycle should be detected")
	})

	t.Run("complete graph K4", func(t *testing.T) {
		// Graph: Complete graph with 4 vertices (every vertex connected to every other)
		g := NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3)

		result := hasCycleDirected(g)

		assert.False(t, result, "Complete graph should contain cycles")
	})

	t.Run("linear chain", func(t *testing.T) {
		// Graph: 0-1-2-3-4 (linear chain, no cycle)
		g := NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)

		result := hasCycleDirected(g)

		assert.False(t, result, "Linear chain should have no cycles")
	})
}
