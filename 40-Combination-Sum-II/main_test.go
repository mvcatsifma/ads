package p40

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum2(t *testing.T) {
	t.Run("example 1: duplicates in candidates", func(t *testing.T) {
		result := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
		expected := [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("example 2: all same candidates", func(t *testing.T) {
		result := combinationSum2([]int{2, 5, 2, 1, 2}, 5)
		expected := [][]int{{1, 2, 2}, {5}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("no valid combinations", func(t *testing.T) {
		result := combinationSum2([]int{1, 2, 3}, 10)
		expected := [][]int{}

		assert.Equal(t, len(expected), len(result))
	})

	t.Run("target equals sum of all candidates", func(t *testing.T) {
		result := combinationSum2([]int{1, 2, 3}, 6)
		expected := [][]int{{1, 2, 3}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("single element equals target", func(t *testing.T) {
		result := combinationSum2([]int{1, 5, 1, 1}, 1)
		expected := [][]int{{1}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("multiple duplicates in combination", func(t *testing.T) {
		result := combinationSum2([]int{1, 1, 1, 1, 1}, 3)
		expected := [][]int{{1, 1, 1}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("no duplicates in candidates", func(t *testing.T) {
		result := combinationSum2([]int{2, 3, 5, 7}, 10)
		expected := [][]int{{2, 3, 5}, {3, 7}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("all combinations use duplicates", func(t *testing.T) {
		result := combinationSum2([]int{1, 1, 2, 2}, 4)
		expected := [][]int{{1, 1, 2}, {2, 2}}

		assert.Equal(t, len(expected), len(result))
		assert.Equal(t, sortCombinations(expected), sortCombinations(result))
	})

	t.Run("verify no duplicate combinations returned", func(t *testing.T) {
		result := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)

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

	t.Run("verify all combinations sum to target", func(t *testing.T) {
		result := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)

		for _, combo := range result {
			sum := 0
			for _, num := range combo {
				sum += num
			}
			assert.Equal(t, 8, sum, "Combination should sum to target")
		}
	})

	t.Run("each candidate used at most once", func(t *testing.T) {
		candidates := []int{1, 2, 2, 3}
		result := combinationSum2(candidates, 5)

		// Verify each combination uses candidates correctly
		for _, combo := range result {
			// Create a copy of candidates to track usage
			remaining := append([]int{}, candidates...)

			// Try to match each element in combo with a candidate
			for _, val := range combo {
				found := false
				for i, cand := range remaining {
					if cand == val {
						remaining = append(remaining[:i], remaining[i+1:]...)
						found = true
						break
					}
				}
				assert.True(t, found, "Combination uses candidate not available: %v from %v", combo, candidates)
			}
		}
	})

	t.Run("target is 0", func(t *testing.T) {
		result := combinationSum2([]int{1, 2, 3}, 0)
		expected := [][]int{}

		assert.Equal(t, len(expected), len(result))
	})

	t.Run("large input with all 1s", func(t *testing.T) {
		candidates := make([]int, 100)
		for i := range candidates {
			candidates[i] = 1
		}
		target := 30

		result := combinationSum2(candidates, target)

		// There's only one way to sum to 30 using 1s
		assert.Equal(t, 1, len(result), "There should be only one combination")

		if len(result) > 0 {
			assert.Equal(t, 30, len(result[0]), "The combination should have 30 ones")

			sum := 0
			for _, num := range result[0] {
				assert.Equal(t, 1, num, "All numbers in the combination should be 1")
				sum += num
			}
			assert.Equal(t, target, sum, "The combination should sum to the target")
		}
	})
}

// Helper function to sort combinations for order-independent comparison
func sortCombinations(combinations [][]int) [][]int {
	// Sort each individual combination
	for _, combo := range combinations {
		sort.Ints(combo)
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
