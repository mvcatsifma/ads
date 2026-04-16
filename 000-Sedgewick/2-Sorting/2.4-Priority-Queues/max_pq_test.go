package priorityqueus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMaxPQ(t *testing.T) {
	pq := NewMaxPQ(10)
	assert.NotNil(t, pq)
	assert.Equal(t, 0, pq.size())
	assert.True(t, pq.isEmpty())
	assert.Equal(t, 11, len(pq.pq)) // maxN+1 for 1-indexed array
}

func TestMaxPQ_Insert(t *testing.T) {
	pq := NewMaxPQ(10)

	pq.insert("A")
	assert.Equal(t, 1, pq.size())
	assert.False(t, pq.isEmpty())
	assert.Equal(t, Key("A"), pq.pq[1])
}

func TestMaxPQ_IsEmpty(t *testing.T) {
	pq := NewMaxPQ(5)
	assert.True(t, pq.isEmpty())

	pq.insert("A")
	assert.False(t, pq.isEmpty())
}

func TestMaxPQ_Size(t *testing.T) {
	pq := NewMaxPQ(10)
	assert.Equal(t, 0, pq.size())

	pq.insert("A")
	assert.Equal(t, 1, pq.size())

	pq.insert("B")
	assert.Equal(t, 2, pq.size())

	pq.insert("C")
	assert.Equal(t, 3, pq.size())
}

func TestMaxPQ_Less(t *testing.T) {
	pq := NewMaxPQ(10)
	// Manually set values to test less() without heap reordering
	pq.pq[1] = "A"
	pq.pq[2] = "B"
	pq.pq[3] = "C"
	pq.n = 3

	// "A" < "B" < "C" lexicographically
	assert.True(t, pq.less(1, 2))   // "A" < "B"
	assert.True(t, pq.less(2, 3))   // "B" < "C"
	assert.False(t, pq.less(2, 1))  // "B" not < "A"
}

func TestMaxPQ_Exch(t *testing.T) {
	pq := NewMaxPQ(10)
	// Manually set values to test exch() without heap reordering
	pq.pq[1] = "A"
	pq.pq[2] = "B"
	pq.n = 2

	assert.Equal(t, Key("A"), pq.pq[1])
	assert.Equal(t, Key("B"), pq.pq[2])

	pq.exch(1, 2)

	assert.Equal(t, Key("B"), pq.pq[1])
	assert.Equal(t, Key("A"), pq.pq[2])
}

func TestMaxPQ_Swim(t *testing.T) {
	pq := NewMaxPQ(10)

	// Manually construct heap with element that needs to swim up
	pq.pq[1] = "B"
	pq.pq[2] = "A"
	pq.pq[3] = "C"
	pq.n = 3

	// "C" at position 3 should swim up past "A" at position 1
	pq.swim(3)

	// After swimming, "C" should be at root since it's largest
	assert.Equal(t, Key("C"), pq.pq[1])
}

func TestMaxPQ_Sink(t *testing.T) {
	pq := NewMaxPQ(10)

	// Manually construct heap with element that needs to sink down
	pq.pq[1] = "A" // Small element at root
	pq.pq[2] = "C" // Larger children
	pq.pq[3] = "B"
	pq.n = 3

	// "A" at position 1 should sink down
	pq.sink(1)

	// After sinking, "C" should be at root (largest of children)
	assert.Equal(t, Key("C"), pq.pq[1])
	assert.Equal(t, Key("A"), pq.pq[2])
}

func TestMaxPQ_MultipleInserts(t *testing.T) {
	pq := NewMaxPQ(10)

	elements := []Key{"D", "B", "F", "A", "E", "C"}
	for _, elem := range elements {
		pq.insert(elem)
	}

	assert.Equal(t, 6, pq.size())
	assert.False(t, pq.isEmpty())
}

func TestMaxPQ_EmptyPQ(t *testing.T) {
	pq := NewMaxPQ(5)

	assert.True(t, pq.isEmpty())
	assert.Equal(t, 0, pq.size())
}

func TestMaxPQ_SingleElement(t *testing.T) {
	pq := NewMaxPQ(5)

	pq.insert("X")

	assert.False(t, pq.isEmpty())
	assert.Equal(t, 1, pq.size())
	assert.Equal(t, Key("X"), pq.pq[1])
}

func TestMaxPQ_HeapProperty(t *testing.T) {
	pq := NewMaxPQ(10)

	// Insert elements in random order - insert will automatically maintain heap property
	elements := []Key{"D", "B", "F", "A", "E", "C"}
	for _, elem := range elements {
		pq.insert(elem)
	}

	// Verify heap property: parent >= both children
	// For max heap with 1-indexed array: pq[k] >= pq[2*k] and pq[k] >= pq[2*k+1]
	for k := 1; k <= pq.n/2; k++ {
		// Check left child
		if 2*k <= pq.n {
			assert.False(t, pq.less(k, 2*k), "Heap property violated: parent at %d (%s) < left child at %d (%s)", k, pq.pq[k], 2*k, pq.pq[2*k])
		}
		// Check right child
		if 2*k+1 <= pq.n {
			assert.False(t, pq.less(k, 2*k+1), "Heap property violated: parent at %d (%s) < right child at %d (%s)", k, pq.pq[k], 2*k+1, pq.pq[2*k+1])
		}
	}

	// Verify the max element is at the root
	assert.Equal(t, Key("F"), pq.pq[1], "Maximum element should be at root")
}
