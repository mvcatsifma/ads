package p138

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// First pass, map original nodes to new nodes
	nodes := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		nodes[curr] = &Node{
			Val: curr.Val,
		}
		curr = curr.Next
	}

	// Second pass, build new list
	dummy := &Node{}
	c1 := head
	c2 := dummy
	for c1 != nil {
		newNode := nodes[c1]
		if c1.Random != nil {
			newNode.Random = nodes[c1.Random]
		}
		c2.Next = newNode
		c1 = c1.Next
		c2 = c2.Next
	}

	// Return head
	return dummy.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
