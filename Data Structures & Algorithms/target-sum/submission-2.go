func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}

	 
	p := (sum+target)/2
	if (sum+target)%2 != 0 || p < 0{
		return 0
	}
	dp := make([]int, p+1)
	dp[0] = 1
	for _, num := range nums {
		for j := p; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
		}
	}

	return dp[p]
}
