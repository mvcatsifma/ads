package p230

import "leetcode/lib"

// kthSmallest finds the kth smallest element in a binary search tree.
// Returns the value of the kth smallest node (1-indexed: k=1 returns smallest, k=2 returns second smallest, etc.)
// Example: tree [3,1,4,null,2] with k=1 returns 1, k=2 returns 2
//
// Approach: Uses DFS in-order traversal (left, node, right) which naturally produces sorted order in a BST.
// Aborts early once k elements are collected, avoiding unnecessary traversal of remaining nodes.
// Time Complexity: O(k) best case, O(n) worst case where n is the number of nodes.
// Space Complexity: O(k) for storing k values, O(h) for recursion stack where h is tree height.
func kthSmallest(root *lib.TreeNode, k int) int {
	vals := []int{}

	// In-order DFS traversal: left subtree -> node -> right subtree
	// This naturally produces values in sorted ascending order for a BST
	var dfs func(*lib.TreeNode)
	dfs = func(node *lib.TreeNode) {
		if node == nil || len(vals) == k {
			return
		}

		dfs(node.Left) // Traverse left subtree first
		if len(vals) < k {
			vals = append(vals, node.Val) // Process node (adds to sorted sequence)
		}
		dfs(node.Right) // Traverse right subtree last
	}
	dfs(root)

	return vals[k-1]
}
