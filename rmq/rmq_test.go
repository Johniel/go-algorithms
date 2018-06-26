package rmq

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	size := 100
	max := 1 << 30
	rmq := Init(func(a, b interface{}) bool { return a.(int) < b.(int) }, size, max)

	a := make([]int, size)
	for i := range a {
		x := rand.Intn(10000)
		a[i] = x
		rmq.Update(i, x)
	}

	for try := 0; try < 100; try++ {
		begin := rand.Intn(size)
		end := rand.Intn(size)
		if end < begin {
			begin, end = end, begin
		}
		mn := max
		for i := begin; i < end; i++ {
			if a[i] < mn {
				mn = a[i]
			}
		}
		assert.Equal(t, rmq.Query(begin, end), mn, "invalid query result")
	}
}
