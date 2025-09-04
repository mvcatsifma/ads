package p100

import (
	"leetcode/lib"
)

// isSameTree determines if two binary trees are structurally identical and have
// the same node values. Two trees are considered the same if they have identical
// structure and all corresponding nodes contain the same values.
//
// The algorithm uses recursive depth-first traversal to compare corresponding nodes
// simultaneously. At each step, it checks:
// 1. Both nodes are nil (base case - subtrees match)
// 2. Exactly one node is nil (structural mismatch)
// 3. Node values differ (value mismatch)
// 4. Recursively compare left and right subtrees
//
// The XOR check efficiently detects structural mismatches - if exactly one of the
// nodes is nil, the trees have different structures at this position.
//
// Time complexity: O(min(m,n)) where m,n are the number of nodes in each tree
// Space complexity: O(min(h1,h2)) where h1,h2 are the heights (recursion stack)
//
// Example:
//   Tree 1:  1      Tree 2:  1
//           / \              / \
//          2   3            2   3
// Returns true (identical structure and values)
//
//   Tree 1:  1      Tree 2:  1
//           /                \
//          2                  2
// Returns false (different structure)
func isSameTree(p *lib.TreeNode, q *lib.TreeNode) bool {
	var dfs func(*lib.TreeNode, *lib.TreeNode) bool
	dfs = func(p *lib.TreeNode, q *lib.TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if xor(p == nil, q == nil) {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		return dfs(p.Left, q.Left) && dfs(p.Right, q.Right)
	}

	return dfs(p, q)
}

func xor(a, b bool) bool {
	return a != b
}
