package p4Graphs

// DepthFirstSearch performs DFS and tracks reachable vertices
type DepthFirstSearch struct {
	marked []bool
	count  int
}

// NewDepthFirstSearch creates a new DFS instance and runs the search
func NewDepthFirstSearch(g *Graph, source int) *DepthFirstSearch {
	dfs := &DepthFirstSearch{
		marked: make([]bool, g.V),
		count:  0,
	}
	dfs.search(g, source)
	return dfs
}

// search performs the actual DFS traversal
func (dfs *DepthFirstSearch) search(g *Graph, v int) {
	dfs.marked[v] = true
	dfs.count++

	for _, w := range g.Adjacent(v) {
		if !dfs.marked[w] {
			dfs.search(g, w)
		}
	}
}

// IsMarked returns true if vertex v is reachable from the source
func (dfs *DepthFirstSearch) IsMarked(v int) bool {
	return dfs.marked[v]
}

// Count returns the number of vertices reachable from the source
func (dfs *DepthFirstSearch) Count() int {
	return dfs.count
}

// ReachableVertices returns all vertices reachable from source via DFS
func ReachableVertices(g *Graph, source int) ([]bool, int) {
	marked := make([]bool, g.V)
	reachable := make([]bool, g.V)

	var dfs func(int)
	dfs = func(v int) {
		marked[v] = true
		reachable[v] = true

		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				dfs(w)
			}
		}
	}

	dfs(source)

	count := 0
	for _, v := range reachable {
		if v {
			count++
		}
	}

	return reachable, count
}
