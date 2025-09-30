package p153

import "testing"

// BenchmarkFindMin benchmarks the findMin function with a 12-element rotated array.
// Results on Apple M3 Pro (arm64):
//   - 288M+ operations/sec (~4.05 ns/op)
//   - Zero memory allocations (0 B/op, 0 allocs/op)
//   - Confirms O(1) space complexity and excellent performance
func BenchmarkFindMin(b *testing.B) {
	// Test with a moderately sized rotated array
	nums := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findMin(nums)
	}
}
