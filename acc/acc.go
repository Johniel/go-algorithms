package acc

type Acc struct {
	sum []int64
}

func NewAcc(values []int64) *Acc {
	acc := &Acc{}
	acc.sum = append(acc.sum, 0)
	var sum int64
	for _, v := range values {
		sum += v
		acc.sum = append(acc.sum, sum)
	}
	return acc
}

func (a *Acc) Query(begin, end int) int64 {
	return a.sum[end] - a.sum[begin]
}
