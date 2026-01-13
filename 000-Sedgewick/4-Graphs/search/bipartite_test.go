package search

import (
	"testing"

	"leetcode/000-Sedgewick/4-Graphs/graphs"

	"github.com/stretchr/testify/assert"
)

func TestIsBipartite(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		g := graphs.NewGraph(0)

		result := isBipartite(g)

		assert.True(t, result, "Empty graph should be bipartite")
	})

	t.Run("single vertex", func(t *testing.T) {
		g := graphs.NewGraph(1)

		result := isBipartite(g)

		assert.True(t, result, "Single vertex should be bipartite")
	})

	t.Run("two vertices with edge", func(t *testing.T) {
		// Graph: 0-1
		g := graphs.NewGraph(2)
		g.AddEdge(0, 1)

		result := isBipartite(g)

		assert.True(t, result, "Two vertices connected by edge should be bipartite")
	})

	t.Run("triangle - odd cycle", func(t *testing.T) {
		// Graph: 0-1-2-0 (triangle - odd cycle)
		g := graphs.NewGraph(3)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 0)

		result := isBipartite(g)

		assert.False(t, result, "Triangle (odd cycle) should not be bipartite")
	})

	t.Run("square - even cycle", func(t *testing.T) {
		// Graph: 0-1-2-3-0 (4-cycle)
		g := graphs.NewGraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 0)

		result := isBipartite(g)

		assert.True(t, result, "Square (even cycle) should be bipartite")
	})

	t.Run("tree structure", func(t *testing.T) {
		// Graph: 0-1-2, 1-3 (tree)
		g := graphs.NewGraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)

		result := isBipartite(g)

		assert.True(t, result, "Tree should be bipartite")
	})

	t.Run("complete bipartite graph K2,2", func(t *testing.T) {
		// Graph: Complete bipartite K2,2 - {0,1} connected to {2,3}
		g := graphs.NewGraph(4)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)

		result := isBipartite(g)

		assert.True(t, result, "Complete bipartite graph should be bipartite")
	})

	t.Run("complete bipartite graph K3,2", func(t *testing.T) {
		// Graph: Complete bipartite K3,2 - {0,1,2} connected to {3,4}
		g := graphs.NewGraph(5)
		g.AddEdge(0, 3)
		g.AddEdge(0, 4)
		g.AddEdge(1, 3)
		g.AddEdge(1, 4)
		g.AddEdge(2, 3)
		g.AddEdge(2, 4)

		result := isBipartite(g)

		assert.True(t, result, "Complete bipartite K3,2 should be bipartite")
	})

	t.Run("5-cycle - odd cycle", func(t *testing.T) {
		// Graph: 0-1-2-3-4-0 (5-cycle)
		g := graphs.NewGraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 0)

		result := isBipartite(g)

		assert.False(t, result, "5-cycle (odd cycle) should not be bipartite")
	})

	t.Run("6-cycle - even cycle", func(t *testing.T) {
		// Graph: 0-1-2-3-4-5-0 (6-cycle)
		g := graphs.NewGraph(6)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 5)
		g.AddEdge(5, 0)

		result := isBipartite(g)

		assert.True(t, result, "6-cycle (even cycle) should be bipartite")
	})

	t.Run("disconnected components all bipartite", func(t *testing.T) {
		// Graph: 0-1 (edge), 2-3-4-5-2 (4-cycle), isolated vertex 6
		g := graphs.NewGraph(7)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 5)
		g.AddEdge(5, 2)
		// vertex 6 is isolated

		result := isBipartite(g)

		assert.True(t, result, "All disconnected components bipartite should result in bipartite graph")
	})

	t.Run("disconnected components one non-bipartite", func(t *testing.T) {
		// Graph: 0-1 (bipartite edge), 2-3-4-2 (triangle - non-bipartite)
		g := graphs.NewGraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 2)

		result := isBipartite(g)

		assert.False(t, result, "One non-bipartite component should make entire graph non-bipartite")
	})

	t.Run("star graph", func(t *testing.T) {
		// Graph: 0 connected to 1,2,3,4 (star with center 0)
		g := graphs.NewGraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)
		g.AddEdge(0, 4)

		result := isBipartite(g)

		assert.True(t, result, "Star graph should be bipartite")
	})

	t.Run("path graph", func(t *testing.T) {
		// Graph: 0-1-2-3-4 (linear path)
		g := graphs.NewGraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)

		result := isBipartite(g)

		assert.True(t, result, "Path graph should be bipartite")
	})

	t.Run("complete graph K4", func(t *testing.T) {
		// Graph: Complete graph with 4 vertices (contains triangles)
		g := graphs.NewGraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)
		g.AddEdge(1, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3)

		result := isBipartite(g)

		assert.False(t, result, "Complete graph K4 contains triangles and should not be bipartite")
	})
}
