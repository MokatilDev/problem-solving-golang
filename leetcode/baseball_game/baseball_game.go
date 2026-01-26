package base_ballgame

import (
	"strconv"
)

func calPoints(operations []string) int {
	var res int
	var arr []int

	for i := range operations {
		switch operations[i] {
		case "D":
			if len(arr) >= 1 {
				arr = append(arr, 2*arr[len(arr)-1])
			}
		case "+":
			if len(arr) >= 2 {
				arr = append(arr, arr[len(arr)-2]+arr[len(arr)-1])
			}
		case "C":
			arr = arr[:len(arr)-1]
		default:
			num, _ := strconv.Atoi(operations[i])
			arr = append(arr, num)
		}
	}

	for _, num := range arr {
		res += num
	}

	return res
}
