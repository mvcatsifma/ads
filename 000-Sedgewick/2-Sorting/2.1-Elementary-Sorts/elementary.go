package elementary

// SelectionSort sorts an array using the selection sort algorithm.
//
// Algorithm: Repeatedly find the minimum element from the unsorted portion
// and place it at the beginning. After each iteration, the sorted portion
// grows by one element from the left.
//
// Performance characteristics:
// - Time: O(N²) - always makes N²/2 comparisons regardless of input order
// - Space: O(1) - in-place sorting
// - Stable: No
// - In-place: Yes
// - Notes: Exactly N exchanges (minimal data movement)
//
// Use Selection Sort when:
// - Minimizing write operations is critical (only N swaps)
// - Working with hardware where writes are expensive
// - Array is very small
func SelectionSort(arr []int) {
	n := len(arr)

	// Iterate through each position in the array
	for i := 0; i < n; i++ {
		// Assume current position holds the minimum
		m := i

		// Find the actual minimum in the remaining unsorted portion
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[m] {
				m = j // Update minimum index
			}
		}

		// Swap the minimum element with the current position
		// This grows the sorted portion by one element
		arr[i], arr[m] = arr[m], arr[i]
	}
}

// InsertionSort sorts an array using the insertion sort algorithm.
//
// Algorithm: Like sorting playing cards in your hand - pick each element
// and insert it into its correct position among the already sorted elements
// by repeatedly swapping it leftward until it's in place.
//
// Performance characteristics:
// - Time: Between N and N² - depends on input order
// - Space: O(1) - in-place sorting
// - Stable: Yes
// - In-place: Yes
// - Notes: Linear time for nearly sorted data
//
// Use Insertion Sort when:
// - Array size < 10-15 elements
// - Array is nearly sorted (best case: O(N))
// - Stability is required and array is small
// - As optimization for recursive sorts on small subarrays
func InsertionSort(arr []int) {
	n := len(arr)

	// Start from second element (first element is trivially sorted)
	for i := 1; i < n; i++ {
		// Insert arr[i] into the sorted portion arr[0:i]
		// Move leftward while current element is smaller than its left neighbor
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			// Swap current element with its left neighbor
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// ShellSort sorts an array using the shell sort algorithm.
//
// Algorithm: An optimized version of insertion sort that allows exchanges of
// elements that are far apart. It works by comparing elements separated by a gap
// (h-sorting), then progressively reducing the gap until it becomes 1 (regular insertion sort).
// The array becomes increasingly sorted with each pass, making the final insertion sort very fast.
//
// Performance characteristics:
// - Time: O(N lg N) - depends on gap sequence
// - Space: O(1) - in-place sorting
// - Stable: No
// - In-place: Yes
// - Notes: Simple code, subquadratic performance
//
// Use Shell Sort when:
// - Simple implementation is desired
// - Small to medium-sized arrays
// - Reasonable performance is acceptable
// - Cannot afford O(N) extra space for merge sort
func ShellSort(arr []int) {
	n := len(arr)

	// Calculate initial gap using Knuth's sequence: h = 3h + 1
	// Sequence: 1, 4, 13, 40, 121, 364, 1093...
	// Start with largest gap less than n/3
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}

	// Perform h-sorting with decreasing gaps
	for h >= 1 {
		// h-sort the array (insertion sort with gap h)
		for i := h; i < n; i++ {
			// Insert arr[i] among arr[i-h], arr[i-2h], arr[i-3h]...
			// Move leftward by h steps while current element is smaller
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				// Swap elements that are h positions apart
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		// Reduce gap by dividing by 3 (inverse of 3h+1)
		h = h / 3
	}
	// Final pass with h=1 is regular insertion sort on a nearly sorted array
}
