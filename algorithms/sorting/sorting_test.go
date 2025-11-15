package sorting

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
	"time"
)

func TestSorting(t *testing.T) {
	t.Run("Selection sorting", func(t *testing.T) {
		nums := []int{81, 87, 47, 59, 81, 18, 25, 40, 56, 0, 94, 96, 16, 67, 7, 3, 5}
		selection(&nums)
		expected := []int{0, 3, 5, 7, 16, 18, 25, 40, 47, 56, 59, 67, 81, 81, 87, 94, 96}

		if !slices.Equal(nums, expected) {
			t.Fatalf("got %v \n expected %v", nums, expected)
		}
	})
}

func generateRandomSlice(size int) []int {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = r.Intn(999)
	}

	return slice
}

func BenchmarkSorting(b *testing.B) {
	benchmarkSizes := []int{100, 1000, 10000}
	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("Selection Sorting Benchmarkd Size: %d", size), func(b *testing.B) {
			nums := generateRandomSlice(size)

			for i := 0; i < b.N; i++ {
				selection(&nums)
			}
		})
	}
}
