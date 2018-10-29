package rmq

// Less is a RMQ comparing function type.
type Less func(a, b interface{}) bool

// RMQ is a Range Minimum Query.
type RMQ struct {
	data []interface{}
	n    int
	less Less
	ini  interface{}
}

// New returns an initialized RMQ.
func New(less Less, size int, ini interface{}) *RMQ {
	rmq := new(RMQ)
	rmq.n = 1
	for rmq.n < size {
		rmq.n *= 2
	}
	rmq.ini = ini
	rmq.less = less
	rmq.data = make([]interface{}, rmq.n*2+1)
	for i := range rmq.data {
		rmq.data[i] = ini
	}
	return rmq
}

// Update updates i-th value with e.
func (rmq *RMQ) Update(i int, e interface{}) {
	i += rmq.n - 1
	rmq.data[i] = e
	for 0 < i {
		i = (i - 1) / 2
		j := i*2 + 1
		k := i*2 + 2
		if rmq.less(rmq.data[j], rmq.data[k]) {
			rmq.data[i] = rmq.data[j]
		} else {
			rmq.data[i] = rmq.data[k]
		}
	}
}

func (rmq *RMQ) query(begin, end, k, l, r int) interface{} {
	if r <= begin || end <= l {
		return rmq.ini
	}
	if begin <= l && r <= end {
		return rmq.data[k]
	}
	v1 := rmq.query(begin, end, k*2+1, l, (l+r)/2)
	v2 := rmq.query(begin, end, k*2+2, (l+r)/2, r)
	if rmq.less(v1, v2) {
		return v1
	} else {
		return v2
	}
}

// Query returns minimum value [begin, end)
func (rmq *RMQ) Query(begin, end int) interface{} {
	return rmq.query(begin, end, 0, 0, rmq.n)
}
