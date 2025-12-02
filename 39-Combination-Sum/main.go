package p39

import "sort"

// combinationSum finds all unique combinations in candidates that sum to the target.
// Each candidate can be used an unlimited number of times.
// Example: combinationSum([2,3,6,7], 7) returns [[2,2,3], [7]]
//
// Approach: Uses backtracking with early pruning. Sorts candidates first to enable
// immediate termination when sum exceeds target. Passes running sum to avoid recalculation.
// Time Complexity: O(N^(T/M)) where N is candidates length, T is target, M is minimum candidate.
// Space Complexity: O(T/M) for recursion stack depth.
func combinationSum(candidates []int, target int) [][]int {
	if target == 0 {
		return [][]int{}
	}

	// Sort for early termination optimization
	sort.Ints(candidates)

	var result [][]int
	var combo []int

	// Pass running sum instead of recalculating
	var dfs func(i int, sum int)
	dfs = func(i int, sum int) {
		// Found a valid combination
		if sum == target {
			result = append(result, append([]int{}, combo...))
			return
		}

		// Prune: sum exceeds target or no more candidates
		if sum > target || i >= len(candidates) {
			return
		}

		// Include: add candidates[i] again (reuse)
		combo = append(combo, candidates[i])
		dfs(i, sum+candidates[i]) // Pass updated sum

		// Exclude: try next candidate (backtrack)
		combo = combo[:len(combo)-1]
		dfs(i+1, sum)
	}

	dfs(0, 0)
	return result
}
