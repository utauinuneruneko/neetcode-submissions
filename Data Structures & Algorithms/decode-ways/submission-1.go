func numDecodings(s string) int {
	memo := make(map[string]int)
	memo[""] = 1
 	var dfs func(string) int
	dfs = func(str string) int {
		if _, ok := memo[str]; ok {
			return memo[str]
		}

		if str[0] == '0' {
			memo[str] = 0
			return memo[str]
		}

		count := dfs(str[1:])
		if len(str) > 1 && (str[0] == '1' || str[0] <= '2' && str[1] <= '6') {
			count += dfs(str[2:])
		}
		memo[str] = count

		return memo[str]
	}
	return dfs(s)
}
