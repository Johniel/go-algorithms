package mincostmaxflow

import (
	"testing"
)

func TestMinCostMaxFlow(t *testing.T) {
	f := NewMinCostMaxFlow(10)

	f.AddEdge(0, 1, 10, 10)
	f.AddEdge(0, 1, 5, 1)
	f.AddEdge(1, 2, 8, 3)
	cost, flowed := f.Run(0, 2, 100)

	if expected := int64(5*1 + 8*3 + 10*2); cost != expected {
		t.Errorf("actual %v, expected %v", cost, expected)
	}
	if expected := int64(3); flowed != expected {
		t.Errorf("actual %v, expected %v", flowed, expected)
	}

	es := f.Edges()
	if expected := 3; len(es) != expected {
		t.Errorf("actual %v, expected %v", len(es), expected)
	}
	expected := []*Edge{
		{
			Src:  0,
			Dst:  1,
			Cap:  10,
			Cost: 10,
			Flow: 2,
		},
		{
			Src:  0,
			Dst:  1,
			Cap:  1,
			Cost: 5,
			Flow: 1,
		},
		{
			Src:  1,
			Dst:  2,
			Cap:  3,
			Cost: 8,
			Flow: 3,
		},
	}
	for i := 0; i < 3; i++ {
		if *es[i] != *expected[i] {
			t.Errorf("actual %v, expected %v", es[i], expected[i])
		}
	}
}
