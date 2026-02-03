package search

import (
	"testing"

	"leetcode/000-Sedgewick/4-Graphs/graphs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPathToVertexDirected(t *testing.T) {
	t.Run("path to self", func(t *testing.T) {
		g := graphs.NewDigraph(3)
		g.AddEdge(0, 1)

		path := PathToVertexDirected(g, 0, 0)

		assert.Equal(t, []int{0}, path, "Path from vertex to itself should be [vertex]")
	})

	t.Run("direct connection", func(t *testing.T) {
		// Graph: 0-1-2
		g := graphs.NewDigraph(3)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)

		path := PathToVertexDirected(g, 0, 1)

		assert.Equal(t, []int{0, 1}, path, "Direct connection should give 2-vertex path")
	})

	t.Run("multi-hop path", func(t *testing.T) {
		// Graph: 0-1-2-3
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)

		path := PathToVertexDirected(g, 0, 3)

		expected := []int{0, 1, 2, 3}
		assert.Equal(t, expected, path, "Should find path through intermediate vertices")
	})

	t.Run("no path exists - disconnected", func(t *testing.T) {
		// Graph: 0-1  2-3 (two separate components)
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)

		path := PathToVertexDirected(g, 0, 2)

		assert.Nil(t, path, "Should return nil when no path exists")
	})

	t.Run("path in cyclic graph", func(t *testing.T) {
		// Graph: 0-1
		//        |/|
		//        2-3 (cycle)
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 3)

		path := PathToVertexDirected(g, 0, 3)

		require.NotNil(t, path, "Should find a path in cyclic graph")
		assert.Equal(t, 0, path[0], "Path should start with source")
		assert.Equal(t, 3, path[len(path)-1], "Path should end with target")

		// Verify path is valid (each consecutive pair is connected)
		for i := 0; i < len(path)-1; i++ {
			assert.True(t, isConnectedDirected(g, path[i], path[i+1]),
				"Vertices %d and %d should be connected", path[i], path[i+1])
		}
	})

	t.Run("star graph - center to leaf", func(t *testing.T) {
		// Graph: 1-0-2 (0 is center)
		//          |
		//          3
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)

		path := PathToVertexDirected(g, 0, 2)

		assert.Equal(t, []int{0, 2}, path, "Direct path from center to leaf")
	})

	t.Run("star graph - no path exists", func(t *testing.T) {
		// Graph: 1-0-2 (0 is center)
		//          |
		//          3
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)

		path := PathToVertexDirected(g, 1, 3)

		assert.Nil(t, path, "Should return nil when no path exists")
	})

	t.Run("complex graph with multiple paths", func(t *testing.T) {
		// Graph: 0-1-3
		//        |   |
		//        2---4
		g := graphs.NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 4)
		g.AddEdge(3, 4)

		path := PathToVertexDirected(g, 0, 4)

		require.NotNil(t, path, "Should find a path")
		assert.Equal(t, 0, path[0], "Path should start with source")
		assert.Equal(t, 4, path[len(path)-1], "Path should end with target")

		// Verify it's a valid path
		for i := 0; i < len(path)-1; i++ {
			assert.True(t, isConnectedDirected(g, path[i], path[i+1]),
				"Consecutive vertices should be connected")
		}
	})

	t.Run("invalid source vertex - negative", func(t *testing.T) {
		g := graphs.NewDigraph(3)
		g.AddEdge(0, 1)

		path := PathToVertexDirected(g, -1, 1)

		assert.Nil(t, path, "Should return nil for negative source vertex")
	})

	t.Run("invalid source vertex - out of bounds", func(t *testing.T) {
		g := graphs.NewDigraph(3)
		g.AddEdge(0, 1)

		path := PathToVertexDirected(g, 5, 1)

		assert.Nil(t, path, "Should return nil for out of bounds source vertex")
	})

	t.Run("invalid target vertex - negative", func(t *testing.T) {
		g := graphs.NewDigraph(3)
		g.AddEdge(0, 1)

		path := PathToVertexDirected(g, 0, -1)

		assert.Nil(t, path, "Should return nil for negative target vertex")
	})

	t.Run("invalid target vertex - out of bounds", func(t *testing.T) {
		g := graphs.NewDigraph(3)
		g.AddEdge(0, 1)

		path := PathToVertexDirected(g, 0, 5)

		assert.Nil(t, path, "Should return nil for out of bounds target vertex")
	})

	t.Run("single vertex graph", func(t *testing.T) {
		g := graphs.NewDigraph(1)

		path := PathToVertexDirected(g, 0, 0)

		assert.Equal(t, []int{0}, path, "Single vertex path to self")
	})

	t.Run("empty graph", func(t *testing.T) {
		g := graphs.NewDigraph(0)

		path := PathToVertexDirected(g, 0, 0)

		assert.Nil(t, path, "Empty graph should return nil")
	})

	t.Run("linear chain - forward", func(t *testing.T) {
		// Graph: 0-1-2-3-4
		g := graphs.NewDigraph(5)
		for i := 0; i < 4; i++ {
			g.AddEdge(i, i+1)
		}

		path := PathToVertexDirected(g, 0, 4)

		expected := []int{0, 1, 2, 3, 4}
		assert.Equal(t, expected, path, "Should find complete chain path")
	})

	t.Run("linear chain - no path exists", func(t *testing.T) {
		// Graph: 0-1-2-3-4
		g := graphs.NewDigraph(5)
		for i := 0; i < 4; i++ {
			g.AddEdge(i, i+1)
		}

		path := PathToVertexDirected(g, 4, 0)

		assert.Nil(t, path, "No path exists")
	})

	t.Run("path validation", func(t *testing.T) {
		// Create a more complex graph and validate path properties
		g := graphs.NewDigraph(6)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 4)
		g.AddEdge(3, 5)
		g.AddEdge(4, 5)

		path := PathToVertexDirected(g, 0, 5)

		require.NotNil(t, path, "Path should exist")
		require.True(t, len(path) >= 2, "Path should have at least 2 vertices")

		// Validate path properties
		assert.Equal(t, 0, path[0], "Path should start with source")
		assert.Equal(t, 5, path[len(path)-1], "Path should end with target")

		// Validate all consecutive vertices are connected
		for i := 0; i < len(path)-1; i++ {
			assert.True(t, isConnectedDirected(g, path[i], path[i+1]),
				"Path should only contain connected vertices: %d -> %d", path[i], path[i+1])
		}

		// Validate no duplicate vertices (simple path)
		seen := make(map[int]bool)
		for _, vertex := range path {
			assert.False(t, seen[vertex], "Path should not contain duplicate vertices: %d", vertex)
			seen[vertex] = true
		}
	})
}

// Helper function to check if two vertices are connected
func isConnectedDirected(g *graphs.Digraph, v, w int) bool {
	for _, neighbor := range g.Adjacent(v) {
		if neighbor == w {
			return true
		}
	}
	return false
}
