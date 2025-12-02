package addBinary

import (
	"strings"
	"testing"
)

func TestAddBinary(t *testing.T) {
	type TestCases struct {
		a        string
		b        string
		expected string
	}

	tests := []TestCases{
		{
			a:        "1",
			b:        "11",
			expected: "100",
		},
		{
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
	}

	for _, tc := range tests {
		got := addBinary(tc.a, tc.b)

		if strings.Compare(got, tc.expected) != 0 {
			t.Errorf("got %s expected %s", got, tc.expected)
		}
	}
}
