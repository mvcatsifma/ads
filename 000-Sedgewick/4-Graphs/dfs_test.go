package p4Graphs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const graphDisconnected = `8
6
0 1
1 2
2 0
3 4
4 5
5 3`

func TestDepthFirstSearch(t *testing.T) {
	graph, err := NewGraphFromReader(strings.NewReader(graphDisconnected))
	if err != nil {
		t.Fatalf("Failed to created Graph: %v", err)
	}

	dfs := NewDepthFirstSearch(graph, 0)

	assert.Equal(t, dfs.count, 3)
	assert.True(t, dfs.marked[0])
	assert.False(t, dfs.marked[3])
}

func TestReachableVertices(t *testing.T) {
	graph, err := NewGraphFromReader(strings.NewReader(graphDisconnected))
	if err != nil {
		t.Fatalf("Failed to created Graph: %v", err)
	}

	reachable, count := ReachableVertices(graph, 0)

	assert.Equal(t, count, 3)
	assert.True(t, reachable[0])
	assert.False(t, reachable[3])
}
