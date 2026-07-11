package find_the_minimum_area_to_cover_all_ones_i

func minimumArea(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	top, bottom := len(grid), -1
	left, right := len(grid[0]), -1

	oneExsit := false

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				oneExsit = true

				if i < top {
					top = i
				}
				if i > bottom {
					bottom = i
				}

				if j < left {
					left = j
				}
				if j > right {
					right = j
				}
			}
		}
	}

	if !oneExsit {
		return 0
	}

	width := right - left + 1
	heigth := bottom - top + 1

	return width * heigth
}
