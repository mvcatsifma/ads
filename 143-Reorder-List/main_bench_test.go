package main

import (
	"fmt"
	"testing"

	lib "leetcode-lib"
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
				list := lib.CreateLinkedListFromNums(nums, -1)
				reorderList(list)
			}
		})
	}
}
