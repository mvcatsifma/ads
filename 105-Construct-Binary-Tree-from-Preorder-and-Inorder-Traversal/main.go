package p105

import (
	"slices"

	"leetcode/lib"
)

// buildTree constructs a binary tree from preorder and inorder traversals.
// Preorder traversal visits nodes in order: root, left subtree, right subtree.
// Inorder traversal visits nodes in order: left subtree, root, right subtree.
// By combining both, the root value and left/right subtree boundaries can be determined uniquely.
// Example: preorder=[3,9,20,15,7], inorder=[9,3,15,20,7] builds tree with root 3, left subtree [9], right subtree [15,20,7]
//
// Approach: Recursively build tree by:
// 1. Taking first element of preorder as root
// 2. Finding root's position in inorder to split left and right subtrees
// 3. Partitioning preorder into left and right subtree ranges
// 4. Recursively building left and right subtrees
// Time Complexity: O(n^2) in worst case due to slices.Index() lookup on each recursion.
// Space Complexity: O(n) for recursion stack and O(n) for sliced arrays.
func buildTree(preorder []int, inorder []int) *lib.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &lib.TreeNode{
		Val: preorder[0],
	}
	mid := slices.Index(inorder, preorder[0])
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}
