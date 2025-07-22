package main

import (
	"fmt"
	"testing"
)

func Benchmark_reorderList(b *testing.B) {
	// Test different list sizes
	sizes := []int{4, 10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			// Create test data - sequence 1 to size
			nums := make([]int, size)
			for i := range nums {
				nums[i] = i + 1
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Create fresh list for each iteration
				list := createLinkedListFromNums(nums, -1)
				reorderList(list)
			}
		})
	}
}

// Helper function if not already defined
func createLinkedListFromNums(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}
