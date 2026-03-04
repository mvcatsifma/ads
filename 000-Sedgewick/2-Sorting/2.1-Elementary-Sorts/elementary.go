package elementary

// SelectionSort sorts an array using the selection sort algorithm.
//
// Algorithm: Repeatedly find the minimum element from the unsorted portion
// and place it at the beginning. After each iteration, the sorted portion
// grows by one element from the left.
//
// Time complexity: O(n²) - always makes n²/2 comparisons regardless of input order
// Space complexity: O(1) - in-place sorting, only uses constant extra space
//
// Best for: Small arrays or when minimizing swaps is important (exactly n swaps)
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
// Time complexity: O(n²) worst case, O(n) best case (already sorted)
// Space complexity: O(1) - in-place sorting
func InsertionSort(arr []int) {
	// TODO: Implement insertion sort
}

// ShellSort sorts an array using the shell sort algorithm.
// Time complexity: O(n^(3/2)) average case (depends on gap sequence)
// Space complexity: O(1) - in-place sorting
func ShellSort(arr []int) {
	// TODO: Implement shell sort
}
