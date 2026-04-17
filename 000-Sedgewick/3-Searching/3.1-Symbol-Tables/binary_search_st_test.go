package symboltables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBinarySearchST(t *testing.T) {
	st := NewBinarySearchST(10)
	assert.NotNil(t, st)
	assert.True(t, st.isEmpty())
	assert.Equal(t, 0, st.size())
}

func TestBinarySearchST_PutAndGet(t *testing.T) {
	st := NewBinarySearchST(10)

	st.put("D", 4)
	value, found := st.get("D")
	assert.True(t, found)
	assert.Equal(t, 4, value)
	assert.Equal(t, 1, st.size())
}

func TestBinarySearchST_GetNonExistent(t *testing.T) {
	st := NewBinarySearchST(10)

	_, found := st.get("A")
	assert.False(t, found)

	st.put("B", 2)
	_, found = st.get("A")
	assert.False(t, found)
}

func TestBinarySearchST_UpdateExistingKey(t *testing.T) {
	st := NewBinarySearchST(10)

	st.put("D", 4)
	assert.Equal(t, 1, st.size())

	st.put("D", 40)
	assert.Equal(t, 1, st.size())

	value, found := st.get("D")
	assert.True(t, found)
	assert.Equal(t, 40, value)
}

func TestBinarySearchST_MultipleKeysInOrder(t *testing.T) {
	st := NewBinarySearchST(10)

	// Insert keys in random order
	st.put("D", 4)
	st.put("B", 2)
	st.put("F", 6)
	st.put("A", 1)
	st.put("E", 5)
	st.put("C", 3)

	assert.Equal(t, 6, st.size())

	// Verify all keys
	value, _ := st.get("A")
	assert.Equal(t, 1, value)
	value, _ = st.get("B")
	assert.Equal(t, 2, value)
	value, _ = st.get("C")
	assert.Equal(t, 3, value)
	value, _ = st.get("D")
	assert.Equal(t, 4, value)
	value, _ = st.get("E")
	assert.Equal(t, 5, value)
	value, _ = st.get("F")
	assert.Equal(t, 6, value)
}

func TestBinarySearchST_Contains(t *testing.T) {
	st := NewBinarySearchST(10)

	assert.False(t, st.contains("A"))

	st.put("A", 1)
	assert.True(t, st.contains("A"))
	assert.False(t, st.contains("B"))
}

func TestBinarySearchST_Delete(t *testing.T) {
	st := NewBinarySearchST(10)

	st.delete("A") // Delete from empty (should not panic)

	st.put("D", 4)
	st.put("B", 2)
	st.put("F", 6)
	assert.Equal(t, 3, st.size())

	st.delete("B")
	assert.Equal(t, 2, st.size())
	assert.False(t, st.contains("B"))
	assert.True(t, st.contains("D"))
	assert.True(t, st.contains("F"))

	st.delete("Z") // Delete non-existent
	assert.Equal(t, 2, st.size())
}

func TestBinarySearchST_DeleteByPuttingZero(t *testing.T) {
	st := NewBinarySearchST(10)

	st.put("A", 1)
	st.put("B", 2)
	assert.Equal(t, 2, st.size())

	st.put("A", 0)
	assert.Equal(t, 1, st.size())
	assert.False(t, st.contains("A"))
	assert.True(t, st.contains("B"))
}

func TestBinarySearchST_MinMax(t *testing.T) {
	st := NewBinarySearchST(10)

	// Empty table
	_, found := st.min()
	assert.False(t, found)
	_, found = st.max()
	assert.False(t, found)

	// Single element
	st.put("D", 4)
	min, found := st.min()
	assert.True(t, found)
	assert.Equal(t, "D", min)
	max, found := st.max()
	assert.True(t, found)
	assert.Equal(t, "D", max)

	// Multiple elements
	st.put("B", 2)
	st.put("F", 6)
	st.put("A", 1)

	min, _ = st.min()
	assert.Equal(t, "A", min)
	max, _ = st.max()
	assert.Equal(t, "F", max)
}

func TestBinarySearchST_FloorCeiling(t *testing.T) {
	st := NewBinarySearchST(10)

	st.put("B", 2)
	st.put("D", 4)
	st.put("F", 6)

	// Floor tests
	floor, found := st.floor("A")
	assert.False(t, found) // No key <= "A"

	floor, found = st.floor("C")
	assert.True(t, found)
	assert.Equal(t, "B", floor)

	floor, found = st.floor("D")
	assert.True(t, found)
	assert.Equal(t, "D", floor)

	floor, found = st.floor("G")
	assert.True(t, found)
	assert.Equal(t, "F", floor)

	// Ceiling tests
	ceiling, found := st.ceiling("A")
	assert.True(t, found)
	assert.Equal(t, "B", ceiling)

	ceiling, found = st.ceiling("C")
	assert.True(t, found)
	assert.Equal(t, "D", ceiling)

	ceiling, found = st.ceiling("D")
	assert.True(t, found)
	assert.Equal(t, "D", ceiling)

	ceiling, found = st.ceiling("G")
	assert.False(t, found) // No key >= "G"
}

func TestBinarySearchST_Rank(t *testing.T) {
	st := NewBinarySearchST(10)

	st.put("B", 2)
	st.put("D", 4)
	st.put("F", 6)
	st.put("H", 8)

	// Keys in order: B, D, F, H
	assert.Equal(t, 0, st.rank("A")) // < all keys
	assert.Equal(t, 0, st.rank("B")) // 0 keys < "B"
	assert.Equal(t, 1, st.rank("C")) // 1 key < "C"
	assert.Equal(t, 1, st.rank("D")) // 1 key < "D"
	assert.Equal(t, 2, st.rank("E")) // 2 keys < "E"
	assert.Equal(t, 4, st.rank("Z")) // all 4 keys < "Z"
}

func TestBinarySearchST_Select(t *testing.T) {
	st := NewBinarySearchST(10)

	st.put("D", 4)
	st.put("B", 2)
	st.put("F", 6)
	st.put("A", 1)

	// Keys in order: A, B, D, F
	key, found := st.selectKey(0)
	assert.True(t, found)
	assert.Equal(t, "A", key)

	key, found = st.selectKey(1)
	assert.True(t, found)
	assert.Equal(t, "B", key)

	key, found = st.selectKey(2)
	assert.True(t, found)
	assert.Equal(t, "D", key)

	key, found = st.selectKey(3)
	assert.True(t, found)
	assert.Equal(t, "F", key)

	_, found = st.selectKey(4)
	assert.False(t, found) // Out of bounds
}

func TestBinarySearchST_DeleteMinMax(t *testing.T) {
	st := NewBinarySearchST(10)

	st.put("D", 4)
	st.put("B", 2)
	st.put("F", 6)
	st.put("A", 1)
	st.put("E", 5)

	// Delete min
	st.deleteMin()
	assert.Equal(t, 4, st.size())
	assert.False(t, st.contains("A"))
	min, _ := st.min()
	assert.Equal(t, "B", min)

	// Delete max
	st.deleteMax()
	assert.Equal(t, 3, st.size())
	assert.False(t, st.contains("F"))
	max, _ := st.max()
	assert.Equal(t, "E", max)
}

func TestBinarySearchST_Keys(t *testing.T) {
	st := NewBinarySearchST(10)

	// Empty table
	keys := st.keys()
	assert.Equal(t, 0, len(keys))

	// Multiple keys should be in sorted order
	st.put("D", 4)
	st.put("B", 2)
	st.put("F", 6)
	st.put("A", 1)

	keys = st.keys()
	assert.Equal(t, 4, len(keys))
	assert.Equal(t, []string{"A", "B", "D", "F"}, keys)
}

func TestBinarySearchST_KeysRange(t *testing.T) {
	st := NewBinarySearchST(10)

	st.put("B", 2)
	st.put("D", 4)
	st.put("F", 6)
	st.put("H", 8)
	st.put("J", 10)

	// Range queries
	keys := st.keysRange("A", "K")
	assert.Equal(t, []string{"B", "D", "F", "H", "J"}, keys)

	keys = st.keysRange("C", "G")
	assert.Equal(t, []string{"D", "F"}, keys)

	keys = st.keysRange("D", "F")
	assert.Equal(t, []string{"D", "F"}, keys)

	keys = st.keysRange("E", "G")
	assert.Equal(t, []string{"F"}, keys)

	keys = st.keysRange("K", "Z")
	assert.Equal(t, 0, len(keys))
}

func TestBinarySearchST_IsEmpty(t *testing.T) {
	st := NewBinarySearchST(10)

	assert.True(t, st.isEmpty())

	st.put("A", 1)
	assert.False(t, st.isEmpty())

	st.delete("A")
	assert.True(t, st.isEmpty())
}

func TestBinarySearchST_Size(t *testing.T) {
	st := NewBinarySearchST(10)

	assert.Equal(t, 0, st.size())

	st.put("A", 1)
	assert.Equal(t, 1, st.size())

	st.put("B", 2)
	assert.Equal(t, 2, st.size())

	st.put("A", 10) // Update
	assert.Equal(t, 2, st.size())

	st.delete("A")
	assert.Equal(t, 1, st.size())
}
