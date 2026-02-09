package search

import (
	"testing"

	"leetcode/000-Sedgewick/4-Graphs/graphs"

	"github.com/stretchr/testify/assert"
)

func TestKosarajuSharirSCC(t *testing.T) {

	t.Run("invalid vertices", func(t *testing.T) {
		g := graphs.NewDigraph(3)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)

		// Test negative vertices
		assert.False(t, KosarajuSharirSCC(g, -1, 0), "Negative vertex should return false")
		assert.False(t, KosarajuSharirSCC(g, 0, -1), "Negative vertex should return false")

		// Test out of bounds vertices
		assert.False(t, KosarajuSharirSCC(g, 3, 0), "Out of bounds vertex should return false")
		assert.False(t, KosarajuSharirSCC(g, 0, 3), "Out of bounds vertex should return false")
		assert.False(t, KosarajuSharirSCC(g, 5, 10), "Both out of bounds should return false")
	})

	t.Run("empty graph", func(t *testing.T) {
		g := graphs.NewDigraph(0)

		// No vertices to test, but function should handle gracefully
		assert.False(t, KosarajuSharirSCC(g, 0, 1), "Empty graph should return false for any vertices")
	})

	t.Run("single vertex", func(t *testing.T) {
		g := graphs.NewDigraph(1)

		// Single vertex is in SCC with itself
		assert.True(t, KosarajuSharirSCC(g, 0, 0), "Single vertex should be in SCC with itself")
	})

	t.Run("single vertex with self loop", func(t *testing.T) {
		g := graphs.NewDigraph(1)
		g.AddEdge(0, 0)

		// Self loop doesn't change SCC membership
		assert.True(t, KosarajuSharirSCC(g, 0, 0), "Vertex with self loop should be in SCC with itself")
	})

	t.Run("two isolated vertices", func(t *testing.T) {
		g := graphs.NewDigraph(2)

		// No edges - each vertex is its own SCC
		assert.True(t, KosarajuSharirSCC(g, 0, 0), "Vertex 0 should be in SCC with itself")
		assert.True(t, KosarajuSharirSCC(g, 1, 1), "Vertex 1 should be in SCC with itself")
		assert.False(t, KosarajuSharirSCC(g, 0, 1), "Isolated vertices should not be in same SCC")
	})

	t.Run("simple cycle", func(t *testing.T) {
		// Graph: 0→1→2→0 (one SCC)
		g := graphs.NewDigraph(3)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 0)

		// All vertices in cycle should be in same SCC
		assert.True(t, KosarajuSharirSCC(g, 0, 1), "0 and 1 should be in same SCC")
		assert.True(t, KosarajuSharirSCC(g, 1, 2), "1 and 2 should be in same SCC")
		assert.True(t, KosarajuSharirSCC(g, 2, 0), "2 and 0 should be in same SCC")
		assert.True(t, KosarajuSharirSCC(g, 0, 2), "0 and 2 should be in same SCC")
	})

	t.Run("linear chain no cycles", func(t *testing.T) {
		// Graph: 0→1→2→3 (no cycles, each vertex is own SCC)
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)

		// Each vertex should be in its own SCC
		assert.True(t, KosarajuSharirSCC(g, 0, 0), "0 should be in SCC with itself")
		assert.True(t, KosarajuSharirSCC(g, 1, 1), "1 should be in SCC with itself")
		assert.True(t, KosarajuSharirSCC(g, 2, 2), "2 should be in SCC with itself")
		assert.True(t, KosarajuSharirSCC(g, 3, 3), "3 should be in SCC with itself")

		// No vertex should be in SCC with any other
		assert.False(t, KosarajuSharirSCC(g, 0, 1), "0 and 1 should not be in same SCC")
		assert.False(t, KosarajuSharirSCC(g, 1, 2), "1 and 2 should not be in same SCC")
		assert.False(t, KosarajuSharirSCC(g, 2, 3), "2 and 3 should not be in same SCC")
		assert.False(t, KosarajuSharirSCC(g, 0, 3), "0 and 3 should not be in same SCC")
	})

	t.Run("two separate cycles", func(t *testing.T) {
		// Graph: 0→1→0 and 2→3→2 (two separate SCCs)
		g := graphs.NewDigraph(4)
		g.AddEdge(0, 1)
		g.AddEdge(1, 0)
		g.AddEdge(2, 3)
		g.AddEdge(3, 2)

		// Within each cycle
		assert.True(t, KosarajuSharirSCC(g, 0, 1), "0 and 1 should be in same SCC")
		assert.True(t, KosarajuSharirSCC(g, 1, 0), "1 and 0 should be in same SCC")
		assert.True(t, KosarajuSharirSCC(g, 2, 3), "2 and 3 should be in same SCC")
		assert.True(t, KosarajuSharirSCC(g, 3, 2), "3 and 2 should be in same SCC")

		// Between cycles
		assert.False(t, KosarajuSharirSCC(g, 0, 2), "0 and 2 should not be in same SCC")
		assert.False(t, KosarajuSharirSCC(g, 0, 3), "0 and 3 should not be in same SCC")
		assert.False(t, KosarajuSharirSCC(g, 1, 2), "1 and 2 should not be in same SCC")
		assert.False(t, KosarajuSharirSCC(g, 1, 3), "1 and 3 should not be in same SCC")
	})

	t.Run("complex graph with multiple SCCs", func(t *testing.T) {
		// Graph: 0→1→2→0 (SCC1), 3→4 (SCC2: 3), (SCC3: 4), 1→3 (connection between SCCs)
		g := graphs.NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 0)
		g.AddEdge(1, 3)
		g.AddEdge(3, 4)

		// SCC1: {0, 1, 2}
		assert.True(t, KosarajuSharirSCC(g, 0, 1), "0 and 1 should be in same SCC")
		assert.True(t, KosarajuSharirSCC(g, 1, 2), "1 and 2 should be in same SCC")
		assert.True(t, KosarajuSharirSCC(g, 0, 2), "0 and 2 should be in same SCC")

		// SCC2: {3}, SCC3: {4}
		assert.True(t, KosarajuSharirSCC(g, 3, 3), "3 should be in SCC with itself")
		assert.True(t, KosarajuSharirSCC(g, 4, 4), "4 should be in SCC with itself")

		// Between different SCCs
		assert.False(t, KosarajuSharirSCC(g, 0, 3), "0 and 3 should not be in same SCC")
		assert.False(t, KosarajuSharirSCC(g, 1, 4), "1 and 4 should not be in same SCC")
		assert.False(t, KosarajuSharirSCC(g, 2, 3), "2 and 3 should not be in same SCC")
		assert.False(t, KosarajuSharirSCC(g, 3, 4), "3 and 4 should not be in same SCC")
	})

	t.Run("larger cycle", func(t *testing.T) {
		// Graph: 0→1→2→3→4→0 (one large SCC)
		g := graphs.NewDigraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 0)

		// All pairs should be in same SCC
		vertices := []int{0, 1, 2, 3, 4}
		for i := 0; i < len(vertices); i++ {
			for j := 0; j < len(vertices); j++ {
				assert.True(t, KosarajuSharirSCC(g, vertices[i], vertices[j]),
					"Vertices %d and %d should be in same SCC", vertices[i], vertices[j])
			}
		}
	})

	t.Run("web structure", func(t *testing.T) {
		// Complex web: 0→1→2→0 (SCC1), 2→3→4→3 (SCC2), 4→5 (SCC3: 5)
		g := graphs.NewDigraph(6)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 0) // SCC1: {0,1,2}
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 3) // SCC2: {3,4}
		g.AddEdge(4, 5) // SCC3: {5}

		// SCC1: {0,1,2}
		assert.True(t, KosarajuSharirSCC(g, 0, 1), "0,1 in same SCC")
		assert.True(t, KosarajuSharirSCC(g, 1, 2), "1,2 in same SCC")
		assert.True(t, KosarajuSharirSCC(g, 0, 2), "0,2 in same SCC")

		// SCC2: {3,4}
		assert.True(t, KosarajuSharirSCC(g, 3, 4), "3,4 in same SCC")

		// SCC3: {5}
		assert.True(t, KosarajuSharirSCC(g, 5, 5), "5 in SCC with itself")

		// Between different SCCs
		assert.False(t, KosarajuSharirSCC(g, 0, 3), "0,3 in different SCCs")
		assert.False(t, KosarajuSharirSCC(g, 1, 4), "1,4 in different SCCs")
		assert.False(t, KosarajuSharirSCC(g, 2, 5), "2,5 in different SCCs")
		assert.False(t, KosarajuSharirSCC(g, 3, 5), "3,5 in different SCCs")
		assert.False(t, KosarajuSharirSCC(g, 4, 5), "4,5 in different SCCs")
	})

	t.Run("strongly connected complete graph", func(t *testing.T) {
		// Complete graph: every vertex connects to every other vertex
		g := graphs.NewDigraph(4)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if i != j {
					g.AddEdge(i, j)
				}
			}
		}

		// All vertices should be in same SCC
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				assert.True(t, KosarajuSharirSCC(g, i, j),
					"Vertices %d and %d should be in same SCC in complete graph", i, j)
			}
		}
	})

	t.Run("real world example - web pages", func(t *testing.T) {
		// Web pages with links:
		// Page 0 (Home) → Page 1 (About) → Page 2 (Contact) → Page 0 (circular navigation)
		// Page 1 → Page 3 (Blog) → Page 4 (Archive) (one-way links)
		g := graphs.NewDigraph(5)
		g.AddEdge(0, 1) // Home → About
		g.AddEdge(1, 2) // About → Contact
		g.AddEdge(2, 0) // Contact → Home (circular navigation)
		g.AddEdge(1, 3) // About → Blog
		g.AddEdge(3, 4) // Blog → Archive

		// Main navigation cycle: {Home, About, Contact}
		assert.True(t, KosarajuSharirSCC(g, 0, 1), "Home and About are mutually reachable")
		assert.True(t, KosarajuSharirSCC(g, 1, 2), "About and Contact are mutually reachable")
		assert.True(t, KosarajuSharirSCC(g, 0, 2), "Home and Contact are mutually reachable")

		// Blog and Archive are separate components
		assert.False(t, KosarajuSharirSCC(g, 1, 3), "About and Blog are not mutually reachable")
		assert.False(t, KosarajuSharirSCC(g, 3, 4), "Blog and Archive are not mutually reachable")
		assert.False(t, KosarajuSharirSCC(g, 0, 4), "Home and Archive are not mutually reachable")
	})
}
