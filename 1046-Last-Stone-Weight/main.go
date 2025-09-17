package p1046

import "container/heap"

// lastStoneWeight simulates the stone smashing game where stones are repeatedly
// smashed together until at most one remains. The game follows these rules:
// 1. Choose the two heaviest stones and smash them together
// 2. If both stones have the same weight, both are destroyed
// 3. If stones have different weights (x < y), the heavier stone becomes (y - x)
// 4. Continue until 0 or 1 stones remain
//
// The algorithm uses a max-heap to efficiently find the two heaviest stones in
// O(log n) time. Since Go's container/heap implements a min-heap by default,
// this assumes IntHeap is configured as a max-heap (Less function returns h[i] > h[j]).
//
// Time complexity: O(n log n) where n is the number of stones
// - O(n log n) to build initial heap
// - O(n log n) for up to n iterations of pop/push operations
//
// Space complexity: O(n) for the heap storage
//
// Example:
// stones = [2,7,4,1,8,1]
// Step 1: Smash 8 and 7 → remaining: [2,4,1,1,1]
// Step 2: Smash 4 and 2 → remaining: [1,1,1,2]
// Step 3: Smash 2 and 1 → remaining: [1,1,1]
// Step 4: Smash 1 and 1 → remaining: [1]
// Result: 1
//
// Returns the weight of the last remaining stone, or 0 if no stones remain.
func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range stones {
		heap.Push(h, num)
	}

	for h.Len() > 1 {
		y := heap.Pop(h).(int)
		x := heap.Pop(h).(int)

		if y == x {
			continue
		} else {
			heap.Push(h, y-x)
		}
	}

	if h.Len() == 1 {
		return heap.Pop(h).(int)
	}

	return 0
}

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // sort in descending order (max-heap)
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() int {
	vals := *h
	x := vals[0]
	return x
}
