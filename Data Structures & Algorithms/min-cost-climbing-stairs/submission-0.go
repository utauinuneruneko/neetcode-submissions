func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp_0, dp_1 := 0, 0
	for i := 2; i <= n; i++ {
		tmp := min(dp_0+cost[i-2], dp_1+cost[i-1])
		dp_0 = dp_1
		dp_1 = tmp
	}
	
	return dp_1
}
