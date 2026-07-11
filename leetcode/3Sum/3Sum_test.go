package three_sum

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "Example 3",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := threeSum(test.nums)
			if !reflect.DeepEqual(result, test.expected) {
				t.Fatalf("got %v, expected %v", result, test.expected)
			}
		})
	}
}

func BenchmarkThreeSum(b *testing.B) {
	benchmarkSizes := []int{100, 1000}

	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			nums := make([]int, size)
			for i := range nums {
				nums[i] = rand.Intn(2000) - 1000
			}

			b.ResetTimer()
			for b.Loop() {
				threeSum(nums)
			}
		})
	}
}
