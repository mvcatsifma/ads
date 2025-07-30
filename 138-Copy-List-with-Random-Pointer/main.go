package p138

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodeMap := make(map[*Node]*Node)

	// Helper function to get or create a node copy
	getOrCreateNode := func(original *Node) *Node {
		if original == nil {
			return nil
		}

		if copied, exists := nodeMap[original]; exists {
			return copied
		}

		// Create new node and store in map
		newNode := &Node{Val: original.Val}
		nodeMap[original] = newNode
		return newNode
	}

	// First pass: create all nodes and set Next pointers
	current := head
	var prev *Node
	var newHead *Node

	for current != nil {
		// Get or create current node copy
		currentCopy := getOrCreateNode(current)

		// Set as head if this is the first node
		if newHead == nil {
			newHead = currentCopy
		}

		// Link previous node to current
		if prev != nil {
			prev.Next = currentCopy
		}

		// Set random pointer
		currentCopy.Random = getOrCreateNode(current.Random)

		prev = currentCopy
		current = current.Next
	}

	return newHead
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
