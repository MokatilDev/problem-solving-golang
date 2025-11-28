package uglyNumber

import "testing"

func TestUgly(t *testing.T) {

	type testCase struct {
		name     string
		input    int
		expected bool
	}

	tests := []testCase{
		{
			name:     "Standard Ugly Number (6)",
			input:    6,
			expected: true,
		},
		{
			name:     "Not Ugly (14)",
			input:    14,
			expected: false,
		},
		{
			name:     "Base Case (1)",
			input:    1,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := isUgly(tc.input)
			if res != tc.expected {
				t.Errorf("For input %d, expected %v but got %v", tc.input, tc.expected, res)
			}
		})
	}
}
