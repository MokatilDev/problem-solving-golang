package count_partitions_with_even_sum_difference

import "testing"

func TestCountPartitions(t *testing.T) {
	type testCase struct {
		input  []int
		output int
	}

	tests := []testCase{
		{
			input:  []int{10, 10, 3, 7, 6},
			output: 4,
		},
		{
			input:  []int{1, 2, 2},
			output: 0,
		},
		{
			input:  []int{2, 4, 6, 8},
			output: 3,
		},
	}

	for _, tc := range tests {
		t.Run("Test Inputs", func(t *testing.T) {
			got := countPartitions(tc.input)
			if got != tc.output {
				t.Errorf("got: %d, expected: %d", got, tc.output)
			}
		})
	}
}
