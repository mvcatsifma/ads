package main

import (
	"fmt"
	"testing"

	lib "leetcode-lib"
)

func BenchmarkFrequenciesOfElements(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			list := lib.CreateLinkedListWithSize(size, false)
			b.ResetTimer()
			b.ReportAllocs() // Added allocation reporting for all sizes

			for i := 0; i < b.N; i++ {
				frequenciesOfElements(list)
			}
		})
	}
}
