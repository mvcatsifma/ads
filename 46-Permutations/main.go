package p46

import "slices"

// permute generates all possible permutations of the given integer slice.
// A permutation is an arrangement of all elements in any order.
// Example: permute([1,2,3]) returns [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
//
// Approach: Uses backtracking to build permutations incrementally.
// For each permutation, tries adding each unused number, then backtracks to try other numbers.
// Time Complexity: O(n!) where n is the length of nums (n! permutations, each taking O(n) to build).
// Space Complexity: O(n) for recursion stack + O(n! * n) for storing all permutations.
func permute(nums []int) [][]int {
	var result [][]int

	// Recursive backtracking function
	var backtrack func(curr []int)
	backtrack = func(curr []int) {
		// Base case: found a complete permutation
		if len(curr) == len(nums) {
			// Create a deep copy to avoid slice reference issues
			result = append(result, append([]int{}, curr...))
			return
		}

		// Try adding each unused number to the current permutation
		for _, num := range nums {
			// Skip if number is already used in current permutation
			if !slices.Contains(curr, num) {
				// Choose: add num to current permutation
				curr = append(curr, num)

				// Explore: recursively build rest of permutation
				backtrack(curr)

				// Unchoose: remove last number (backtrack)
				curr = curr[:len(curr)-1]
			}
		}
	}

	backtrack([]int{})
	return result
}
