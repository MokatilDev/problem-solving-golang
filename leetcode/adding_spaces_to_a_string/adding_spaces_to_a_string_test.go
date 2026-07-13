package adding_spaces_to_a_string

import (
	"testing"
)

func TestAddSpaces(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		sapces   []int
		expected string
	}{
		{
			name:     "Example 1",
			s:        "LeetcodeHelpsMeLearn",
			sapces:   []int{8, 13, 15},
			expected: "Leetcode Helps Me Learn",
		},
		{
			name:     "Example 2",
			s:        "icodeinpython",
			sapces:   []int{1, 5, 7, 9},
			expected: "i code in py thon",
		},
		{
			name:     "Example 3",
			s:        "spacing",
			sapces:   []int{0, 1, 2, 3, 4, 5, 6},
			expected: " s p a c i n g",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := addSpaces(test.s, test.sapces)
			if result != test.expected {
				t.Fatalf("got %v, expected %v", result, test.expected)
			}
		})
	}
}

func BenchmarkAddSpaces(b *testing.B) {
	s := "LeetcodeHelpsMeLearn"
	spaces := []int{8, 13, 15}

	b.ResetTimer()

	for b.Loop() {
		addSpaces(s, spaces)
	}
}
