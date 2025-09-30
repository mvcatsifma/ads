package p153

import "math"

// findMin finds the minimum element in a rotated sorted array using binary search.
// The array was originally sorted in ascending order, then rotated at some pivot.
// Time complexity: O(log n), Space complexity: O(1)
//
// Example:
//   [4,5,6,7,0,1,2] -> 0 (rotated at index 4)
//   [2,3,4,5,1] -> 1 (rotated at index 4)
//   [1,2,3,4,5] -> 1 (no rotation)
func findMin(nums []int) int {
	res := nums[0]
	l, r := 0, len(nums)-1

	for l <= r {
		if nums[l] < nums[r] {
			res = int(math.Min(float64(res), float64(nums[l])))
			break
		}

		m := (l + r) / 2
		res = int(math.Min(float64(res), float64(nums[m])))
		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}
