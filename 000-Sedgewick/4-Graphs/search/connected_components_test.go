package search

import (
	"sort"
	"testing"

	"leetcode/000-Sedgewick/4-Graphs/graphs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnectedComponents(t *testing.T) {
	t.Run("single vertex", func(t *testing.T) {
		g := graphs.NewGraph(1)

		count, components := connectedComponents(g)

		assert.Equal(t, 1, count, "Single vertex should form one component")
		assert.Equal(t, [][]int{{0}}, components, "Component should contain vertex 0")
	})

	t.Run("fully connected graph", func(t *testing.T) {
		// Graph: 0-1-2 (all connected)
		g := graphs.NewGraph(3)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)

		count, components := connectedComponents(g)

		assert.Equal(t, 1, count, "Connected graph should have one component")
		require.Len(t, components, 1)

		// Sort for consistent comparison
		sort.Ints(components[0])
		assert.Equal(t, []int{0, 1, 2}, components[0], "All vertices should be in same component")
	})

	t.Run("two separate components", func(t *testing.T) {
		// Graph: 0-1  2-3 (two separate pairs)
		g := graphs.NewGraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(2, 3)

		count, components := connectedComponents(g)

		assert.Equal(t, 2, count, "Should find two separate components")
		require.Len(t, components, 2)

		// Sort components for consistent comparison
		for i := range components {
			sort.Ints(components[i])
		}
		sort.Slice(components, func(i, j int) bool {
			return components[i][0] < components[j][0]
		})

		expected := [][]int{{0, 1}, {2, 3}}
		assert.Equal(t, expected, components, "Should have two separate components")
	})

	t.Run("isolated vertices", func(t *testing.T) {
		// Graph: 0  1  2  3 (all isolated)
		g := graphs.NewGraph(4)

		count, components := connectedComponents(g)

		assert.Equal(t, 4, count, "Each isolated vertex should be its own component")
		require.Len(t, components, 4)

		// Sort components by first element
		sort.Slice(components, func(i, j int) bool {
			return components[i][0] < components[j][0]
		})

		expected := [][]int{{0}, {1}, {2}, {3}}
		assert.Equal(t, expected, components, "Each vertex should be in separate component")
	})

	t.Run("mixed components", func(t *testing.T) {
		// Graph: 0-1-2  3  4-5 (three components: size 3, 1, 2)
		g := graphs.NewGraph(6)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(4, 5)

		count, components := connectedComponents(g)

		assert.Equal(t, 3, count, "Should find three components")
		require.Len(t, components, 3)

		// Sort components for consistent comparison
		for i := range components {
			sort.Ints(components[i])
		}
		sort.Slice(components, func(i, j int) bool {
			return len(components[i]) > len(components[j]) ||
				(len(components[i]) == len(components[j]) && components[i][0] < components[j][0])
		})

		expected := [][]int{{0, 1, 2}, {4, 5}, {3}}
		assert.Equal(t, expected, components, "Should have components of sizes 3, 2, 1")
	})

	t.Run("cycle graph", func(t *testing.T) {
		// Graph: 0-1-2-3-0 (4-vertex cycle)
		g := graphs.NewGraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 0)

		count, components := connectedComponents(g)

		assert.Equal(t, 1, count, "Cycle should form one connected component")
		require.Len(t, components, 1)

		sort.Ints(components[0])
		assert.Equal(t, []int{0, 1, 2, 3}, components[0], "All vertices should be in same component")
	})

	t.Run("star graph", func(t *testing.T) {
		// Graph: 1-0-2 (0 is center connected to all others)
		//          |
		//          3
		g := graphs.NewGraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(0, 3)

		count, components := connectedComponents(g)

		assert.Equal(t, 1, count, "Star graph should be one connected component")
		require.Len(t, components, 1)

		sort.Ints(components[0])
		assert.Equal(t, []int{0, 1, 2, 3}, components[0], "All vertices should be connected through center")
	})

	t.Run("complex multi-component graph", func(t *testing.T) {
		// Graph: 0-1-2  3-4  5  6-7-8-9
		//        (4 components of sizes 3, 2, 1, 4)
		g := graphs.NewGraph(10)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(3, 4)
		g.AddEdge(6, 7)
		g.AddEdge(7, 8)
		g.AddEdge(8, 9)

		count, components := connectedComponents(g)

		assert.Equal(t, 4, count, "Should find four components")
		require.Len(t, components, 4)

		// Sort components for consistent comparison
		for i := range components {
			sort.Ints(components[i])
		}
		sort.Slice(components, func(i, j int) bool {
			return components[i][0] < components[j][0]
		})

		expected := [][]int{{0, 1, 2}, {3, 4}, {5}, {6, 7, 8, 9}}
		assert.Equal(t, expected, components, "Should correctly identify all four components")
	})

	t.Run("empty graph", func(t *testing.T) {
		g := graphs.NewGraph(0)

		count, components := connectedComponents(g)

		assert.Equal(t, 0, count, "Empty graph should have zero components")
		assert.Equal(t, [][]int{}, components, "Should return empty component list")
	})

	t.Run("nil graph", func(t *testing.T) {
		count, components := connectedComponents(nil)

		assert.Equal(t, 0, count, "Nil graph should have zero components")
		assert.Equal(t, [][]int{}, components, "Should return empty component list")
	})

	t.Run("verify component properties", func(t *testing.T) {
		// Create a graph with known structure
		g := graphs.NewGraph(7)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(3, 4)
		// Vertices 5 and 6 are isolated

		count, components := connectedComponents(g)

		// Verify basic properties
		assert.Equal(t, 4, count, "Should find correct number of components")
		require.Len(t, components, 4)

		// Verify all vertices are accounted for
		totalVertices := 0
		allVertices := make(map[int]bool)

		for _, component := range components {
			totalVertices += len(component)
			for _, vertex := range component {
				assert.False(t, allVertices[vertex], "Vertex %d should not appear in multiple components", vertex)
				allVertices[vertex] = true
				assert.True(t, vertex >= 0 && vertex < g.V, "Vertex %d should be within bounds", vertex)
			}
		}

		assert.Equal(t, g.V, totalVertices, "All vertices should be assigned to exactly one component")
		assert.Equal(t, g.V, len(allVertices), "All vertices should be accounted for")
	})

	t.Run("large graph with multiple components", func(t *testing.T) {
		// Create larger graph: chains of length 10 each
		const chainLength = 10
		const numChains = 5
		g := graphs.NewGraph(chainLength * numChains)

		// Create separate chains
		for chain := 0; chain < numChains; chain++ {
			start := chain * chainLength
			for i := 0; i < chainLength-1; i++ {
				g.AddEdge(start+i, start+i+1)
			}
		}

		count, components := connectedComponents(g)

		assert.Equal(t, numChains, count, "Should find one component per chain")
		require.Len(t, components, numChains)

		// Verify each component has correct size
		for _, component := range components {
			assert.Equal(t, chainLength, len(component), "Each chain should have %d vertices", chainLength)
		}
	})
}
