package pivot

import "testing"

type Pivot struct {
	n     int
	pivot int
	got   int
}

func TestPivotInteger(t *testing.T) {
	cases := []Pivot{
		{n: 8, pivot: 6, got: pivotInteger(8)},
		{n: 1, pivot: 1, got: pivotInteger(1)},
		{n: 4, pivot: -1, got: pivotInteger(4)},
	}

	for i,case := range cases {
		
	}
}
