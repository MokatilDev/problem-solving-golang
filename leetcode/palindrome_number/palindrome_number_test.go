package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	type testCase struct {
		name     string
		input    int
		expected bool
	}

	tests := []testCase{
		{
			name:     "Positive palindrome (odd length)",
			input:    121,
			expected: true,
		},
		{
			name:     "Positive palindrome (even length)",
			input:    1221,
			expected: true,
		},
		{
			name:     "Negative number",
			input:    -121,
			expected: false,
		},
		{
			name:     "Non-palindrome positive",
			input:    10,
			expected: false,
		},
		{
			name:     "Single digit zero",
			input:    0,
			expected: true,
		},
		{
			name:     "Large palindrome",
			input:    123454321,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

		})
	}

}
