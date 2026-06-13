package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	res := 0

	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			res = max(res, profit)
		} else {
			l = r
		}

		r++
	}

	return res
}
