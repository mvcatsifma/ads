package main

import (
	"fmt"
	"testing"
)

func BenchmarkRemoveNthFromEnd(b *testing.B) {
	// Test different list sizes
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		// Test different n values for each size
		nValues := []int{1, size / 2, size} // remove from end, middle, start

		for _, n := range nValues {
			b.Run(fmt.Sprintf("size_%d_n_%d", size, n), func(b *testing.B) {
				b.ResetTimer()
				b.ReportAllocs()

				for i := 0; i < b.N; i++ {
					// Create fresh list for each iteration
					list := createLinkedListWithSize(size, false)
					removeNthFromEnd(list, n)
				}
			})
		}
	}
}

func createLinkedListWithSize(size int, withCycle bool) *ListNode {
	if size == 0 {
		return nil
	}

	head := &ListNode{Val: 0}
	current := head
	var cyclePoint *ListNode

	// Create the list
	for i := 1; i < size; i++ {
		current.Next = &ListNode{Val: i}
		current = current.Next
		// Store middle node for creating cycle
		if i == size/2 {
			cyclePoint = current
		}
	}

	// Create cycle if requested
	if withCycle {
		current.Next = cyclePoint
	}

	return head
}
