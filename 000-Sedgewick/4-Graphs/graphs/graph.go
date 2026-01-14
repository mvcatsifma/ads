package graphs

type Graph struct {
	V   int     // number of vertices
	E   int     // number of edges
	Adj [][]int // adjacency lists
}

func (g *Graph) AddEdge(v, w int) {
	g.Adj[v] = append(g.Adj[v], w)
	g.Adj[w] = append(g.Adj[w], v)
	g.E++
}

func (g *Graph) Adjacent(v int) []int {
	return g.Adj[v]
}

func NewGraph(v int) *Graph {
	adj := make([][]int, v)
	for i := range v {
		adj[i] = make([]int, 0)
	}
	return &Graph{
		V:   v,
		E:   0,
		Adj: adj,
	}
}
