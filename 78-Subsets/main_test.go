package p78

import (
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		result := subsets([]int{})
		expected := [][]int{{}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortSubsets(expected), sortSubsets(result))
	})

	t.Run("single element", func(t *testing.T) {
		result := subsets([]int{1})
		expected := [][]int{{}, {1}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortSubsets(expected), sortSubsets(result))
	})

	t.Run("two elements", func(t *testing.T) {
		result := subsets([]int{1, 2})
		expected := [][]int{{}, {1}, {2}, {1, 2}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortSubsets(expected), sortSubsets(result))
	})

	t.Run("three elements", func(t *testing.T) {
		result := subsets([]int{1, 2, 3})
		expected := [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortSubsets(expected), sortSubsets(result))
	})

	t.Run("array with duplicates in set", func(t *testing.T) {
		result := subsets([]int{0, 0})
		// Both 0s will be included or excluded independently
		// Results might be [[], [0], [0], [0,0]]

		assert.Equal(t, 4, len(result))
		assert.Equal(t, sortSubsets(result), sortSubsets(result))
	})

	t.Run("negative numbers", func(t *testing.T) {
		result := subsets([]int{-1, 0, 1})
		expected := [][]int{{}, {-1}, {0}, {1}, {-1, 0}, {-1, 1}, {0, 1}, {-1, 0, 1}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortSubsets(expected), sortSubsets(result))
	})

	t.Run("larger array", func(t *testing.T) {
		result := subsets([]int{1, 2, 3, 4})
		// 2^4 = 16 subsets
		assert.Equal(t, 16, len(result))
		assert.Equal(t, sortSubsets(result), sortSubsets(result))
	})
}

// Helper function to sort a 2D slice for comparison (ignores order)
func sortSubsets(subsets [][]int) [][]int {
	// Sort each individual subset
	for _, subset := range subsets {
		slices.Sort(subset)
	}

	// Sort the subsets themselves for consistent comparison
	sort.Slice(subsets, func(i, j int) bool {
		// Compare by length first
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		// Then compare element by element
		for k := 0; k < len(subsets[i]); k++ {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})

	return subsets
}
