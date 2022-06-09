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
}
