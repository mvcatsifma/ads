package p347

import (
	"container/heap"
)

type pair struct {
	num  int
	freq int
}

type pairHeap []pair

func (p pairHeap) Len() int {
	return len(p)
}

func (p pairHeap) Less(i, j int) bool {
	return p[i].freq > p[j].freq // max-heap!
}

func (p pairHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pairHeap) Push(x any) {
	*p = append(*p, x.(pair))
}

func (p *pairHeap) Pop() any {
	old := *p         // store the underlying slice in a local variable
	n := len(old)     // get the length of the slice
	x := old[n-1]     // extract the last element
	*p = old[0 : n-1] // shrink the slice by one element
	return x          // return the value that was at the end of the slice
}

func topKFrequent(nums []int, topK int) []int {
	var counts = make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	pairHeap := &pairHeap{}
	for num, freq := range counts {
		heap.Push(pairHeap, pair{num, freq})
	}

	var result []int
	for i := 0; i < topK; i++ {
		result = append(result, heap.Pop(pairHeap).(pair).num)
	}

	return result
}
