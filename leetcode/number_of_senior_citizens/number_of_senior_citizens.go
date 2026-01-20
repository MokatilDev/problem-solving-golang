package number_of_senior_citizens

import "strconv"

func countSeniors(details []string) int {
	var res int

	for _, v := range details {
		age_str := string(v[len(v)-4]) + string(v[len(v)-3])

		if age_int, _ := strconv.Atoi(age_str); age_int >= 60 {
			res++
		}
	}

	return res
}
