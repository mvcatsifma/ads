package p88

// merge combines two sorted integer arrays into one sorted array in-place.
// Assumes nums1 has sufficient capacity (m+n) to hold all elements.
//
// Parameters:
//   - nums1: first sorted array with extra space at the end
//   - m: number of valid elements in nums1
//   - nums2: second sorted array to merge
//   - n: number of elements in nums2
//
// The function appends nums2 elements to nums1 and sorts the result.
// Modifies nums1 in-place.
//
// Example: merge([1,2,3,0,0,0], 3, [2,5,6], 3) → nums1 becomes [1,2,2,3,5,6]
func merge(nums1 []int, m int, nums2 []int, n int) {
	var j = 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[j]
		j++
	}

	bubbleSort(nums1)
}

func bubbleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// Early termination if no swaps occurred
		if !swapped {
			break
		}
	}
}
