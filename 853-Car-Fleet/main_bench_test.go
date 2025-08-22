package p853

import "testing"

func BenchmarkCarFleet(b *testing.B) {
	// Test case with moderate size - realistic scenario
	target := 1000
	position := make([]int, 100)
	speed := make([]int, 100)

	// Generate test data: cars at various positions with random speeds
	for i := 0; i < 100; i++ {
		position[i] = i * 10    // positions: 0, 10, 20, ..., 990
		speed[i] = (i % 10) + 1 // speeds: 1-10 (cyclical)
	}

	b.ResetTimer()
	b.ReportAllocs() // Enable allocation reporting
	for i := 0; i < b.N; i++ {
		carFleet(target, position, speed)
	}
}
