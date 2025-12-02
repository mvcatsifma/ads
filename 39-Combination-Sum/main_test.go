package p39

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	t.Run("example 1: target found with multiple combinations", func(t *testing.T) {
		result := combinationSum([]int{2, 3, 6, 7}, 7)
		expected := [][]int{{2, 2, 3}, {7}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("example 2: multiple ways to reach target", func(t *testing.T) {
		result := combinationSum([]int{2, 3, 5}, 8)
		expected := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("example 3: single candidate reaches target", func(t *testing.T) {
		result := combinationSum([]int{2}, 1)
		expected := [][]int{}

		assert.Equal(t, len(expected), len(result))
	})

	t.Run("target equals single candidate", func(t *testing.T) {
		result := combinationSum([]int{1, 2, 3}, 3)
		expected := [][]int{{1, 1, 1}, {1, 2}, {3}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("target is 0", func(t *testing.T) {
		result := combinationSum([]int{1, 2, 3}, 0)
		expected := [][]int{}

		assert.Equal(t, len(expected), len(result))
	})

	t.Run("single candidate with repetition", func(t *testing.T) {
		result := combinationSum([]int{5}, 10)
		expected := [][]int{{5, 5}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("large target", func(t *testing.T) {
		result := combinationSum([]int{2, 3, 7}, 18)
		// Multiple valid combinations exist

		// Verify all combinations sum to target
		for _, combo := range result {
			sum := 0
			for _, num := range combo {
				sum += num
			}
			assert.Equal(t, 18, sum, "Combination should sum to target")
		}
	})

	t.Run("candidates can be reused", func(t *testing.T) {
		result := combinationSum([]int{3}, 9)
		expected := [][]int{{3, 3, 3}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("verify no duplicate combinations", func(t *testing.T) {
		result := combinationSum([]int{2, 3, 6, 7}, 7)

		// Convert to map to check for duplicates
		seen := make(map[string]bool)
		for _, combo := range result {
			sorted := append([]int{}, combo...)
			sort.Ints(sorted)
			key := fmt.Sprintf("%v", sorted)
			assert.False(t, seen[key], "Duplicate combination found: %v", combo)
			seen[key] = true
		}
	})
}

// Helper function to sort combinations for order-independent comparison
func sortCombinations(combinations [][]int) [][]int {
	// Sort each individual combination
	for _, combo := range combinations {
		slices.Sort(combo)
	}

	// Sort the combinations themselves for consistent comparison
	sort.Slice(combinations, func(i, j int) bool {
		// Compare by length first
		if len(combinations[i]) != len(combinations[j]) {
			return len(combinations[i]) < len(combinations[j])
		}
		// Then compare element by element
		for k := 0; k < len(combinations[i]); k++ {
			if combinations[i][k] != combinations[j][k] {
				return combinations[i][k] < combinations[j][k]
			}
		}
		return false
	})

	return combinations
}
