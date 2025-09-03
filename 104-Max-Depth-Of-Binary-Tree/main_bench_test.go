package p104

import (
	"testing"

	"leetcode/lib"
)

// Benchmark pairs with same node count
func BenchmarkMaxDepth_100Nodes_Balanced(b *testing.B) {
	root := createBalancedTreeWithNodes(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepth(root)
	}
}

func BenchmarkMaxDepth_100Nodes_Skewed(b *testing.B) {
	root := createSkewedTreeWithNodes(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepth(root)
	}
}

func BenchmarkMaxDepth_1000Nodes_Balanced(b *testing.B) {
	root := createBalancedTreeWithNodes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepth(root)
	}
}

func BenchmarkMaxDepth_1000Nodes_Skewed(b *testing.B) {
	root := createSkewedTreeWithNodes(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepth(root)
	}
}

func BenchmarkMaxDepth_10000Nodes_Balanced(b *testing.B) {
	root := createBalancedTreeWithNodes(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepth(root)
	}
}

func BenchmarkMaxDepth_10000Nodes_Skewed(b *testing.B) {
	root := createSkewedTreeWithNodes(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepth(root)
	}
}

// Test to verify node counts are correct
func TestNodeCounts(t *testing.T) {
	testCases := []int{100, 1000, 10000}

	for _, nodeCount := range testCases {
		balanced := createBalancedTreeWithNodes(nodeCount)
		skewed := createSkewedTreeWithNodes(nodeCount)

		balancedCount := countNodes(balanced)
		skewedCount := countNodes(skewed)

		if balancedCount != nodeCount {
			t.Errorf("Balanced tree: expected %d nodes, got %d", nodeCount, balancedCount)
		}

		if skewedCount != nodeCount {
			t.Errorf("Skewed tree: expected %d nodes, got %d", nodeCount, skewedCount)
		}

		t.Logf("%d nodes - Balanced depth: %d, Skewed depth: %d",
			nodeCount, maxDepth(balanced), maxDepth(skewed))
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

// Helper to count nodes (for verification)
func countNodes(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
