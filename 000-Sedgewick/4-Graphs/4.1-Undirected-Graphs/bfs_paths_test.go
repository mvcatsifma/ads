package undirected

import (
	"testing"


	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortestPathToVertex(t *testing.T) {
	t.Run("path to self", func(t *testing.T) {
		g := NewGraph(3)
		g.AddEdge(0, 1)

		path := ShortestPathToVertex(g, 0, 0)

		assert.Equal(t, []int{0}, path, "Path from vertex to itself should be [vertex]")
	})

	t.Run("direct connection", func(t *testing.T) {
		// Graph: 0-1-2
		g := NewGraph(3)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)

		path := ShortestPathToVertex(g, 0, 1)

		assert.Equal(t, []int{0, 1}, path, "Direct connection should give shortest 2-vertex path")
	})

	t.Run("shortest path vs longer path", func(t *testing.T) {
		// Graph: 0-1-3
		//        |   |
		//        2---4
		// Shortest 0->4: 0->2->4 (length 3)
		// Longer path: 0->1->3->4 (length 4)
		g := NewGraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 4)
		g.AddEdge(3, 4)

		path := ShortestPathToVertex(g, 0, 4)

		require.NotNil(t, path)
		assert.Equal(t, 0, path[0], "Path should start with source")
		assert.Equal(t, 4, path[len(path)-1], "Path should end with target")
		assert.Equal(t, 3, len(path), "Should find shortest path of length 3")

		// Verify it's one of the valid shortest paths
		validPaths := [][]int{{0, 2, 4}}
		assert.Contains(t, validPaths, path, "Should be the shortest path")
	})

	t.Run("no path exists - disconnected", func(t *testing.T) {
		// Graph: 0-1  2-3 (two separate components)
		g := NewGraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)

		path := ShortestPathToVertex(g, 0, 2)

		assert.Nil(t, path, "Should return nil when no path exists")
	})

	t.Run("linear chain", func(t *testing.T) {
		// Graph: 0-1-2-3-4
		g := NewGraph(5)
		for i := 0; i < 4; i++ {
			g.AddEdge(i, i+1)
		}

		path := ShortestPathToVertex(g, 0, 4)

		expected := []int{0, 1, 2, 3, 4}
		assert.Equal(t, expected, path, "Should find complete chain as shortest path")
	})

	t.Run("star graph - center to leaf", func(t *testing.T) {
		// Graph: 1-0-2 (0 is center)
		//          |
		//          3
		g := NewGraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)

		path := ShortestPathToVertex(g, 0, 2)

		assert.Equal(t, []int{0, 2}, path, "Direct path from center to leaf")
	})

	t.Run("star graph - leaf to leaf", func(t *testing.T) {
		// Graph: 1-0-2 (0 is center)
		//          |
		//          3
		g := NewGraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)

		path := ShortestPathToVertex(g, 1, 3)

		expected := []int{1, 0, 3}
		assert.Equal(t, expected, path, "Shortest path between leaves goes through center")
	})

	t.Run("cycle graph", func(t *testing.T) {
		// Graph: 0-1-2-3-0 (4-vertex cycle)
		g := NewGraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 0)

		path := ShortestPathToVertex(g, 0, 2)

		require.NotNil(t, path)
		assert.Equal(t, 0, path[0], "Path should start with source")
		assert.Equal(t, 2, path[len(path)-1], "Path should end with target")
		assert.Equal(t, 3, len(path), "Shortest path should have length 3")

		// Should be either [0,1,2] or [0,3,2]
		validPaths := [][]int{{0, 1, 2}, {0, 3, 2}}
		assert.Contains(t, validPaths, path, "Should be one of the two shortest paths")
	})

	t.Run("complete graph", func(t *testing.T) {
		// Complete graph K4: every vertex connected to every other
		g := NewGraph(4)
		for i := 0; i < 4; i++ {
			for j := i + 1; j < 4; j++ {
				g.AddEdge(i, j)
			}
		}

		path := ShortestPathToVertex(g, 0, 3)

		expected := []int{0, 3}
		assert.Equal(t, expected, path, "Direct connection should be shortest in complete graph")
	})

	t.Run("invalid source vertex - negative", func(t *testing.T) {
		g := NewGraph(3)
		g.AddEdge(0, 1)

		path := ShortestPathToVertex(g, -1, 1)

		assert.Nil(t, path, "Should return nil for negative source vertex")
	})

	t.Run("invalid source vertex - out of bounds", func(t *testing.T) {
		g := NewGraph(3)
		g.AddEdge(0, 1)

		path := ShortestPathToVertex(g, 5, 1)

		assert.Nil(t, path, "Should return nil for out of bounds source vertex")
	})

	t.Run("invalid target vertex - negative", func(t *testing.T) {
		g := NewGraph(3)
		g.AddEdge(0, 1)

		path := ShortestPathToVertex(g, 0, -1)

		assert.Nil(t, path, "Should return nil for negative target vertex")
	})

	t.Run("invalid target vertex - out of bounds", func(t *testing.T) {
		g := NewGraph(3)
		g.AddEdge(0, 1)

		path := ShortestPathToVertex(g, 0, 5)

		assert.Nil(t, path, "Should return nil for out of bounds target vertex")
	})

	t.Run("empty graph", func(t *testing.T) {
		g := NewGraph(0)

		path := ShortestPathToVertex(g, 0, 0)

		assert.Nil(t, path, "Empty graph should return nil")
	})

	t.Run("single vertex graph", func(t *testing.T) {
		g := NewGraph(1)

		path := ShortestPathToVertex(g, 0, 0)

		assert.Equal(t, []int{0}, path, "Single vertex path to self")
	})

	t.Run("path validation - all edges exist", func(t *testing.T) {
		// Create complex graph and validate returned path
		g := NewGraph(6)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 4)
		g.AddEdge(3, 5)
		g.AddEdge(4, 5)

		path := ShortestPathToVertex(g, 0, 5)

		require.NotNil(t, path, "Path should exist")
		require.True(t, len(path) >= 2, "Path should have at least 2 vertices")

		// Validate path properties
		assert.Equal(t, 0, path[0], "Path should start with source")
		assert.Equal(t, 5, path[len(path)-1], "Path should end with target")

		// Validate all consecutive vertices are connected
		for i := 0; i < len(path)-1; i++ {
			assert.True(t, isConnected(g, path[i], path[i+1]),
				"Path should only contain connected vertices: %d -> %d", path[i], path[i+1])
		}

		// Validate no duplicate vertices (simple path)
		seen := make(map[int]bool)
		for _, vertex := range path {
			assert.False(t, seen[vertex], "Path should not contain duplicate vertices: %d", vertex)
			seen[vertex] = true
		}
	})

	t.Run("BFS finds shortest path not just any path", func(t *testing.T) {
		// Graph with multiple paths of different lengths
		// 0-1-2-5
		// |     |
		// 3-4---+
		g := NewGraph(6)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 5)
		g.AddEdge(0, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 5)

		path := ShortestPathToVertex(g, 0, 5)

		require.NotNil(t, path)
		// Both paths have same length (3), so either is valid shortest path
		validShortestPaths := [][]int{{0, 1, 2, 5}, {0, 3, 4, 5}}
		assert.Contains(t, validShortestPaths, path, "Should find one of the shortest paths")
		assert.Equal(t, 4, len(path), "Shortest path should have length 4")
	})

	t.Run("large graph performance", func(t *testing.T) {
		// Create larger connected graph to test performance
		const size = 100
		g := NewGraph(size)

		// Create chain: 0-1-2-...-99
		for i := 0; i < size-1; i++ {
			g.AddEdge(i, i+1)
		}

		path := ShortestPathToVertex(g, 0, size-1)

		require.NotNil(t, path)
		assert.Equal(t, size, len(path), "Path should traverse entire chain")
		assert.Equal(t, 0, path[0], "Should start at 0")
		assert.Equal(t, size-1, path[len(path)-1], "Should end at last vertex")
	})
}
