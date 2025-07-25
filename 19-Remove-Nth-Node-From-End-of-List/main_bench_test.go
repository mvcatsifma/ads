package p19

import (
	"fmt"
	"testing"

	lib "leetcode/lib"
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
					list := lib.CreateLinkedListWithSize(size, false)
					removeNthFromEnd(list, n)
				}
			})
		}
	}
}
