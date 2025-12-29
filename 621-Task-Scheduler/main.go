package p621

import (
	"container/heap"
)

func leastInterval(tasks []byte, n int) int {
	// Count frequency of each task
	freq := make(map[byte]int)
	for _, task := range tasks {
		freq[task]++
	}

	// Max-heap to store frequencies
	h := &IntHeap{}
	heap.Init(h)
	for _, count := range freq {
		heap.Push(h, count)
	}

	time := 0

	// Process tasks in cycles of (n+1) length
	for h.Len() > 0 {
		cycle := n + 1
		temp := []int{}
		tasksInCycle := 0

		// Fill one cycle (n+1 slots)
		for cycle > 0 && h.Len() > 0 {
			// Take most frequent task
			maxFreq := heap.Pop(h).(int)
			tasksInCycle++

			// If task has remaining instances, store for later
			if maxFreq > 1 {
				temp = append(temp, maxFreq-1)
			}
			cycle--
		}

		// Put remaining tasks back in heap
		for _, freq := range temp {
			heap.Push(h, freq)
		}

		// Add time for this cycle
		if h.Len() == 0 {
			// Last cycle - only count actual tasks
			time += tasksInCycle
		} else {
			// Full cycle needed (including idle time)
			time += n + 1
		}
	}

	return time
}

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
