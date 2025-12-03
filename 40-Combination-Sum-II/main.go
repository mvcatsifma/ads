package p40

import (
	"encoding/json"
	"slices"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	if target == 0 {
		return [][]int{}
	}

	// Sort for early termination optimization
	sort.Ints(candidates)

	var combo []int
	seen := make(map[string]bool)

	// Pass running sum instead of recalculating
	var dfs func(i int, sum int)
	dfs = func(i int, sum int) {
		// Found a valid combination
		if sum == target {
			slices.Sort(combo)
			key := sliceToKey(combo)
			if !seen[key] {
				seen[key] = true
			}
			return
		}

		// Prune: sum exceeds target or no more candidates
		if sum > target || i >= len(candidates) {
			return
		}

		// Include: add candidates[i] again (reuse)
		combo = append(combo, candidates[i])
		dfs(i+1, sum+candidates[i]) // Pass updated sum

		// Exclude: try next candidate (backtrack)
		combo = combo[:len(combo)-1]
		dfs(i+1, sum)
	}

	dfs(0, 0)

	var result [][]int
	for k := range seen {
		result = append(result, keyToSlice(k))
	}

	return result
}

func sliceToKey(slice []int) string {
	bytes, _ := json.Marshal(slice)
	return string(bytes)
}

func keyToSlice(key string) []int {
	var slice []int
	_ = json.Unmarshal([]byte(key), &slice)
	return slice
}
