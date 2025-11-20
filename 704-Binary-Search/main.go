package p704

// search performs a binary search on the sorted slice nums to find the index
// of the target value. If the target is not found, it returns -1.
//
// Time complexity: O(log n), where n is the length of the input slice nums.
// Space complexity: O(1), as the function uses only a constant amount of extra space.
//
// Parameters:
//   - nums: a sorted slice of integers to search
//   - target: the value to search for in the slice
//
// Returns:
//   - int: the index of the target value in the slice, or -1 if the target is not found
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}

	return -1
}
