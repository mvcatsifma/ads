package p703

import "container/heap"

type KthLargest struct {
	k    int
	heap *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}

	for h.Len() > k {
		heap.Pop(h)
	}

	return KthLargest{
		k:    k,
		heap: h,
	}
}

func (l *KthLargest) Add(val int) int {
	if l.heap.Len() < l.k || val > l.heap.Peek() {
		heap.Push(l.heap, val)
		if l.heap.Len() > l.k {
			heap.Pop(l.heap)
		}
	}

	return l.heap.Peek()
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // sort in ascending order (min-heap)
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
