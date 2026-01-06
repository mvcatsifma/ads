package p133

// Node represents a vertex in an undirected graph.
type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph performs a deep copy using a clear two-pass design:
//   Pass 1 (DFS): discover all reachable nodes, record originals, and allocate a clone per value.
//   Pass 2: wire the cloned neighbor edges using the originals' adjacency lists.
//
// Assumptions match LeetCode 133:
//   - Unique Val in [1..100], graph is connected, entry node has Val == 1.
// Time:  O(V+E)  Space: O(V)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// seen maps node value -> original node pointer.
	// It doubles as the "visited" set during DFS and as our list of discovered nodes.
	seen := make(map[int]*Node)

	// clones maps node value -> cloned node pointer (structure only; neighbors filled in Pass 2).
	clones := make(map[int]*Node)

	// dfs records the original node and ensures a corresponding clone exists.
	var dfs func(*Node)
	dfs = func(n *Node) {
		// Guard against nil and revisits (handles cycles).
		if n == nil || seen[n.Val] != nil {
			return
		}
		seen[n.Val] = n

		// Allocate the clone the first time we encounter this value.
		if clones[n.Val] == nil {
			clones[n.Val] = &Node{Val: n.Val}
		}

		// Explore neighbors to discover the whole connected component.
		for _, nb := range n.Neighbors {
			dfs(nb)
		}
	}
	dfs(node)

	// Pass 2: for each discovered original node, append the corresponding cloned neighbors.
	for val, orig := range seen {
		c := clones[val]
		// Pre-size neighbor slice to avoid reallocations on dense nodes.
		c.Neighbors = make([]*Node, 0, len(orig.Neighbors))
		for _, nb := range orig.Neighbors {
			c.Neighbors = append(c.Neighbors, clones[nb.Val])
		}
	}

	// Return the clone that corresponds to the input reference (robust even if root != 1).
	return clones[node.Val]
}
