func canPartition(nums []int) bool {
    sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 != 0 {
		return false
	}
	target := sum/2
	
	dp := make([]bool, target+1)
	for i := 0; i < len(nums); i++ {
		dp[0] = true
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[target]
}
