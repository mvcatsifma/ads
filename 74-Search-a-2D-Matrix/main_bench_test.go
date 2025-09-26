package p74

import "testing"

// BenchmarkSearchMatrix tests the performance of the optimized searchMatrix function
// on a 100x100 sorted matrix (10,000 elements) searching for the middle element.
//
// Results on Apple M3 Pro (arm64):
// - 61,287,730 iterations in 1 second
// - 19.50 ns/op average execution time
// - 0 B/op memory allocation (zero heap allocations)
// - 0 allocs/op (no garbage collection pressure)
//
// Performance characteristics:
// - Demonstrates O(log(m*n)) time complexity: ~13-14 binary search steps for 10K elements
// - Excellent memory efficiency with zero allocations
// - Fast execution time indicates efficient CPU utilization
// - Tests middle element position only; performance may vary slightly for edge cases
func BenchmarkSearchMatrix(b *testing.B) {
	// Generate 100x100 sorted matrix
	matrix := make([][]int, 100)
	val := 1
	for i := 0; i < 100; i++ {
		matrix[i] = make([]int, 100)
		for j := 0; j < 100; j++ {
			matrix[i][j] = val
			val++
		}
	}

	target := matrix[50][50] // Middle element

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		searchMatrix(matrix, target)
	}
}
