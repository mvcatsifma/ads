package p88

// merge combines two sorted integer arrays into one sorted array in-place using
// optimal backward merging. Assumes nums1 has sufficient capacity (m+n) to hold all elements.
//
// Parameters:
//   - nums1: first sorted array with extra space at the end
//   - m: number of valid elements in nums1
//   - nums2: second sorted array to merge
//   - n: number of elements in nums2
//
// The function merges from right to left, comparing the largest unprocessed elements
// and placing them at the end of nums1. This avoids overwriting unprocessed elements.
// Modifies nums1 in-place.
//
// Time Complexity: O(m+n)
// Space Complexity: O(1)
//
// Example: merge([1,2,3,0,0,0], 3, [2,5,6], 3) → nums1 becomes [1,2,2,3,5,6]
func merge(nums1 []int, m int, nums2 []int, n int) {
	last := m + n - 1

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[last] = nums1[m-1]
			m--
		} else {
			nums1[last] = nums2[n-1]
			n--
		}

		last--
	}

	for n > 0 {
		nums1[last] = nums2[n-1]
		n--
		last--
	}
}
