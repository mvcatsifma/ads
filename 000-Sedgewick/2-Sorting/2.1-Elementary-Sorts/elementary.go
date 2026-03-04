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
//
// Algorithm: Like sorting playing cards in your hand - pick each element
// and insert it into its correct position among the already sorted elements
// by repeatedly swapping it leftward until it's in place.
//
// Time complexity: O(n²) worst case (reverse sorted), O(n) best case (already sorted)
// Space complexity: O(1) - in-place sorting, only uses constant extra space
//
// Best for: Small arrays or nearly sorted data (very efficient for partially sorted arrays)
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
// Time complexity: O(n^(3/2)) average case (depends on gap sequence)
// Space complexity: O(1) - in-place sorting
func ShellSort(arr []int) {
	// TODO: Implement shell sort
}
