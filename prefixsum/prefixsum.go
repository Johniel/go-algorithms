package prefixsum

type PrefixSum struct {
	sum []int64
}

func NewPrefixSum(values []int64) *PrefixSum {
	acc := &PrefixSum{}
	acc.sum = append(acc.sum, 0)
	var sum int64
	for _, v := range values {
		sum += v
		acc.sum = append(acc.sum, sum)
	}
	return acc
}

func (a *PrefixSum) Query(begin, end int) int64 {
	return a.sum[end] - a.sum[begin]
}
