package kruskal

import (
	"sort"

	"github.com/johniel/go-algorithms/union_find"
)

// Edge struct is weighted graph edge.
type Edge struct {
	src  int
	dst  int
	cost int
}

// Kruskal is MST algorithm
func Kruskal(v int, es []Edge) (int, []Edge) {
	sort.Slice(es, func(i, j int) bool {
		return es[i].cost < es[j].cost
	})

	cost := 0
	mst := make([]Edge, 0)
	u := UnionFind.New(v)
	for _, e := range es {
		if !u.IsSameSet(e.src, e.dst) {
			u.Union(e.src, e.dst)
			mst = append(mst, e)
			cost += e.cost
		}
	}

	return cost, mst
}
