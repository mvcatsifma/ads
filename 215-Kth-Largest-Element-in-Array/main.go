package p215

import "container/heap"

// findKthLargest finds the kth largest element in an unsorted array.
// k is 1-indexed, meaning k=1 returns the largest element, k=2 returns the second largest, etc.
// Example: findKthLargest([3,2,1,5,6,4], 2) returns 5 (second largest)
//
// Approach: Uses a max-heap to efficiently extract the k largest elements.
// Adds all elements to the heap, then pops k times to reach the kth largest.
// Time Complexity: O(n log n) where n is the length of nums (heap operations).
// Space Complexity: O(n) for storing all elements in the heap.
func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	// Add all numbers to max-heap
	for _, num := range nums {
		heap.Push(h, num)
	}

	var retval int
	// Pop k times to get the kth largest element
	for range k {
		retval = heap.Pop(h).(int)
	}

	return retval
}

// IntHeap implements a max-heap of integers.
// Satisfies the heap.Interface for use with the container/heap package.
// Larger values have higher priority (max-heap behavior).
type IntHeap []int

// Len returns the number of elements in the heap.
func (h IntHeap) Len() int { return len(h) }

// Less compares two elements for max-heap ordering (larger value = higher priority).
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

// Swap exchanges two elements in the heap.
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap. Uses pointer receiver because it modifies slice length.
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the largest element from the heap. Uses pointer receiver.
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
