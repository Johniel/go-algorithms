package UnionFind

// UnionFind is disjoint set data structure
type UnionFind struct {
	p []int
	r []int
}

// New returns new UnionFind pointer with given size.
func New(size int) *UnionFind {
	uf := new(UnionFind)
	uf.p = make([]int, size)
	uf.r = make([]int, size)
	for i := range uf.p {
		uf.p[i] = i
		uf.r[i] = 1
	}
	return uf
}

// Union
func (uf *UnionFind) Union(a, b int) {
	if uf.r[a] > uf.r[b] {
		uf.p[b] = a
	} else {
		uf.p[a] = b
		if uf.r[a] == uf.r[b] {
			uf.r[a]++
		}
	}
}

// Find return root element
func (uf *UnionFind) Find(x int) int {
	if uf.p[x] != x {
		uf.p[x] = uf.Find(uf.p[x])
	}
	return uf.p[x]
}

// IsSameSet return true if given two elements are included in a single set.
func (uf *UnionFind) IsSameSet(a, b int) bool {
	a = uf.Find(a)
	b = uf.Find(b)
	return a == b
}
