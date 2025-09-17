package p1046

import (
	"math/rand"
	"testing"
)

func BenchmarkLastStoneWeight(b *testing.B) {
	// Create test data with varying stone weights
	stones := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		stones[i] = rand.Intn(1000) + 1 // weights 1-1000
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		// Make a copy since the function modifies the heap
		testStones := make([]int, len(stones))
		copy(testStones, stones)

		lastStoneWeight(testStones)
	}
}
