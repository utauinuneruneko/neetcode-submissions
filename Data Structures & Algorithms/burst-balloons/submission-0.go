func maxCoins(nums []int) int {
    n := len(nums)
	arr := append(append([]int{1}, nums...), 1)

	memo := make([][]int, n+2)
	for i := range memo {
		memo[i] = make([]int, n+2)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(left, right int) int {
		if right - left <= 1 {
			return 0
		}

		if memo[left][right] != -1 {
			return memo[left][right]
		}

		res := 0
		for k := left+1; k < right; k++ {
			total := dfs(left, k) + dfs(k, right) + arr[k]*arr[left]*arr[right]
			res = max(res, total)
		}
		memo[left][right] = res
		return res
	}
	return dfs(0, n+1)
}