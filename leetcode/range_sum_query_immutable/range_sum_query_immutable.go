package range_sum_query_immutable

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	return NumArray{arr: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		res += this.arr[i]
	}
	return res
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
