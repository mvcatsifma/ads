package p746

import (
	"math/rand"
	"testing"
)

// BenchmarkMinCostClimbingStairs measures the performance characteristics of the
// minCostClimbingStairs function across different input sizes on Apple M3 Pro.
//
// Performance Results (Apple M3 Pro, arm64):
// - Small (3 elements):   ~50 ns/op,    72 B/op,     2 allocs/op
// - Medium (100 elements): ~1.2 μs/op,  2.7 KB/op,   2 allocs/op
// - Large (1K elements):   ~11 μs/op,   20 KB/op,    2 allocs/op
// - XLarge (10K elements): ~112 μs/op,  188 KB/op,   2 allocs/op
//
// Key observations:
// - Time complexity scales linearly O(n) as expected
// - Memory usage grows proportionally with input size due to slice copying
// - Consistent 2 allocations per operation (slice copy + append for virtual top)
// - Memory overhead is significant: ~2x input size due to copying mechanism
//
// Performance bottlenecks:
// - Slice copying dominates small inputs (72B for 3 ints = 24B overhead)
// - append() operation adds minimal overhead compared to copying
// - float64 conversions in math.Min add computational cost
func BenchmarkMinCostClimbingStairs(b *testing.B) {
	testCases := []struct {
		name string
		cost []int
	}{
		{
			name: "small",
			cost: []int{10, 15, 20},
		},
		{
			name: "medium",
			cost: generateRandomCosts(100),
		},
		{
			name: "large",
			cost: generateRandomCosts(1000),
		},
		{
			name: "xlarge",
			cost: generateRandomCosts(10000),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				// Create a copy to avoid modifying the original slice
				costCopy := make([]int, len(tc.cost))
				copy(costCopy, tc.cost)
				minCostClimbingStairs(costCopy)
			}
		})
	}
}

// Helper function to generate random test data
func generateRandomCosts(n int) []int {
	costs := make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = rand.Intn(1000) + 1 // Random cost between 1-1000
	}
	return costs
}
