package p739

import "testing"

func BenchmarkDailyTemperatures(b *testing.B) {
	// Generate test data with realistic temperature variations
	temperatures := make([]int, 10000)
	for i := 0; i < len(temperatures); i++ {
		temperatures[i] = 60 + (i % 40) // Range: 60-99°F with patterns
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = dailyTemperatures(temperatures)
	}
}
