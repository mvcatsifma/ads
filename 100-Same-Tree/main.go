package p100

import (
	"math"
	"reflect"

	"leetcode/lib"
)

// isSameTree determines if two binary trees are structurally identical and have
// the same node values. Two trees are considered the same if they have identical
// structure and all corresponding nodes contain the same values.
//
// The algorithm uses preorder traversal (root -> left -> right) to serialize both
// trees into slices, using math.MaxInt as a sentinel value to represent nil nodes.
// This ensures that structural differences are captured - trees with different
// structures will produce different serializations even if they contain the same values.
//
// For example:
//   Tree 1:  1      Tree 2:  1
//           / \              \
//          2   3              2
//                            /
//                           3
// Serializes to: [1,2,MaxInt,MaxInt,3,MaxInt,MaxInt] vs [1,MaxInt,2,3,MaxInt,MaxInt,MaxInt]
//
// Time complexity: O(n) where n is the number of nodes in the larger tree
// Space complexity: O(n) for the serialization slices + O(h) for recursion stack
//
// Note: Uses preorder traversal with nil sentinels to ensure structural comparison.
func isSameTree(p *lib.TreeNode, q *lib.TreeNode) bool {
	pVals := make([]int, 0)
	qVals := make([]int, 0)

	var dfs func(*lib.TreeNode, *[]int)
	dfs = func(node *lib.TreeNode, vals *[]int) {
		if node == nil {
			*vals = append(*vals, math.MaxInt)
			return
		}
		*vals = append(*vals, node.Val)

		dfs(node.Left, vals)
		dfs(node.Right, vals)
	}

	dfs(p, &pVals)
	dfs(q, &qVals)

	return reflect.DeepEqual(pVals, qVals)
}
