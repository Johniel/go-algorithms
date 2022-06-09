package mincostmaxflow

import (
	"container/heap"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

type heapEntry struct {
	node int
	cost int64
}

type minHeap []*heapEntry

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].cost < h[j].cost }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*heapEntry)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type edge struct {
	src        int
	dst        int
	cap        int64
	flow       int64
	cost       int64
	rev        int
	isResidual bool
}

type Edge struct {
	Src  int
	Dst  int
	Cap  int64
	Flow int64
	Cost int64
}

func (e *edge) residue() int64 {
	return e.cap - e.flow
}

type path struct {
	node int
	idx  int
}

type MinCostMaxFlow struct {
	n         int
	dist      []int64
	path      []*path
	potential []int64
	g         [][]*edge
}

func NewMinCostMaxFlow(n int) *MinCostMaxFlow {
	return &MinCostMaxFlow{
		n:         n,
		dist:      make([]int64, n),
		path:      make([]*path, n),
		potential: make([]int64, n),
		g:         make([][]*edge, n),
	}
}

func (f *MinCostMaxFlow) AddEdge(src int, dst int, cost, cap int64) {
	f.g[src] = append(f.g[src], &edge{src: src, dst: dst, cap: cap, flow: 0, cost: +cost, rev: len(f.g[dst]), isResidual: false})
	f.g[dst] = append(f.g[dst], &edge{src: dst, dst: src, cap: cap, flow: cap, cost: -cost, rev: len(f.g[src]) - 1, isResidual: true})
}

func (f *MinCostMaxFlow) sssp(src int, dst int) bool {
	const inf int64 = (1 << 60)
	for i := 0; i < f.n; i++ {
		f.dist[i] = inf
	}
	f.dist[src] = 0

	f.path[src] = &path{node: src, idx: -1}

	q := &minHeap{}
	heap.Push(q, &heapEntry{node: src, cost: 0})
	for q.Len() != 0 {
		m := heap.Pop(q).(*heapEntry)
		if f.dist[m.node] != m.cost {
			continue
		}
		if m.node == dst {
			break
		}
		for i, e := range f.g[m.node] {
			if e.residue() <= 0 {
				continue
			}
			rcost := e.cost + (f.potential[e.src] - f.potential[e.dst])
			if f.dist[e.dst] > rcost+f.dist[e.src] {
				f.dist[e.dst] = rcost + f.dist[e.src]
				f.path[e.dst] = &path{node: e.src, idx: i}
				heap.Push(q, &heapEntry{node: e.dst, cost: f.dist[e.dst]})
			}
		}
	}

	return f.dist[dst] != inf
}

func (f *MinCostMaxFlow) Run(src int, snk int, req int64) (cost int64, flowed int64) {
	for i := 0; i < f.n; i++ {
		f.potential[i] = 0
	}

	for 0 < req && f.sssp(src, snk) {
		for i := 0; i < f.n; i++ {
			f.potential[i] += f.dist[i]
		}
		mn := req
		for i := snk; i != f.path[i].node; i = f.path[i].node {
			v := f.path[i].node
			e := f.path[i].idx
			mn = min(mn, f.g[v][e].residue())
		}
		for i := snk; i != f.path[i].node; i = f.path[i].node {
			v := f.path[i].node
			e := f.path[i].idx
			cost += mn * f.g[v][e].cost
			f.g[v][e].flow += mn
			f.g[f.g[v][e].dst][f.g[v][e].rev].flow -= mn
		}
		req -= mn
		flowed += mn
	}

	return cost, flowed
}

func (f *MinCostMaxFlow) Edges() []*Edge {
	es := []*Edge{}
	for _, a := range f.g {
		for _, e := range a {
			if !e.isResidual {
				es = append(es, &Edge{
					Src:  e.src,
					Dst:  e.dst,
					Cap:  e.cap,
					Flow: e.flow,
					Cost: e.cost,
				})
			}
		}
	}
	return es
}
