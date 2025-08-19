package p287

import "testing"

func BenchmarkFindDuplicate(b *testing.B) {
	// Create test array: [1,2,3,4,5,6,7,8,9,5] - duplicate is 5
	nums := make([]int, 1000)
	for i := 0; i < 999; i++ {
		nums[i] = (i % 999) + 1
	}
	nums[999] = 500 // Add duplicate

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findDuplicate(nums)
	}
}
