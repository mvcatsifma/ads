package p40

import (
	"slices"
)

// combinationSum2 finds all unique combinations in candidates where the chosen numbers sum to target.
// Each number in candidates may only be used once in the combination.
// Example: combinationSum2([10,1,2,7,6,1,5], 8) returns [[1,1,6],[1,2,5],[1,7],[2,6]]
//
// Approach: Uses backtracking with sorting to handle duplicates efficiently.
// Time Complexity: O(2^n) where n is the number of candidates (each candidate can be included or not).
// Space Complexity: O(n) for recursion stack + O(2^n) for storing all combinations in the worst case.
func combinationSum2(candidates []int, target int) [][]int {
	// Edge case: no valid combinations if target is 0
	if target == 0 {
		return [][]int{}
	}

	// Sort candidates to group duplicates together
	slices.Sort(candidates)

	var result [][]int

	// Recursive backtracking function
	var backtrack func(cur []int, pos int, target int)
	backtrack = func(cur []int, pos int, target int) {
		// Base case: found a valid combination
		if target == 0 {
			result = append(result, append([]int{}, cur...))
		}
		// Pruning: stop if target becomes negative
		if target <= 0 {
			return
		}

		prev := -1 // Track previous candidate to skip duplicates
		for i := pos; i < len(candidates); i++ {
			// Skip duplicates at the same level of recursion
			if candidates[i] == prev {
				continue
			}

			// Choose current candidate
			cur = append(cur, candidates[i])

			// Explore next candidates (i+1 ensures each candidate is used only once)
			backtrack(cur, i+1, target-candidates[i])

			// Unchoose (backtrack)
			cur = cur[:len(cur)-1]

			// Update prev to skip duplicates in next iteration
			prev = candidates[i]
		}
	}

	backtrack([]int{}, 0, target)

	return result
}
