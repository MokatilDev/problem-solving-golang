package count_partitions_with_even_sum_difference

func countPartitions(nums []int) int {
	var k int

	for i := 0; i <= len(nums)-2; i++ {
		var sum1, sum2 int

		for j := 0; j <= i; j++ {
			sum1 += nums[j]
		}

		for j := i + 1; j < len(nums); j++ {
			sum2 += nums[j]
		}

		if (sum1-sum2)%2 == 0 {
			k++
		}
	}

	return k
}
