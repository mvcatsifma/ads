package p230

import (
	"container/heap"

	"leetcode/lib"
)

// kthSmallest finds the kth smallest element in a binary search tree.
// Returns the value of the kth smallest node (1-indexed: k=1 returns smallest, k=2 returns second smallest, etc.)
// Example: tree [3,1,4,null,2] with k=1 returns 1, k=2 returns 2
//
// Approach: Uses DFS to collect all node values into a min-heap, then pops k-1 elements
// to reach the kth smallest element at the top of the heap.
// Note: This approach is inefficient for BSTs since in-order traversal would naturally produce sorted order.
// Time Complexity: O(n log n) where n is the number of nodes (build heap + k pops).
// Space Complexity: O(n) for storing all nodes in the heap.
func kthSmallest(root *lib.TreeNode, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	// DFS to collect all node values
	var dfs func(*lib.TreeNode)
	dfs = func(node *lib.TreeNode) {
		if node == nil {
			return
		}
		heap.Push(h, node.Val)

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	// Pop k-1 elements to reach the kth smallest
	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	// Return the kth smallest element
	return heap.Pop(h).(int)
}

// IntHeap is a min-heap implementation for integers.
// Implements the heap.Interface for use with container/heap package.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap: smaller values at top
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds element x to the heap. Uses pointer receiver because it modifies slice length.
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the smallest element from the heap. Uses pointer receiver.
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
