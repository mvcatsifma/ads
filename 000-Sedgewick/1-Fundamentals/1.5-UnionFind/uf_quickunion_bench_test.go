package unionfind

import (
	"testing"
)

func BenchmarkQuickUnion_Find(b *testing.B) {
	// Quick-Find has O(1) find - should be very fast
	uf := NewQuickUnion(10000)
	for i := 0; i < 9999; i++ {
		uf.union(i, i+1)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		uf.find(i % 10000)
	}
}

func BenchmarkQuickUnion_Union(b *testing.B) {
	size := 1000
	var count int // Prevent optimization

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		uf := NewQuickUnion(size)
		uf.union(0, size-1)
		count += uf.count // Use the result
	}

	_ = count // Prevent "unused variable" error
}
