package p133

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[int]bool)
	currNodes := make(map[int]*Node)
	newNodes := make(map[int]*Node)

	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil || visited[n.Val] {
			return
		}
		visited[n.Val] = true
		currNodes[n.Val] = n

		if _, ok := newNodes[n.Val]; !ok {
			newNodes[n.Val] = &Node{
				Val: n.Val,
			}
		}

		for _, neighbor := range n.Neighbors {
			dfs(neighbor)
		}
	}
	dfs(node)

	for _, currNode := range currNodes {
		newNode := newNodes[currNode.Val]
		for _, neighbor := range currNode.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, newNodes[neighbor.Val])
		}
	}

	return newNodes[1]
}
