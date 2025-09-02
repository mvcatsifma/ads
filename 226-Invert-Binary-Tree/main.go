package p226

// invertTree inverts a binary tree by swapping left and right children of all nodes.
// The inversion is performed recursively in-place, modifying the original tree structure.
// Returns the root of the inverted tree, or nil if the input tree is empty.
//
// Example:
//   Original:     Inverted:
//       4             4
//      / \           / \
//     2   7         7   2
//    / \ / \       / \ / \
//   1  3 6  9     9  6 3  1
//
// Time complexity: O(n) where n is the number of nodes
// Space complexity: O(h) where h is the height of the tree (recursion stack)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
