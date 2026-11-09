func wordBreak(s string, wordDict []string) bool {
    dict := make(map[string]struct{})
	for _, word := range wordDict {
		dict[word] = struct{}{}
	}

	n := len(s)

	memo := make(map[int]bool)
	memo[n] = true

	var dfs func(int) bool
	dfs = func(start int) bool {
		if start == n {
			return true
		}

		if _, ok := memo[start]; ok {
			return memo[start]
		}

		res := false
		for i := start; !res && i < len(s); i++ {
			if _, ok := dict[s[start:i+1]]; ok {
				res = res || dfs(i+1)
			}
		}
		memo[start] = res
		return res
	}
	dfs(0)
	return memo[0]
}
