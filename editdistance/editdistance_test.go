package editdistance

import (
	"testing"
)

func Test(t *testing.T) {
	if x := EditDistance("a", "a"); x != 0 {
		t.Errorf("actual %v, expected %v", x, 0)
	}
	if x := EditDistance("a", "b"); x != 1 {
		t.Errorf("actual %v, expected %v", x, 1)
	}
	if x := EditDistance("a", "bb"); x != 2 {
		t.Errorf("actual %v, expected %v", x, 2)
	}
	if x := EditDistance("aa", "b"); x != 2 {
		t.Errorf("actual %v, expected %v", x, 2)
	}
	if x := EditDistance("axa", "bxb"); x != 2 {
		t.Errorf("actual %v, expected %v", x, 2)
	}
	if x := EditDistance("axa", "bxa"); x != 1 {
		t.Errorf("actual %v, expected %v", x, 1)
	}
}
