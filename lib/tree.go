package lib

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree constructs a binary tree from a level-order (breadth-first) slice representation.
// The input slice represents nodes level by level, left to right.
// Returns nil if the input slice is empty.
//
// Example:
//   nums := []int{3, 9, 20, 15, 7}
//   Creates tree:
//       3
//      / \
//     9  20
//       /  \
//      15   7
//
// Time complexity: O(n) where n is len(nums)
// Space complexity: O(n) for the queue and tree nodes
func BuildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if i < len(nums) {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
			i++
		}

		if i < len(nums) {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
			i++
		}
	}

	return root
}
