package p118

import (
	"fmt"
	"testing"
)

func BenchmarkGenerateOptimized(b *testing.B) {
	sizes := []int{10, 20, 30, 50, 100}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("rows_%d", size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_ = generate(size)
			}
		})
	}
}
