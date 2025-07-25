package p138

import (
	"fmt"
	"testing"
)

func BenchmarkCopyRandomList(b *testing.B) {
	// Test different list sizes
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		// Create test list once per size
		original := createTestList(size)

		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				copied := copyRandomList(original)
				// Prevent compiler optimization
				if copied == nil {
					b.Fatal("unexpected nil result")
				}
			}
		})
	}
}

// Helper to create a linked list with random pointers for testing
func createTestList(size int) *Node {
	// Create nodes
	nodes := make([]*Node, size)
	for i := 0; i < size; i++ {
		nodes[i] = &Node{Val: i}
	}

	// Connect Next pointers
	for i := 0; i < size-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// Set Random pointers with deterministic but varied patterns
	for i := 0; i < size; i++ {
		switch i % 4 {
		case 0:
			// Point to self
			nodes[i].Random = nodes[i]
		case 1:
			// Point to next (or first if last)
			nodes[i].Random = nodes[(i+1)%size]
		case 2:
			// Point to null
			nodes[i].Random = nil
		case 3:
			// Point to first node
			nodes[i].Random = nodes[0]
		}
	}

	return nodes[0]
}
