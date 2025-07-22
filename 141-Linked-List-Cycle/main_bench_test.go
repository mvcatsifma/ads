package main

import (
	"fmt"
	"testing"
)

func Benchmark_hasCycle(b *testing.B) {
	// Test cases with different sizes
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		// Test non-cyclic list
		b.Run(fmt.Sprintf("NoCycle_Size_%d", size), func(b *testing.B) {
			list := createLinkedListWithSize(size, false)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hasCycle(list)
			}
		})

		// Test cyclic list
		b.Run(fmt.Sprintf("WithCycle_Size_%d", size), func(b *testing.B) {
			list := createLinkedListWithSize(size, true)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				hasCycle(list)
			}
		})
	}
}
