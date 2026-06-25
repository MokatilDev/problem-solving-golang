package koko_eating_bananas

import "slices"

func minEatingSpeed(piles []int, h int) int {
	slices.Sort(piles)

	l, r := 1, piles[len(piles)-1]
	res := r

	for l <= r {
		k := (l + r) / 2
		hoursSpent := 0

		for _, p := range piles {
			hoursSpent += (p + k - 1) / k
		}

		if hoursSpent <= h {
			if k < res {
				res = k
			}

			r = k - 1
		} else {
			l = k + 1
		}
	}

	return res
}
