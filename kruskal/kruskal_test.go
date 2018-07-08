package kruskal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallelEdge(t *testing.T) {
	es := []Edge{
		Edge{src: 0, dst: 1, cost: 1},
		Edge{src: 1, dst: 0, cost: 2},
	}
	cost, mst := Kruskal(2, es)
	assert.Equal(t, cost, 1)
	assert.Len(t, mst, 1)
}

func TestCost(t *testing.T) {
	es := []Edge{
		Edge{src: 0, dst: 1, cost: 0},
		Edge{src: 1, dst: 2, cost: 1},
		Edge{src: 0, dst: 2, cost: 2},
		Edge{src: 8, dst: 9, cost: 0},
	}
	cost, mst := Kruskal(10, es)
	assert.Equal(t, cost, 1)
	assert.Len(t, mst, 3)
}
