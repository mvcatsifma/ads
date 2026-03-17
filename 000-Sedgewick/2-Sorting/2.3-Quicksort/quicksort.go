package quicksort

// QuickSort sorts an array using the quicksort algorithm.
//
// Algorithm: Divide-and-conquer approach that selects a pivot element, partitions
// the array so all elements less than the pivot are on its left and all greater
// elements are on its right, then recursively sorts the left and right subarrays.
//
// Time complexity: O(n log n) average case, O(n²) worst case (already sorted array)
// Space complexity: O(log n) - recursion stack depth
//
// Best for: General-purpose sorting with in-place partitioning. Fastest in practice
// due to cache efficiency and low constant factors. Use 3-way variant for data with
// many duplicates. Randomize input or use median-of-3 to avoid worst case.
func QuickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

// QuickSort3Way sorts an array using 3-way quicksort (optimized for duplicates).
// Time complexity: O(n log n) average case, O(n) when all elements are equal
// Space complexity: O(log n) - recursion stack
func QuickSort3Way(arr []int) {
	// TODO: Implement 3-way quicksort
}

// sort recursively sorts arr[lo:hi+1] using quicksort partitioning
func sort(arr []int, lo int, hi int) {
	// Base case: subarray with 0 or 1 elements is already sorted
	if hi <= lo {
		return
	}

	// Partition: rearrange so arr[lo:j] < arr[j] <= arr[j+1:hi]
	j := partition(arr, lo, hi)

	// Recursively sort left and right subarrays (pivot is already in place)
	sort(arr, lo, j-1)   // Sort left subarray (elements < pivot)
	sort(arr, j+1, hi)   // Sort right subarray (elements > pivot)
}

// partition rearranges arr[lo:hi+1] so that arr[lo:j] < arr[j] <= arr[j+1:hi]
// Returns the final position of the pivot element
func partition(arr []int, lo int, hi int) int {
	// Initialize scan pointers: i scans left→right, j scans right→left
	i, j := lo, hi+1

	// Choose pivot element (using first element - could be randomized)
	v := arr[lo]

	for {
		// Scan from left: find element >= pivot to swap
		i++
		for arr[i] < v {
			if i == hi { // Reached end, stop scanning
				break
			}
			i++
		}

		// Scan from right: find element <= pivot to swap
		j--
		for v < arr[j] {
			if j == lo { // Reached start, stop scanning
				break
			}
			j--
		}

		// Check if pointers have crossed
		if i >= j {
			break
		}

		// Swap elements that are out of place
		arr[i], arr[j] = arr[j], arr[i]
	}

	// Place pivot in its final sorted position
	arr[lo], arr[j] = arr[j], arr[lo]

	return j // Return pivot's final position
}
