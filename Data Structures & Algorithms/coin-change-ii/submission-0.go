func change(amount int, coins []int) int {
    dp := make([]int, amount+1)

	for _, coin := range coins {
		dp[0] = 1
		for j := coin; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coin]
		}
	}

	return dp[amount]
}
