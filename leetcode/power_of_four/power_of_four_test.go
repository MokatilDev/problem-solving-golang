package power_of_four

import "testing"

func TestCountPartitions(t *testing.T) {
	type testCase struct {
		input  int
		output bool
	}

	tests := []testCase{
		{
			input:  16,
			output: true,
		},
		{
			input:  8,
			output: false,
		},
		{
			input:  1,
			output: true,
		},
	}

	for _, tc := range tests {
		t.Run("Test Inputs", func(t *testing.T) {
			got := isPowerOfFour(tc.input)
			if got != tc.output {
				t.Errorf("got: %v, expected: %v", got, tc.output)
			}
		})
	}
}
