package p543

import (
	"testing"

	"leetcode/lib"
)

// Benchmark pairs with same node count
func BenchmarkDiameterOfBinaryTree_100Nodes_Balanced(b *testing.B) {
	root := createBalancedTreeWithNodes(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		diameterOfBinaryTree(root)
	}
}

func BenchmarkDiameterOfBinaryTree_100Nodes_Skewed(b *testing.B) {
	root := createSkewedTreeWithNodes(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		diameterOfBinaryTree(root)
	}
}

func BenchmarkDiameterOfBinaryTree_1000Nodes_Balanced(b *testing.B) {
	root := createBalancedTreeWithNodes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		diameterOfBinaryTree(root)
	}
}

func BenchmarkDiameterOfBinaryTree_1000Nodes_Skewed(b *testing.B) {
	root := createSkewedTreeWithNodes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		diameterOfBinaryTree(root)
	}
}

func BenchmarkDiameterOfBinaryTree_10000Nodes_Balanced(b *testing.B) {
	root := createBalancedTreeWithNodes(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		diameterOfBinaryTree(root)
	}
}

func BenchmarkDiameterOfBinaryTree_10000Nodes_Skewed(b *testing.B) {
	root := createSkewedTreeWithNodes(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		diameterOfBinaryTree(root)
	}
}

// Create balanced tree with exactly nodeCount nodes
func createBalancedTreeWithNodes(nodeCount int) *lib.TreeNode {
	if nodeCount <= 0 {
		return nil
	}

	nodes := make([]*lib.TreeNode, nodeCount)
	for i := 0; i < nodeCount; i++ {
		nodes[i] = &lib.TreeNode{Val: i + 1}
	}

	// Build complete binary tree level by level
	for i := 0; i < nodeCount/2; i++ {
		leftIdx := 2*i + 1
		rightIdx := 2*i + 2

		if leftIdx < nodeCount {
			nodes[i].Left = nodes[leftIdx]
		}
		if rightIdx < nodeCount {
			nodes[i].Right = nodes[rightIdx]
		}
	}

	return nodes[0]
}

// Create skewed tree with exactly nodeCount nodes (left-skewed)
func createSkewedTreeWithNodes(nodeCount int) *lib.TreeNode {
	if nodeCount <= 0 {
		return nil
	}

	root := &lib.TreeNode{Val: 1}
	current := root

	for i := 1; i < nodeCount; i++ {
		current.Left = &lib.TreeNode{Val: i + 1}
		current = current.Left
	}

	return root
}
