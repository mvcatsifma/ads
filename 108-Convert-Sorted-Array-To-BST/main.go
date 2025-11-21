package p108

import "leetcode/lib"

// sortedArrayToBST converts a sorted array into a height-balanced binary search tree.
// Uses a divide-and-conquer approach by recursively selecting the middle element
// as the root to ensure the resulting BST maintains optimal O(log n) height.
//
// Time Complexity: O(n) - visits each element once
// Space Complexity: O(log n) - recursion stack depth for balanced tree
//
// Parameters:
//   nums: sorted integer array in ascending order
//
// Returns:
//   *TreeNode: root of the constructed height-balanced BST, or nil if input is empty
func sortedArrayToBST(nums []int) *lib.TreeNode {
	var dfs func(l int, r int) *lib.TreeNode
	dfs = func(l int, r int) *lib.TreeNode {
		if l > r {
			return nil
		}
		m := (l + r) / 2
		root := &lib.TreeNode{
			Val: nums[m],
		}

		root.Left = dfs(l, m-1)
		root.Right = dfs(m+1, r)
		return root
	}

	return dfs(0, len(nums)-1)
}
