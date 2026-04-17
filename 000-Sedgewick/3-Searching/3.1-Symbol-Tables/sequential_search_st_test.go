package symboltables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSequentialSearchST(t *testing.T) {
	st := NewSequentialSearchST()
	assert.NotNil(t, st)
	assert.True(t, st.isEmpty())
	assert.Equal(t, 0, st.size())
}

func TestSequentialSearchST_PutAndGet(t *testing.T) {
	st := NewSequentialSearchST()

	// Put single key-value pair
	st.put("A", 1)
	value, found := st.get("A")
	assert.True(t, found)
	assert.Equal(t, 1, value)
	assert.Equal(t, 1, st.size())
	assert.False(t, st.isEmpty())
}

func TestSequentialSearchST_GetNonExistent(t *testing.T) {
	st := NewSequentialSearchST()

	// Get from empty table
	_, found := st.get("A")
	assert.False(t, found)

	// Get non-existent key from non-empty table
	st.put("B", 2)
	_, found = st.get("A")
	assert.False(t, found)
}

func TestSequentialSearchST_UpdateExistingKey(t *testing.T) {
	st := NewSequentialSearchST()

	// Insert key
	st.put("A", 1)
	assert.Equal(t, 1, st.size())

	// Update same key
	st.put("A", 10)
	assert.Equal(t, 1, st.size(), "Size should not increase when updating")

	value, found := st.get("A")
	assert.True(t, found)
	assert.Equal(t, 10, value)
}

func TestSequentialSearchST_MultipleKeys(t *testing.T) {
	st := NewSequentialSearchST()

	// Insert multiple keys
	st.put("A", 1)
	st.put("B", 2)
	st.put("C", 3)
	st.put("D", 4)

	assert.Equal(t, 4, st.size())

	// Verify all keys
	value, _ := st.get("A")
	assert.Equal(t, 1, value)
	value, _ = st.get("B")
	assert.Equal(t, 2, value)
	value, _ = st.get("C")
	assert.Equal(t, 3, value)
	value, _ = st.get("D")
	assert.Equal(t, 4, value)
}

func TestSequentialSearchST_Contains(t *testing.T) {
	st := NewSequentialSearchST()

	assert.False(t, st.contains("A"))

	st.put("A", 1)
	assert.True(t, st.contains("A"))
	assert.False(t, st.contains("B"))

	st.put("B", 2)
	assert.True(t, st.contains("A"))
	assert.True(t, st.contains("B"))
}

func TestSequentialSearchST_Delete(t *testing.T) {
	st := NewSequentialSearchST()

	// Delete from empty table (should not panic)
	st.delete("A")
	assert.Equal(t, 0, st.size())

	// Add and delete single key
	st.put("A", 1)
	assert.Equal(t, 1, st.size())
	st.delete("A")
	assert.Equal(t, 0, st.size())
	assert.False(t, st.contains("A"))

	// Delete from multiple keys
	st.put("A", 1)
	st.put("B", 2)
	st.put("C", 3)
	assert.Equal(t, 3, st.size())

	st.delete("B")
	assert.Equal(t, 2, st.size())
	assert.True(t, st.contains("A"))
	assert.False(t, st.contains("B"))
	assert.True(t, st.contains("C"))

	// Delete non-existent key
	st.delete("Z")
	assert.Equal(t, 2, st.size())
}

func TestSequentialSearchST_DeleteByPuttingZero(t *testing.T) {
	st := NewSequentialSearchST()

	st.put("A", 1)
	st.put("B", 2)
	assert.Equal(t, 2, st.size())

	// Putting zero value should delete the key
	st.put("A", 0)
	assert.Equal(t, 1, st.size())
	assert.False(t, st.contains("A"))
	assert.True(t, st.contains("B"))
}

func TestSequentialSearchST_Keys(t *testing.T) {
	st := NewSequentialSearchST()

	// Empty table
	keys := st.keys()
	assert.Equal(t, 0, len(keys))

	// Single key
	st.put("A", 1)
	keys = st.keys()
	assert.Equal(t, 1, len(keys))
	assert.Contains(t, keys, "A")

	// Multiple keys
	st.put("B", 2)
	st.put("C", 3)
	keys = st.keys()
	assert.Equal(t, 3, len(keys))
	assert.Contains(t, keys, "A")
	assert.Contains(t, keys, "B")
	assert.Contains(t, keys, "C")
}

func TestSequentialSearchST_IsEmpty(t *testing.T) {
	st := NewSequentialSearchST()

	assert.True(t, st.isEmpty())

	st.put("A", 1)
	assert.False(t, st.isEmpty())

	st.delete("A")
	assert.True(t, st.isEmpty())
}

func TestSequentialSearchST_Size(t *testing.T) {
	st := NewSequentialSearchST()

	assert.Equal(t, 0, st.size())

	st.put("A", 1)
	assert.Equal(t, 1, st.size())

	st.put("B", 2)
	assert.Equal(t, 2, st.size())

	st.put("A", 10) // Update existing
	assert.Equal(t, 2, st.size())

	st.delete("A")
	assert.Equal(t, 1, st.size())

	st.delete("B")
	assert.Equal(t, 0, st.size())
}

func TestSequentialSearchST_LargerTable(t *testing.T) {
	st := NewSequentialSearchST()

	// Insert many keys (values start at 1 to avoid zero-value deletion)
	for i := 0; i < 26; i++ {
		key := string(rune('A' + i))
		st.put(key, i+1)
	}

	assert.Equal(t, 26, st.size())

	// Verify all keys present
	for i := 0; i < 26; i++ {
		key := string(rune('A' + i))
		value, found := st.get(key)
		assert.True(t, found)
		assert.Equal(t, i+1, value)
	}

	// Delete some keys
	st.delete("M")
	st.delete("Z")
	assert.Equal(t, 24, st.size())
	assert.False(t, st.contains("M"))
	assert.False(t, st.contains("Z"))
}

func TestSequentialSearchST_OverwriteSequence(t *testing.T) {
	st := NewSequentialSearchST()

	// Multiple overwrites of same key
	st.put("A", 1)
	st.put("A", 2)
	st.put("A", 3)
	st.put("A", 4)

	assert.Equal(t, 1, st.size())
	value, found := st.get("A")
	assert.True(t, found)
	assert.Equal(t, 4, value)
}
