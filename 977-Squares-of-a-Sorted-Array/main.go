package p977

// sortedSquares returns the squares of a sorted array in sorted order using two pointers.
//
// Algorithm: Since input is sorted, largest absolute values are at the edges.
// Use two pointers at both ends, compare their squares, and place the larger square
// at the end of result array. Work backwards from largest to smallest.
//
// Time complexity: O(n) - single pass through array
// Space complexity: O(n) - for result array (or O(1) if not counting output)
//
// Key insight: Filling back-to-front ensures smallest value (zero if present) naturally
// ends up at position 0 without special handling.
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	// Two pointers: left starts at beginning, right at end
	left, right := 0, len(nums)-1

	// Fill result from back to front (largest to smallest)
	for k := len(nums) - 1; k >= 0; k-- {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		// Compare squares from both ends
		// Larger square goes at current position
		if leftSquare > rightSquare {
			result[k] = leftSquare
			left++
		} else {
			result[k] = rightSquare
			right--
		}
	}

	return result
}
