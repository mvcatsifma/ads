package priorityqueus

// sort performs heapsort on arr using 1-indexed heap operations.
// arr[0] is unused; actual elements are at indices 1 to len(arr)-1.
func sort(arr []int) {
	n := len(arr) - 1 // Exclude arr[0] (unused in 1-indexed heap)

	// Phase 1: Build max heap (heapify)
	// Iterate from last parent node (n/2) toward root (1)
	// For each node, sink moves it down the tree toward leaves
	// Result: max heap property satisfied (parent >= children at all nodes)
	//         Maximum element is now at root (position 1)
	//         Array is heap-ordered but NOT yet sorted
	for k := n / 2; k >= 1; k-- {
		sink(arr, k, n)
	}

	// Phase 2: Sort down
	// Repeatedly:
	// 1. Pick max (at root position 1)
	// 2. Swap with last element in heap (position n) - places max in final sorted position
	// 3. Shrink heap (decrement n) to exclude sorted element
	// 4. Heapify: sink new root to restore max heap property
	// 5. Repeat until heap is empty
	for n > 1 {
		exch(arr, 1, n) // Move max to position n
		n -= 1          // Reduce heap size
		sink(arr, 1, n) // Restore heap property
	}
}

// exch swaps elements at positions i and k in the array.
func exch(arr []int, i int, k int) {
	arr[i], arr[k] = arr[k], arr[i]
}

// sink restores heap order by moving element at position k down the tree.
// Only considers elements in range [1, n].
func sink(arr []int, k int, n int) {
	for 2*k <= n {
		j := 2 * k
		// Choose the larger of the two children
		if j < n && less(arr, j, j+1) {
			j++
		}
		// If parent is not less than larger child, heap order is satisfied
		if !less(arr, k, j) {
			break
		}
		exch(arr, k, j)
		k = j
	}
}

// less compares elements at positions k and j.
// Returns true if arr[k] < arr[j].
func less(arr []int, k int, j int) bool {
	return arr[k] < arr[j]
}
