package search

import "leetcode/000-Sedgewick/4-Graphs/graphs"

func ReachableVerticesDirected(g *graphs.Digraph, source int) ([]bool, int) {
	if source < 0 || source >= g.V {
		return make([]bool, g.V), 0
	}

	marked := make([]bool, g.V)

	var dfs func(v int)
	dfs = func(v int) {
		marked[v] = true
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				dfs(w)
			}
		}
	}
	dfs(source)

	count := 0
	for _, m := range marked {
		if m {
			count++
		}
	}

	return marked, count
}
