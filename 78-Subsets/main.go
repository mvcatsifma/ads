package p78

// subsets generates all possible subsets (the power set) of the input array.
// A subset is any combination of elements from the array, including the empty set and the full array.
// Example: subsets([1,2,3]) returns [[], [1], [2], [3], [1,2], [1,3], [2,3], [1,2,3]] (in any order)
//
// Approach: Uses backtracking to explore two choices at each element: include it or exclude it.
// For each element at index i, recursively builds subsets with the element included,
// then backtracks and builds subsets with the element excluded.
// Time Complexity: O(n * 2^n) where n is the length of nums (2^n subsets, O(n) to copy each).
// Space Complexity: O(n) for recursion stack depth + O(n * 2^n) for storing all subsets.
func subsets(nums []int) [][]int {
	var result [][]int
	var subset []int

	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(nums) {
			result = append(result, append([]int{}, subset...))
			return
		}

		subset = append(subset, nums[i])
		dfs(i + 1)

		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}
	dfs(0)
	return result
}
