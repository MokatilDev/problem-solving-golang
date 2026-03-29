package kids_with_the_greatest_number_of_candies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	greatest := 0

	for _, v := range candies {
		if greatest < v {
			greatest = v
		}
	}

	for i := range candies {
		v := candies[i] + extraCandies

		if v >= greatest {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	return res
}
