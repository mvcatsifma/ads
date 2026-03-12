package mergesort

// MergeSort sorts an array using the merge sort algorithm (top-down recursive).
//
// Algorithm: Classic divide-and-conquer approach that recursively splits the array
// in half until reaching single-element subarrays, then merges them back together
// in sorted order. Uses an auxiliary array to avoid expensive in-place merging.
//
// Time complexity: O(n log n) - always makes log(n) levels of splits with n work per level
// Space complexity: O(n) - requires auxiliary array for merging (allocated once, reused)
//
// Best for: General-purpose sorting when O(n log n) guarantee is needed regardless of input.
// Unlike QuickSort, performance is consistent but requires extra space.
func MergeSort(arr []int) {
	// Allocate auxiliary array once and reuse across all recursive calls
	aux := make([]int, len(arr))
	sort(arr, aux, 0, len(arr)-1)
}

// MergeSortBU sorts an array using bottom-up merge sort (iterative).
//
// Algorithm: Non-recursive variant of merge sort that builds sorted subarrays from
// the bottom up. Starts by merging subarrays of size 1, then 2, then 4, and so on,
// doubling the size with each pass until the entire array is sorted.
//
// Time complexity: O(n log n) - always, log(n) passes with n work per pass
// Space complexity: O(n) - requires auxiliary array for merging
//
// Best for: Situations where recursion is problematic (limited stack space) or when
// better cache locality is desired. Same asymptotic performance as top-down but
// avoids function call overhead.
func MergeSortBU(arr []int) {
	n := len(arr)
	// Allocate auxiliary array once
	aux := make([]int, n)

	// Outer loop: double subarray size each pass (1 → 2 → 4 → 8 → ...)
	// length represents the size of subarrays to merge in this pass
	for length := 1; length < n; length *= 2 {
		// Inner loop: merge all pairs of subarrays of size 'length'
		// Step by 2*length since we're merging two subarrays at a time
		for lo := 0; lo < n-length; lo += length + length {
			// Merge arr[lo:lo+length] with arr[lo+length:lo+2*length]
			// Use min() to handle last subarray which may be shorter than 'length'
			merge(arr, aux, lo, lo+length-1, min(lo+length+length-1, n-1))
		}
	}
}

// sort recursively sorts arr[lo:hi+1] using divide-and-conquer
func sort(arr []int, aux []int, lo int, hi int) {
	// Base case: single element (or empty) is already sorted
	if hi <= lo {
		return
	}

	// Divide: find midpoint (avoids overflow compared to (lo+hi)/2)
	mid := lo + (hi-lo)/2

	// Conquer: recursively sort left and right halves
	sort(arr, aux, lo, mid)   // Sort left half
	sort(arr, aux, mid+1, hi) // Sort right half

	// Combine: merge the two sorted halves
	merge(arr, aux, lo, mid, hi)
}

// merge combines two sorted subarrays arr[lo:mid] and arr[mid+1:hi] into a single sorted array
func merge(arr []int, aux []int, lo int, mid int, hi int) {
	// Initialize pointers to start of left and right subarrays
	i, j := lo, mid+1

	// Copy current subarray to auxiliary array
	// We merge from aux back into arr to avoid overwriting elements we haven't processed yet
	for k := lo; k <= hi; k++ {
		aux[k] = arr[k]
	}

	// Merge back to arr in sorted order
	for k := lo; k <= hi; k++ {
		if i > mid {
			// Left subarray exhausted, take from right
			arr[k] = aux[j]
			j++
		} else if j > hi {
			// Right subarray exhausted, take from left
			arr[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			// Both available: right element is smaller, take from right
			arr[k] = aux[j]
			j++
		} else {
			// Both available: left element is smaller or equal, take from left
			arr[k] = aux[i]
			i++
		}
	}
}
