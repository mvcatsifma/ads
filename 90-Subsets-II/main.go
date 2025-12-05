package p90

import "slices"

// subsetsWithDup generates all possible subsets (the power set) of the input array, without duplicate subsets.
// Handles duplicate elements in the input by skipping over consecutive duplicates when excluding an element.
// Example: subsetsWithDup([1,2,2]) returns [[], [1], [1,2], [1,2,2], [2], [2,2]]
//
// Approach: Uses backtracking with sorting to handle duplicates efficiently.
// After including an element, explores all possibilities. When excluding an element, skips all
// consecutive duplicates to avoid generating duplicate subsets.
// Time Complexity: O(n * 2^n) where n is the length of nums (2^n subsets, O(n) to copy each).
// Space Complexity: O(n) for recursion stack + O(n * 2^n) for storing all subsets.
func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	var cur []int

	slices.Sort(nums)

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(nums) {
			result = append(result, append([]int{}, cur...))
			return
		}

		cur = append(cur, nums[i])
		backtrack(i + 1)
		cur = cur[:len(cur)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		backtrack(i + 1)
	}
	backtrack(0)
	return result
}
