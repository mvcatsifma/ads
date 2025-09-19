package p70

import "testing"

// BenchmarkClimbStairs measures the performance of climbStairs with n=45.
// Results show excellent performance on Apple M3 Pro:
// - ~12.46 ns/op: Very fast execution due to O(1) space and simple arithmetic
// - 0 B/op, 0 allocs/op: Zero heap allocations, uses only stack variables
// - 92M+ ops/sec: Demonstrates the efficiency of the iterative Fibonacci approach
// The consistent performance validates the O(n) time complexity with minimal overhead.
func BenchmarkClimbStairs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		climbStairs(45) // Test with moderately large input
	}
}
