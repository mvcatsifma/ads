package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigraph_Reverse(t *testing.T) {
	// 0 -> 1 -> 2
	g := NewDigraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)

	// 0 <- 1 <- 2
	rg := g.Reverse()
	assert.Len(t, rg.Adjacent(0), 0)
	assert.Len(t, rg.Adjacent(1), 1)
	assert.Equal(t, 0, rg.Adjacent(1)[0])
	assert.Len(t, rg.Adjacent(2), 1)
	assert.Equal(t, 1, rg.Adjacent(2)[0])
}
