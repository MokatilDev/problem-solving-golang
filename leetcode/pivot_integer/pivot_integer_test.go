package pivot

import "testing"

type testCase struct {
	n     int
	pivot int
}

func TestPivotInteger(t *testing.T) {
	testCases := []testCase{
		{n: 8, pivot: 6},
		{n: 1, pivot: 1},
		{n: 4, pivot: -1},
	}

	for _, tc := range testCases {
		t.Run("The pivot", func(t *testing.T) {
			n := tc.n
			got := pivotInteger(n)
			expected := tc.pivot
			if tc.pivot != pivotInteger(tc.n) {
				t.Errorf("for n=%d expected: %d  got: %d\n", n, expected, got)
			}
		})
	}
}
