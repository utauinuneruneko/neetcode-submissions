func climbStairs(n int) int {
    dp_0, dp_1 := 1, 1
	for i := 2; i <= n; i++ {
		tmp := dp_0 + dp_1
		dp_0 = dp_1
		dp_1 = tmp
	}
	
	return dp_1
}
