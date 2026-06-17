package largest_rectangle_in_histogram

func largestRectangleArea(heights []int) int {
	var res int
	stack := make([]int, 0)

	for i := 0; i <= len(heights)+1; i++ {
		h := 0
		if i < len(heights) {
			h = heights[i]
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := heights[top]
			width := 0

			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if res < area {
				res = area
			}
		}

		stack = append(stack, i)
	}

	return res
}
