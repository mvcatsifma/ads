package p000Backtracking

// permute generates all possible permutations of the input array.
// A permutation is a unique arrangement of all elements.
// Example: permute([1,2,3]) returns [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
//
// Approach: Uses backtracking to explore all possible orderings.
// For each position, tries each unused number, then backtracks to try other numbers.
// Time Complexity: O(n! * n) where n is the length of nums (n! permutations, n to copy each).
// Space Complexity: O(n) for recursion stack depth + O(n! * n) for storing all permutations.
func permute(nums []int) [][]int {
	var result [][]int
	var current []int
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			result = append(result, append([]int{}, current...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			current = append(current, nums[i])
			used[i] = true

			backtrack()

			current = current[:len(current)-1]
			used[i] = false
		}
	}
	backtrack()
	return result
}
