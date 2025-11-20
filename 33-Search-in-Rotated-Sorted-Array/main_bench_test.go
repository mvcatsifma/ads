package p33

import "testing"

func BenchmarkSearch(b *testing.B) {
	// Create a rotated sorted array for testing
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i + 1
	}
	// Rotate at position 3333: [3334, 3335, ..., 10000, 1, 2, 3, ..., 3333]
	rotated := append(nums[3333:], nums[:3333]...)
	target := 5000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search(rotated, target)
	}
}
