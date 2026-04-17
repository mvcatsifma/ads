package quicksort

// QuickSort sorts an array using the quicksort algorithm.
//
// Algorithm: Divide-and-conquer approach that selects a pivot element, partitions
// the array so all elements less than the pivot are on its left and all greater
// elements are on its right, then recursively sorts the left and right subarrays.
//
// Performance characteristics:
// - Time: O(N lg N) - probabilistic guarantee (average case)
// - Space: O(lg N) - recursion stack
// - Stable: No
// - In-place: Yes
// - Notes: Fastest general-purpose sort in practice
//
// Use Quick Sort when:
// - Average case performance is priority (fastest in practice)
// - Space is limited (in-place, no auxiliary array)
// - Stability is not required
// - Input can be shuffled/randomized to avoid worst case
func QuickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

// QuickSort3Way sorts an array using 3-way quicksort (optimized for duplicates).
//
// Algorithm: Variant of quicksort that partitions array into THREE parts instead of two:
// elements less than pivot, elements equal to pivot, and elements greater than pivot.
// This is Dijkstra's 3-way partitioning scheme, also known as the Dutch National Flag problem.
//
// Performance characteristics:
// - Time: Between N and N lg N - depends on duplicate keys
// - Space: O(lg N) - recursion stack
// - Stable: No
// - In-place: Yes
// - Notes: Linear time when many duplicates (O(N) for all equal)
//
// Use 3-way Quick Sort when:
// - Many duplicate keys are expected
// - Same benefits as regular quicksort
// - Data has low cardinality (few distinct values)
func QuickSort3Way(arr []int) {
	sort3Way(arr, 0, len(arr)-1)
}

// sort3Way recursively sorts arr[lo:hi+1] using 3-way partitioning
func sort3Way(arr []int, lo int, hi int) {
	// Base case: subarray with 0 or 1 elements is already sorted
	if hi <= lo {
		return
	}

	// Maintain invariant: arr[lo:lt] < v, arr[lt:gt] = v, arr[gt:hi] > v
	// Three pointers:
	//   lt: leftmost position of elements equal to v
	//   i:  current scanning position
	//   gt: rightmost position of elements equal to v
	lt, i, gt := lo, lo+1, hi
	v := arr[lo] // Choose pivot element

	// Single pass partitioning into three regions
	for i <= gt {
		if arr[i] < v {
			// Element belongs in left region (< v)
			arr[lt], arr[i] = arr[i], arr[lt]
			i++  // Advance scan pointer (we've seen arr[lt] before, it equals v)
			lt++ // Grow less-than region
		} else if arr[i] > v {
			// Element belongs in right region (> v)
			arr[gt], arr[i] = arr[i], arr[gt]
			gt-- // Shrink greater-than region
			// Don't increment i! We haven't examined the element we just swapped from arr[gt]
		} else {
			// Element equals pivot, already in correct region
			i++ // Just advance scan pointer
		}
	}

	// After partitioning: arr[lo:lt] < v, arr[lt:gt+1] = v, arr[gt+1:hi+1] > v
	// Recursively sort the less-than and greater-than regions
	// Skip the equal region - it's already in final position!
	sort3Way(arr, lo, lt-1)  // Sort left region (< v)
	sort3Way(arr, gt+1, hi)  // Sort right region (> v)
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
	sort(arr, lo, j-1) // Sort left subarray (elements < pivot)
	sort(arr, j+1, hi) // Sort right subarray (elements > pivot)
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
