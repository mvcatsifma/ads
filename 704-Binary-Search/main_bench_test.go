package p704

import (
	"math/rand"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	// Generate a sorted slice of 1 million integers
	nums := make([]int, 1000000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// Search for a random target in the slice
		target := rand.Intn(1000000)
		search(nums, target)
	}
}
