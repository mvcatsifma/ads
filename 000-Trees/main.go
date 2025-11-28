package p000Trees

import (
	"container/list"

	"leetcode/lib"
)

// Example: tree [1,2,3] returns [1,2,3]
//
// Approach: Uses recursive DFS, processing each node before its children.
// Time Complexity: O(n) where n is the number of nodes (visit each node once).
// Space Complexity: O(h) for recursion stack where h is the height of the tree.
func dfsPreorder(root *lib.TreeNode) []int {
	var retval []int
	var dfs func(*lib.TreeNode)
	dfs = func(node *lib.TreeNode) {
		if node == nil {
			return
		}
		retval = append(retval, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return retval
}

// dfsInorder performs a depth-first inorder traversal of a binary tree.
// Inorder traversal visits nodes in order: left subtree, root, right subtree.
// For a binary search tree, inorder traversal produces values in sorted ascending order.
// Example: tree [5,3,7,2,4,6,8] returns [2,3,4,5,6,7,8]
//        5
//       / \
//      3   7
//     / \ / \
//    2  4 6  8
// Inorder visits: 2(left-left), 3(left), 4(left-right), 5(root), 6(right-left), 7(right), 8(right-right)
//
// Approach: Uses recursive DFS, processing each node between its left and right children.
// Time Complexity: O(n) where n is the number of nodes (visit each node once).
// Space Complexity: O(h) for recursion stack where h is the height of the tree.
func dfsInorder(root *lib.TreeNode) []int {
	var retval []int
	var dfs func(*lib.TreeNode)
	dfs = func(node *lib.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		retval = append(retval, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return retval
}

// dfsPostorder performs a depth-first postorder traversal of a binary tree.
// Postorder traversal visits nodes in order: left subtree, right subtree, root.
// Postorder is useful for operations that require processing children before parents (e.g., tree deletion, calculating tree height).
// Example: tree [5,3,7,2,4,6,8] returns [2,4,3,6,8,7,5]
//        5
//       / \
//      3   7
//     / \ / \
//    2  4 6  8
// Postorder visits: 2(left-left), 4(left-right), 3(left), 6(right-left), 8(right-right), 7(right), 5(root)
//
// Approach: Uses recursive DFS, processing each node after its children.
// Time Complexity: O(n) where n is the number of nodes (visit each node once).
// Space Complexity: O(h) for recursion stack where h is the height of the tree.
func dfsPostorder(root *lib.TreeNode) []int {
	var retval []int
	var dfs func(*lib.TreeNode)
	dfs = func(node *lib.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		retval = append(retval, node.Val)
	}

	dfs(root)
	return retval
}

// bfs performs a breadth-first search (level-order) traversal of a binary tree.
// Returns all node values in level-order: root, then all nodes at depth 1, then depth 2, etc.
// Example: tree [5,3,7,2,4,6,8] returns [5,3,7,2,4,6,8]
//        5
//       / \
//      3   7
//     / \ / \
//    2  4 6  8
// Level order: [5] at level 0, [3,7] at level 1, [2,4,6,8] at level 2
//
// Approach: Uses a queue (container/list) to process nodes level-by-level.
// For each level, process all nodes currently in the queue before moving to the next level.
// Only adds non-nil children to the queue to avoid wasting space and iterations.
// Time Complexity: O(n) where n is the number of nodes (visit each node once).
// Space Complexity: O(w) where w is the maximum width of the tree (max nodes at any level).
func bfs(root *lib.TreeNode) []int {
	var retval []int

	// Initialize queue with root
	l := list.New()
	l.PushBack(root)

	// Process level by level
	for l.Len() > 0 {
		length := l.Len() // Number of nodes at current level

		// Process all nodes at current level
		for i := 0; i < length; i++ {
			node := l.Remove(l.Front()).(*lib.TreeNode)
			retval = append(retval, node.Val)

			// Only add non-nil children to queue
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
	}

	return retval
}
