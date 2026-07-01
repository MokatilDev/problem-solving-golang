package coin_change

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for i := range dp {
		for _, coin := range coins {
			if (i - coin) >= 0 {
				if dp[i-coin] == math.MaxInt {
					continue
				}

				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}

	return dp[amount]
}

// Depth Search approch
/* func coinChange(coins []int, amount int) int {
	var dfs func(remain int) int
	dfs = func(remain int) int {
		if remain == 0 {
			return 0
		}
		if remain < 0 {
			return math.MaxInt
		}

		numberOfCoins := math.MaxInt
		for _, coin := range coins {
			if coin > remain {
				continue
			}

			result := dfs(remain - coin)
			if result != math.MaxInt {
				numberOfCoins = min(numberOfCoins, 1+result)
			}
		}

		return numberOfCoins
	}

	res := dfs(amount)
	if res == math.MaxInt {
		return -1
	}

	return res
}*/
