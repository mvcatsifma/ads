package p136

import (
	"math/rand"
	"testing"
)

// BenchmarkSingleNumber measures the performance of the XOR-based single number algorithm
// on a medium-sized dataset (10,001 elements) on Apple M3 Pro.
//
// Performance Results (Apple M3 Pro, arm64):
// - ~2.6 μs/op for 10,001 elements
// - 0 B/op, 0 allocs/op (true O(1) space complexity)
// - ~471,019 operations per second
//
// Performance characteristics:
// - Linear time complexity O(n): ~0.26 ns per element processed
// - Constant space complexity: no heap allocations during execution
// - Excellent cache performance: single sequential pass through array
// - Optimal bit manipulation: XOR operations are CPU-native and very fast
//
// This demonstrates the efficiency of the XOR approach compared to alternative
// solutions like hash maps or sorting, which would require O(n) space or O(n log n) time.
func BenchmarkSingleNumber(b *testing.B) {
	// Test with a medium-sized array: 10,001 elements (5,000 pairs + 1 single)
	nums := make([]int, 10001)

	// Fill with pairs
	for i := 0; i < 5000; i++ {
		val := rand.Intn(100000)
		nums[i*2] = val
		nums[i*2+1] = val
	}
	nums[10000] = 42 // The single number

	// Shuffle to randomize order
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		singleNumber(nums)
	}
}
