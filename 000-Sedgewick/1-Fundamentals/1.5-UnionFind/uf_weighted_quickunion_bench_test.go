package unionfind

import (
	"testing"
)

func BenchmarkWeightedQuickUnion_Find(b *testing.B) {
	// Setup: Create structure with many existing unions
	size := 10000
	uf := NewWeightedQuickUnion(size)

	// Build tree by sequential unions
	for i := 0; i < size-1; i++ {
		uf.union(i, i+1)
	}

	b.ResetTimer()

	// Benchmark find operations
	for i := 0; i < b.N; i++ {
		uf.find(i % size)
	}
}

func BenchmarkWeightedQuickUnion_Union(b *testing.B) {
	size := 1000000
	uf := NewWeightedQuickUnion(size)

	b.ResetTimer()

	// Benchmark union operations
	for i := 0; i < b.N; i++ {
		p := (i * 2) % (size - 1)
		q := (i*2 + 1) % size
		uf.union(p, q)
	}
}
