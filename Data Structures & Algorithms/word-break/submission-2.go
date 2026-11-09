func wordBreak(s string, wordDict []string) bool {
    dict := make(map[string]struct{})
	for _, word := range wordDict {
		dict[word] = struct{}{}
	}

	n := len(s)

	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; !dp[i] && j < i; j++ {
			if dp[j] {
				if _, ok := dict[s[j:i]]; ok {
					dp[i] = true
				}
			}
		}
	}

	return dp[n]
}
