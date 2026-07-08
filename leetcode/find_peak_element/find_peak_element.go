package find_peak_element

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if mid > 0 && nums[mid] < nums[mid-1] {
			r = mid - 1
		} else if mid < len(nums)-1 && nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			return nums[mid]
		}
	}

	return -1
}
