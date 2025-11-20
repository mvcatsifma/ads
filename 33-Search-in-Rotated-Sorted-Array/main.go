package p33

// search performs binary search on a rotated sorted array to find the target value.
// The algorithm first finds the pivot point (smallest element) where the rotation occurred,
// then performs binary search on both the left part [0, pivot-1] and right part [pivot, n-1].
// Returns the index of target if found, -1 otherwise.
//
// Example: [4,5,6,7,0,1,2] rotated at index 4
// - Left part: [4,5,6,7] (indices 0-3)
// - Right part: [0,1,2] (indices 4-6)
//
// Time Complexity: O(log n) - binary search to find pivot + binary search on each part
// Space Complexity: O(1) - only uses constant extra space for variables
func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1

	// Find the index of the pivot element (the smallest element)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[n-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	binarySearch := func(leftBoundary int, rightBoundary int, target int) int {
		left, right := leftBoundary, rightBoundary
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == target {
				return mid
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return -1
	}

	// Binary search over elements on the pivot's left
	if answer := binarySearch(0, left-1, target); answer != -1 {
		return answer
	}

	// Binary search over elements on the pivot element's right
	return binarySearch(left, n-1, target)
}
