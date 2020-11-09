package disjointset

import (
	"testing"
)

func TestCycle(t *testing.T) {
	edges := [][]int{
		{0, 1}, {1, 2}, {1, 3}, {3, 4}, {2, 5},
	}

	djset := New(6)

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		if isCycle, _ := djset.Union(a, b); isCycle == 0 {
			t.Errorf("expected not cycle: %v - %v", a, b)
		}
	}
}
