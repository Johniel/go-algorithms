package acc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var values []int64
	for i := 1; i <= 10; i++ {
		values = append(values, int64(i))
	}

	a := NewAcc(values)
	for i := 0; i < len(values); i++ {
		for j := i; j < len(values); j++ {
			var sum int64
			for k := i; k < j+1; k++ {
				sum += values[k]
			}
			assert.Equal(t, a.Query(i, j+1), sum)
		}
	}
}
