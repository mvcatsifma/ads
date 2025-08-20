package p46

import "testing"

// Benchmark test
func BenchmarkLRUCache(b *testing.B) {
	cache := Constructor(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Put(i%2000, i)
		cache.Get(i % 1000)
	}
}
