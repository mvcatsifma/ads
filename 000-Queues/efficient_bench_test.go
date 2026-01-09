package p000Queues

import "testing"

func BenchmarkComparison(b *testing.B) {
	const queueSize = 10000

	b.Run("Efficient", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			q := NewEfficientIntQueue()
			for j := 0; j < queueSize; j++ {
				q.Enqueue(j)
			}
			b.StartTimer()

			for j := 0; j < queueSize; j++ {
				q.Dequeue()
			}
		}
	})

	b.Run("Simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			q := &SimpleFIFOIntQueue{items: make([]int, 0, queueSize)}

			dequeue := func(q *SimpleFIFOIntQueue) (int, bool) {
				if len(q.items) == 0 {
					return 0, false
				}
				item := q.items[0]
				q.items = q.items[1:] // O(n) slice operation
				return item, true
			}

			for j := 0; j < queueSize; j++ {
				q.items = append(q.items, j)
			}
			b.StartTimer()

			for j := 0; j < queueSize; j++ {
				dequeue(q)
			}
		}
	})
}
