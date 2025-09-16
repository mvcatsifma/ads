package p703

import (
	"math/rand"
	"testing"
)

func BenchmarkKthLargest_Add(b *testing.B) {
	// Setup: Create KthLargest with moderate initial data
	initialNums := make([]int, 100)
	for i := 0; i < 100; i++ {
		initialNums[i] = rand.Intn(1000)
	}

	kth := Constructor(3, initialNums)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		kth.Add(rand.Intn(1000))
	}
}
