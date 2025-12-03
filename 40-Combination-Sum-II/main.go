package p40

import (
	"slices"
)

func combinationSum2(candidates []int, target int) [][]int {
	if target == 0 {
		return [][]int{}
	}
	slices.Sort(candidates)
	var result [][]int
	var backtrack func(cur []int, pos int, target int)
	backtrack = func(cur []int, pos int, target int) {
		if target == 0 {
			result = append(result, append([]int{}, cur...))
		}
		if target <= 0 {
			return
		}

		prev := -1
		for i := pos; i < len(candidates); i++ {
			if candidates[i] == prev {
				continue
			}
			cur = append(cur, candidates[i])
			backtrack(cur, i+1, target-candidates[i])
			cur = cur[:len(cur)-1]
			prev = candidates[i]
		}
	}
	backtrack([]int{}, 0, target)

	return result
}
