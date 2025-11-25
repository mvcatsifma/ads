package p163

// findMissingRanges finds all missing ranges between lower and upper bounds in a sorted array.
// Returns a 2D array where each element is [start, end] representing a continuous missing range.
// Example: nums=[1,3,4,6], lower=0, upper=7 returns [[0,0], [2,2], [5,5], [7,7]]
//
// Approach: Check gaps before first element, between consecutive elements, and after last element.
// Time Complexity: O(n) where n is the length of nums (single pass through array).
// Space Complexity: O(1) excluding output array (only uses constant extra variables).
func findMissingRanges(nums []int, lower int, upper int) [][]int {
	n := len(nums)
	missingRanges := make([][]int, 0)
	if n == 0 {
		missingRanges = append(missingRanges, []int{lower, upper})
		return missingRanges
	}

	if lower < nums[0] {
		missingRanges = append(missingRanges, []int{lower, nums[0] - 1})
	}

	for i := 0; i < n-1; i++ {
		if nums[i+1]-nums[i] <= 1 {
			continue
		}
		missingRanges = append(missingRanges, []int{nums[i] + 1, nums[i+1] - 1})
	}

	if upper > nums[n-1] {
		missingRanges = append(missingRanges, []int{nums[n-1] + 1, upper})
	}

	return missingRanges
}
