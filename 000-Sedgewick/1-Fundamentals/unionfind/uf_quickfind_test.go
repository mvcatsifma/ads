package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickFind(t *testing.T) {

	t.Run("initialization", func(t *testing.T) {
		uf := NewQuickFind(5)

		assert.Equal(t, 5, uf.count, "Should have 5 components initially")
		assert.Len(t, uf.id, 5, "ID array should have 5 elements")

		// Each element should be in its own component
		for i := 0; i < 5; i++ {
			assert.Equal(t, i, uf.id[i], "Element %d should be in component %d", i, i)
		}
	})

	t.Run("find returns component id", func(t *testing.T) {
		uf := NewQuickFind(5)

		// Initially, each element is its own component
		for i := 0; i < 5; i++ {
			assert.Equal(t, i, uf.find(i), "find(%d) should return %d", i, i)
		}
	})

	t.Run("connected initially false", func(t *testing.T) {
		uf := NewQuickFind(5)

		// No elements should be connected initially
		assert.False(t, uf.connected(0, 1), "0 and 1 should not be connected")
		assert.False(t, uf.connected(1, 2), "1 and 2 should not be connected")
		assert.False(t, uf.connected(0, 4), "0 and 4 should not be connected")

		// Element should be connected to itself
		for i := 0; i < 5; i++ {
			assert.True(t, uf.connected(i, i), "Element %d should be connected to itself", i)
		}
	})

	t.Run("union connects two elements", func(t *testing.T) {
		uf := NewQuickFind(5)

		uf.union(0, 1)

		assert.True(t, uf.connected(0, 1), "0 and 1 should be connected after union")
		assert.True(t, uf.connected(1, 0), "Connection should be symmetric")
		assert.Equal(t, 4, uf.count, "Should have 4 components after one union")
	})

	t.Run("union is transitive", func(t *testing.T) {
		uf := NewQuickFind(5)

		uf.union(0, 1)
		uf.union(1, 2)

		// All three should be connected
		assert.True(t, uf.connected(0, 1), "0 and 1 should be connected")
		assert.True(t, uf.connected(1, 2), "1 and 2 should be connected")
		assert.True(t, uf.connected(0, 2), "0 and 2 should be connected (transitive)")
		assert.Equal(t, 3, uf.count, "Should have 3 components")
	})

	t.Run("union already connected elements", func(t *testing.T) {
		uf := NewQuickFind(5)

		uf.union(0, 1)
		initialCount := uf.count

		uf.union(0, 1) // Union same elements again

		assert.Equal(t, initialCount, uf.count, "Count should not change when unioning already connected elements")
		assert.True(t, uf.connected(0, 1), "0 and 1 should still be connected")
	})

	t.Run("multiple unions create larger component", func(t *testing.T) {
		uf := NewQuickFind(10)

		// Create component {0,1,2,3,4}
		uf.union(0, 1)
		uf.union(1, 2)
		uf.union(2, 3)
		uf.union(3, 4)

		// All should be connected
		assert.True(t, uf.connected(0, 4), "0 and 4 should be connected")
		assert.True(t, uf.connected(1, 3), "1 and 3 should be connected")
		assert.True(t, uf.connected(0, 2), "0 and 2 should be connected")

		// Should not be connected to elements outside component
		assert.False(t, uf.connected(0, 5), "0 and 5 should not be connected")
		assert.False(t, uf.connected(4, 9), "4 and 9 should not be connected")

		assert.Equal(t, 6, uf.count, "Should have 6 components (one with 5 elements, five with 1 element)")
	})

	t.Run("separate components", func(t *testing.T) {
		uf := NewQuickFind(8)

		// Create component {0,1,2}
		uf.union(0, 1)
		uf.union(1, 2)

		// Create component {3,4,5}
		uf.union(3, 4)
		uf.union(4, 5)

		// Within components
		assert.True(t, uf.connected(0, 2), "0 and 2 should be connected")
		assert.True(t, uf.connected(3, 5), "3 and 5 should be connected")

		// Between components
		assert.False(t, uf.connected(0, 3), "0 and 3 should not be connected")
		assert.False(t, uf.connected(2, 4), "2 and 4 should not be connected")
		assert.False(t, uf.connected(1, 5), "1 and 5 should not be connected")

		// Isolated elements
		assert.False(t, uf.connected(0, 6), "0 and 6 should not be connected")
		assert.False(t, uf.connected(3, 7), "3 and 7 should not be connected")

		assert.Equal(t, 4, uf.count, "Should have 4 components")
	})

	t.Run("merge two components", func(t *testing.T) {
		uf := NewQuickFind(6)

		// Create component {0,1,2}
		uf.union(0, 1)
		uf.union(1, 2)

		// Create component {3,4,5}
		uf.union(3, 4)
		uf.union(4, 5)

		assert.Equal(t, 2, uf.count, "Should have 2 components")

		// Merge the two components
		uf.union(2, 3)

		assert.Equal(t, 1, uf.count, "Should have 1 component after merge")

		// All elements should be connected
		assert.True(t, uf.connected(0, 5), "0 and 5 should be connected")
		assert.True(t, uf.connected(1, 4), "1 and 4 should be connected")
		assert.True(t, uf.connected(2, 3), "2 and 3 should be connected")
	})

	t.Run("all elements in one component", func(t *testing.T) {
		uf := NewQuickFind(5)

		// Connect all elements
		uf.union(0, 1)
		uf.union(1, 2)
		uf.union(2, 3)
		uf.union(3, 4)

		assert.Equal(t, 1, uf.count, "Should have 1 component")

		// All pairs should be connected
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				assert.True(t, uf.connected(i, j), "Elements %d and %d should be connected", i, j)
			}
		}
	})

	t.Run("count decreases correctly", func(t *testing.T) {
		uf := NewQuickFind(10)

		assert.Equal(t, 10, uf.count, "Should start with 10 components")

		uf.union(0, 1)
		assert.Equal(t, 9, uf.count, "Should have 9 components after first union")

		uf.union(2, 3)
		assert.Equal(t, 8, uf.count, "Should have 8 components")

		uf.union(0, 2) // Merge {0,1} with {2,3}
		assert.Equal(t, 7, uf.count, "Should have 7 components")

		uf.union(1, 3) // Already connected, should not change count
		assert.Equal(t, 7, uf.count, "Count should not change for already connected elements")
	})

	t.Run("single element", func(t *testing.T) {
		uf := NewQuickFind(1)

		assert.Equal(t, 1, uf.count, "Should have 1 component")
		assert.True(t, uf.connected(0, 0), "Element should be connected to itself")
		assert.Equal(t, 0, uf.find(0), "find(0) should return 0")
	})
}
