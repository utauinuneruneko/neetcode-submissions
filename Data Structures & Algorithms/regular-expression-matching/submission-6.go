func isMatch(s string, p string) bool {
    pat := make([]string, 0)
	for i := 0; i < len(p); {
		if i < len(p)-1 && p[i+1] == '*' {
			pat = append(pat, p[i:i+2])
			i += 2
		} else {
			pat = append(pat, p[i:i+1])
			i += 1
		}
	}
	m, n := len(s), len(pat)
	dp := make(map[int]map[int]bool)
	for i := 0; i <= m; i++ {
		dp[i] = make(map[int]bool)
	}
	dp[0][0] = true

	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if _, ok := dp[i][j]; ok {
			return dp[i][j]
		}

		if i == 0 {
			if len(pat[j-1]) == 2 {
				dp[i][j] = dfs(i, j-1)
				return dp[i][j]
			}
			return false
		}
		if j == 0 {
			return i == 0
		}

		if len(pat[j-1]) == 1 {
			if (pat[j-1][0]=='.' || pat[j-1][0] == s[i-1]) {
				dp[i][j] = dfs(i-1, j-1)
			} else {
				dp[i][j] = false
			}
		} else {
			if pat[j-1][0]=='.' || pat[j-1][0]==s[i-1] {
				dp[i][j] = dfs(i-1, j-1) || dfs(i-1, j) || dfs(i, j-1)
			} else {
				dp[i][j] = dfs(i, j-1)
			}
		}
		return dp[i][j]
	}

	return dfs(m, n)
}
