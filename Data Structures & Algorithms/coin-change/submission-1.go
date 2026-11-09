func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i + coin <= amount {
				dp[i+coin] = min(dp[i+coin], dp[i]+1)
			}	
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
