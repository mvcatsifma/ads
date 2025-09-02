package lib

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree constructs a binary tree from a level-order slice representation.
// Uses -1 as a sentinel value to represent nil nodes in the input slice.
// Non-nil nodes (including negative values other than -1) are created normally.
//
// Example:
//   nums := []int{3, 9, 20, -1, -1, 15, 7}
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
	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		// Process left child
		if i < len(nums) {
			if nums[i] != -1 {
				node.Left = &TreeNode{Val: nums[i]}
				queue = append(queue, node.Left)
			}
			i++
		}

		// Process right child
		if i < len(nums) {
			if nums[i] != -1 {
				node.Right = &TreeNode{Val: nums[i]}
				queue = append(queue, node.Right)
			}
			i++
		}
	}

	return root
}
