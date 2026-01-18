package valid_sudoku

import "strconv"

func isValidSudoku(board [][]byte) bool {
	res := true

	exist := make(map[string]bool)

	for i := range 9 {
		for j := range 9 {
			char := board[i][j]

			if char != '.' {
				row_key := "row" + strconv.Itoa(i)
				colum_key := "colum" + strconv.Itoa(j)
				box_key := "box" + strconv.Itoa(i/3) + strconv.Itoa(j/3) + string(char)

				if exist[row_key] || exist[colum_key] || exist[box_key] {
					res = false
					break
				}

				exist[row_key], exist[colum_key], exist[box_key] = true, true, true
			}
		}
	}

	return res
}
