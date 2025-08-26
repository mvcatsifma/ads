package p594

import "testing"

func BenchmarkFindLHS(b *testing.B) {
	// Test data with mixed scenarios
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7, 1, 1, 2, 3, 3, 4, 4, 4, 5}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		findLHS(nums)
	}
}
