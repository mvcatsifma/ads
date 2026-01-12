package graphs

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

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

func NewGraphFromReader(reader io.Reader) (*Graph, error) {
	scanner := bufio.NewScanner(reader)

	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read number of vertices")
	}
	V, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return nil, fmt.Errorf("invalid number of vertices: %v", err)
	}

	g := NewGraph(V)

	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read number of edges")
	}
	E, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return nil, fmt.Errorf("invalid number of edges: %v", err)
	}

	for i := 0; i < E; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("failed to read edge %d", i+1)
		}

		parts := strings.Fields(scanner.Text())
		if len(parts) > 2 {
			return nil, fmt.Errorf("invalid edge format for edge %d (expected 2 vertices)", i+1)
		}

		v, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid vertex in edge %d: %v", i+1, err)
		}

		w, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid vertex in edge %d: %v", i+1, err)
		}

		g.AddEdge(v, w)
	}

	return g, scanner.Err()
}
