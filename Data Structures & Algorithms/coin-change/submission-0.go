func coinChange(coins []int, amount int) int {
    memo := make(map[int]int)
	var dfs func(int) int
	dfs = func(val int) int {
		if val == 0 {
			return 0
		}
		if _, ok := memo[val]; ok {
			return memo[val]
		}
		res := math.MaxInt32
		for _, coin := range coins {
			if val < coin {
				continue
			}
			res = min(dfs(val-coin)+1, res)
		}
		memo[val] = res
		return memo[val]
	}
	res := dfs(amount)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
