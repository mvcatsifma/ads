package priorityqueus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIndexMinPQ(t *testing.T) {
	pq := NewIndexMinPQ(10)
	assert.NotNil(t, pq)
	assert.Equal(t, 0, pq.size())
	assert.True(t, pq.isEmpty())

	// All qp entries should be -1 initially
	for i := 0; i <= 10; i++ {
		assert.Equal(t, -1, pq.qp[i])
		assert.False(t, pq.contains(i))
	}
}

func TestIndexMinPQ_Insert(t *testing.T) {
	pq := NewIndexMinPQ(10)

	pq.insert(3, "C")
	assert.Equal(t, 1, pq.size())
	assert.False(t, pq.isEmpty())
	assert.True(t, pq.contains(3))
	assert.Equal(t, Key("C"), pq.minKey())
	assert.Equal(t, 3, pq.minIndex())
}

func TestIndexMinPQ_MultipleInserts(t *testing.T) {
	pq := NewIndexMinPQ(10)

	// Insert with different client indices and keys
	pq.insert(5, "E")
	pq.insert(2, "B")
	pq.insert(7, "G")
	pq.insert(1, "A")
	pq.insert(4, "D")

	assert.Equal(t, 5, pq.size())
	assert.Equal(t, Key("A"), pq.minKey()) // "A" is minimum
	assert.Equal(t, 1, pq.minIndex())       // at client index 1
}

func TestIndexMinPQ_MinKey(t *testing.T) {
	pq := NewIndexMinPQ(10)

	pq.insert(3, "C")
	pq.insert(1, "A")
	pq.insert(2, "B")

	assert.Equal(t, Key("A"), pq.minKey())
}

func TestIndexMinPQ_MinIndex(t *testing.T) {
	pq := NewIndexMinPQ(10)

	pq.insert(3, "C")
	pq.insert(1, "A")
	pq.insert(2, "B")

	assert.Equal(t, 1, pq.minIndex())
}

func TestIndexMinPQ_DelMin(t *testing.T) {
	pq := NewIndexMinPQ(10)

	pq.insert(5, "E")
	pq.insert(2, "B")
	pq.insert(7, "G")
	pq.insert(1, "A")
	pq.insert(4, "D")

	// Delete minimum ("A" at index 1)
	minIdx := pq.delMin()
	assert.Equal(t, 1, minIdx)
	assert.Equal(t, 4, pq.size())
	assert.False(t, pq.contains(1))

	// Next minimum should be "B" at index 2
	assert.Equal(t, Key("B"), pq.minKey())
	assert.Equal(t, 2, pq.minIndex())
}

func TestIndexMinPQ_DelMinOrder(t *testing.T) {
	pq := NewIndexMinPQ(10)

	pq.insert(0, "D")
	pq.insert(1, "B")
	pq.insert(2, "F")
	pq.insert(3, "A")
	pq.insert(4, "E")
	pq.insert(5, "C")

	// Should delete in sorted order: A, B, C, D, E, F
	expected := []struct {
		index int
		key   Key
	}{
		{3, "A"},
		{1, "B"},
		{5, "C"},
		{0, "D"},
		{4, "E"},
		{2, "F"},
	}

	for _, exp := range expected {
		assert.Equal(t, exp.key, pq.minKey())
		minIdx := pq.delMin()
		assert.Equal(t, exp.index, minIdx)
	}

	assert.True(t, pq.isEmpty())
}

func TestIndexMinPQ_ChangeKey(t *testing.T) {
	pq := NewIndexMinPQ(10)

	pq.insert(0, "D")
	pq.insert(1, "B")
	pq.insert(2, "F")

	assert.Equal(t, Key("B"), pq.minKey())
	assert.Equal(t, 1, pq.minIndex())

	// Change index 1 from "B" to "Z" (now largest)
	pq.changeKey(1, "Z")
	assert.Equal(t, Key("D"), pq.minKey())
	assert.Equal(t, 0, pq.minIndex())

	// Change index 2 from "F" to "A" (now smallest)
	pq.changeKey(2, "A")
	assert.Equal(t, Key("A"), pq.minKey())
	assert.Equal(t, 2, pq.minIndex())
}

func TestIndexMinPQ_Delete(t *testing.T) {
	pq := NewIndexMinPQ(10)

	pq.insert(0, "D")
	pq.insert(1, "B")
	pq.insert(2, "F")
	pq.insert(3, "A")

	assert.Equal(t, 4, pq.size())
	assert.True(t, pq.contains(1))

	// Delete index 1 (key "B")
	pq.delete(1)
	assert.Equal(t, 3, pq.size())
	assert.False(t, pq.contains(1))

	// Minimum should still be "A" at index 3
	assert.Equal(t, Key("A"), pq.minKey())
	assert.Equal(t, 3, pq.minIndex())

	// Delete index 3 (key "A", the minimum)
	pq.delete(3)
	assert.Equal(t, 2, pq.size())
	assert.False(t, pq.contains(3))

	// Now minimum should be "D" at index 0
	assert.Equal(t, Key("D"), pq.minKey())
	assert.Equal(t, 0, pq.minIndex())
}

func TestIndexMinPQ_Contains(t *testing.T) {
	pq := NewIndexMinPQ(10)

	assert.False(t, pq.contains(0))
	assert.False(t, pq.contains(5))

	pq.insert(0, "A")
	pq.insert(5, "E")

	assert.True(t, pq.contains(0))
	assert.True(t, pq.contains(5))
	assert.False(t, pq.contains(3))

	pq.delMin()
	assert.False(t, pq.contains(0))
	assert.True(t, pq.contains(5))
}

func TestIndexMinPQ_HeapProperty(t *testing.T) {
	pq := NewIndexMinPQ(10)

	// Insert elements
	pq.insert(5, "E")
	pq.insert(2, "B")
	pq.insert(7, "G")
	pq.insert(1, "A")
	pq.insert(4, "D")
	pq.insert(6, "F")
	pq.insert(3, "C")

	// Verify min heap property: parent <= both children
	// For min heap with 1-indexed array: keys[pq[k]] <= keys[pq[2*k]] and keys[pq[k]] <= keys[pq[2*k+1]]
	for k := 1; k <= pq.n/2; k++ {
		parentKey := pq.keys[pq.pq[k]]

		// Check left child
		if 2*k <= pq.n {
			leftKey := pq.keys[pq.pq[2*k]]
			assert.False(t, parentKey > leftKey,
				"Heap property violated: parent at %d (%s) > left child at %d (%s)",
				k, parentKey, 2*k, leftKey)
		}

		// Check right child
		if 2*k+1 <= pq.n {
			rightKey := pq.keys[pq.pq[2*k+1]]
			assert.False(t, parentKey > rightKey,
				"Heap property violated: parent at %d (%s) > right child at %d (%s)",
				k, parentKey, 2*k+1, rightKey)
		}
	}

	// Verify the min element is at the root
	assert.Equal(t, Key("A"), pq.keys[pq.pq[1]], "Minimum element should be at root")
}

func TestIndexMinPQ_InverseMapping(t *testing.T) {
	pq := NewIndexMinPQ(10)

	pq.insert(3, "C")
	pq.insert(1, "A")
	pq.insert(5, "E")

	// Verify pq and qp are inverses
	for i := 1; i <= pq.n; i++ {
		clientIdx := pq.pq[i]
		heapPos := pq.qp[clientIdx]
		assert.Equal(t, i, heapPos, "pq[%d] = %d, but qp[%d] != %d", i, clientIdx, clientIdx, i)
	}

	// Verify qp[pq[i]] = i
	for i := 1; i <= pq.n; i++ {
		assert.Equal(t, i, pq.qp[pq.pq[i]], "qp[pq[%d]] should equal %d", i, i)
	}
}

func TestIndexMinPQ_IsEmpty(t *testing.T) {
	pq := NewIndexMinPQ(5)
	assert.True(t, pq.isEmpty())

	pq.insert(0, "A")
	assert.False(t, pq.isEmpty())

	pq.delMin()
	assert.True(t, pq.isEmpty())
}

func TestIndexMinPQ_Size(t *testing.T) {
	pq := NewIndexMinPQ(10)
	assert.Equal(t, 0, pq.size())

	pq.insert(0, "A")
	assert.Equal(t, 1, pq.size())

	pq.insert(1, "B")
	assert.Equal(t, 2, pq.size())

	pq.insert(2, "C")
	assert.Equal(t, 3, pq.size())

	pq.delMin()
	assert.Equal(t, 2, pq.size())

	pq.delete(1)
	assert.Equal(t, 1, pq.size())
}
