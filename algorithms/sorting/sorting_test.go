package sorting

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
)

func TestSorting(t *testing.T) {
	t.Run("Selection sorting", func(t *testing.T) {
		nums := []int{81, 87, 47, 59, 81, 18, 25, 40, 56, 0, 94, 96, 16, 67, 7, 3, 5}
		selectionSort(&nums)
		expected := []int{0, 3, 5, 7, 16, 18, 25, 40, 47, 56, 59, 67, 81, 81, 87, 94, 96}

		if !slices.Equal(nums, expected) {
			t.Fatalf("got %v \n expected %v", nums, expected)
		}
	})

	t.Run("Insertion sorting", func(t *testing.T) {
		nums := []int{81, 87, 47, 59, 81, 18, 25, 40, 56, 0, 94, 96, 16, 67, 7, 3, 5}
		insertionSort(&nums)
		expected := []int{0, 3, 5, 7, 16, 18, 25, 40, 47, 56, 59, 67, 81, 81, 87, 94, 96}

		if !slices.Equal(nums, expected) {
			t.Fatalf("got %v \n expected %v", nums, expected)
		}
	})

	t.Run("Bubble sorting", func(t *testing.T) {
		nums := []int{81, 87, 47, 59, 81, 18, 25, 40, 56, 0, 94, 96, 16, 67, 7, 3, 5}
		bubbleSort(&nums)
		expected := []int{0, 3, 5, 7, 16, 18, 25, 40, 47, 56, 59, 67, 81, 81, 87, 94, 96}

		if !slices.Equal(nums, expected) {
			t.Fatalf("got %v \n expected %v", nums, expected)
		}
	})
}

func generateRandomSlice(size int) []int {
	nums := make([]int, size)

	for i := range size {
		nums[i] = rand.IntN(size)
	}

	return nums
}

func BenchmarkSorting(b *testing.B) {
	benchmarkSizes := []int{100, 1000, 10000}
	for _, size := range benchmarkSizes {
		b.StopTimer()
		nums := generateRandomSlice(size)
		b.StartTimer()

		b.Run(fmt.Sprintf("Selection Sor Benchmarkd Size: %d", size), func(b *testing.B) {
			for b.Loop() {
				selectionSort(&nums)
			}
		})

		b.Run(fmt.Sprintf("Insertion Sort Benchmarkd Size: %d", size), func(b *testing.B) {
			for b.Loop() {
				insertionSort(&nums)
			}
		})

		b.Run(fmt.Sprintf("Bubble Sort Benchmarkd Size: %d", size), func(b *testing.B) {
			for b.Loop() {
				bubbleSort(&nums)
			}
		})
	}
}
