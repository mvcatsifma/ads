package graphs

// Digraph represents a directed graph using adjacency lists
type Digraph struct {
	V   int     // Number of vertices (nodes) in the graph
	E   int     // Number of edges in the graph
	Adj [][]int // Adjacency lists: Adj[v] contains all vertices w where v->w exists
}

// NewDigraph creates a new directed graph with v vertices and no edges
func NewDigraph(v int) *Digraph {
	// Initialize adjacency lists for each vertex
	adj := make([][]int, v)
	for i := range v {
		adj[i] = make([]int, 0) // Empty slice for each vertex's neighbors
	}
	return &Digraph{
		V:   v,   // Set vertex count
		E:   0,   // Start with zero edges
		Adj: adj, // Assign adjacency lists
	}
}

// AddEdge adds a directed edge from vertex v to vertex w
func (g *Digraph) AddEdge(v, w int) {
	g.Adj[v] = append(g.Adj[v], w) // Add w to v's adjacency list
	g.E++                          // Increment edge count
}

// Adjacent returns all vertices adjacent to vertex v (outgoing edges from v)
func (g *Digraph) Adjacent(v int) []int {
	return g.Adj[v] // Return slice of neighbors
}

// Reverse creates a new digraph with all edges reversed
// If original has edge v->w, reversed graph has edge w->v
func (g *Digraph) Reverse() *Digraph {
	dg := NewDigraph(g.V) // Create new graph with same number of vertices

	// Iterate through all vertices and their outgoing edges
	for v := 0; v < g.V; v++ {
		for _, w := range g.Adjacent(v) {
			dg.AddEdge(w, v) // Reverse edge: w->v instead of v->w
		}
	}
	return dg
}
