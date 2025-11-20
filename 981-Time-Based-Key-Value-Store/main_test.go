package p981

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeMap(t *testing.T) {
	t.Run("constructor creates empty TimeMap", func(t *testing.T) {
		tm := Constructor()

		// Should return empty string for non-existent key
		result := tm.Get("nonexistent", 1)
		assert.Equal(t, "", result)
	})

	t.Run("set and get single value", func(t *testing.T) {
		tm := Constructor()

		tm.Set("key1", "value1", 10)

		result := tm.Get("key1", 10)
		assert.Equal(t, "value1", result)
	})

	t.Run("get with exact timestamp", func(t *testing.T) {
		tm := Constructor()

		tm.Set("foo", "bar", 1)
		tm.Set("foo", "bar2", 4)

		assert.Equal(t, "bar", tm.Get("foo", 1))
		assert.Equal(t, "bar2", tm.Get("foo", 4))
	})

	t.Run("get with timestamp before any values", func(t *testing.T) {
		tm := Constructor()

		tm.Set("foo", "bar", 5)

		// Timestamp before any stored value should return empty
		result := tm.Get("foo", 3)
		assert.Equal(t, "", result)
	})

	t.Run("get with timestamp after stored values", func(t *testing.T) {
		tm := Constructor()

		tm.Set("foo", "bar", 1)
		tm.Set("foo", "bar2", 4)

		// Should return the latest value at or before timestamp
		result := tm.Get("foo", 10)
		assert.Equal(t, "bar2", result)
	})

	t.Run("get with timestamp between stored values", func(t *testing.T) {
		tm := Constructor()

		tm.Set("foo", "bar", 1)
		tm.Set("foo", "bar2", 4)
		tm.Set("foo", "bar3", 7)

		// Timestamp 3 should return value from timestamp 1
		assert.Equal(t, "bar", tm.Get("foo", 3))

		// Timestamp 5 should return value from timestamp 4
		assert.Equal(t, "bar2", tm.Get("foo", 5))
	})

	t.Run("multiple keys", func(t *testing.T) {
		tm := Constructor()

		tm.Set("love", "high", 10)
		tm.Set("love", "low", 20)
		tm.Set("hate", "very_low", 15)

		assert.Equal(t, "high", tm.Get("love", 15))
		assert.Equal(t, "low", tm.Get("love", 25))
		assert.Equal(t, "very_low", tm.Get("hate", 20))
		assert.Equal(t, "", tm.Get("peace", 10))
	})

	t.Run("overwrite same timestamp", func(t *testing.T) {
		tm := Constructor()

		tm.Set("key", "value1", 10)
		tm.Set("key", "value2", 10) // Same timestamp

		// Should get the latest value for that timestamp
		result := tm.Get("key", 10)
		assert.Equal(t, "value2", result)
	})

	t.Run("empty key", func(t *testing.T) {
		tm := Constructor()

		tm.Set("", "empty_key_value", 5)

		assert.Equal(t, "empty_key_value", tm.Get("", 5))
		assert.Equal(t, "empty_key_value", tm.Get("", 10))
	})

	t.Run("large number of operations", func(t *testing.T) {
		tm := Constructor()

		// Set many values for the same key
		for i := 0; i < 1000; i += 10 {
			tm.Set("stress_test", "value_"+string(rune(i)), i)
		}

		// Test retrieval at various timestamps
		assert.Equal(t, "value_"+string(rune(100)), tm.Get("stress_test", 105))
		assert.Equal(t, "value_"+string(rune(990)), tm.Get("stress_test", 1000))
		assert.Equal(t, "", tm.Get("stress_test", -1))
	})
}
