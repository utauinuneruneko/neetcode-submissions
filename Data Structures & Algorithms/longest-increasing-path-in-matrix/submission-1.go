func longestIncreasingPath(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	res := 0
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni < 0 || nj < 0 || ni >= m || nj >= n || matrix[ni][nj] >= matrix[i][j] {
				continue
			}
			memo[i][j] = max(memo[i][j], dfs(ni, nj))
		}
		memo[i][j] += 1
		return memo[i][j]
	}
	for i := range matrix {
		for j := range matrix[i] {
			res = max(res, dfs(i, j))
		}
	}
	return res
}
