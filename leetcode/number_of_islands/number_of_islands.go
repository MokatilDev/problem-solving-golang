package number_of_islands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	n, m := len(grid), len(grid[0])
	islands := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				islands++

				bfs(grid, i, j)
			}
		}
	}

	return islands
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func bfs(grid [][]byte, i, j int) {
	queue := [][2]int{}
	queue = append(queue, [2]int{i, j})

	grid[i][j] = '0'
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, direction := range directions {
			a := current[0] + direction[0]
			b := current[1] + direction[1]

			if a >= 0 && a < len(grid) && b >= 0 && b < len(grid[0]) && grid[a][b] == '1' {
				queue = append(queue, [2]int{a, b})
				grid[a][b] = '0'
			}
		}
	}
}
