func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([]bool, m*n)
	dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	flag := false
	var dfs func(int, int, int)
	dfs = func(i, j int, index int) {
		if flag || index == len(word) {
			flag = true
			return
		}

		for _, dir := range dirs {
			ix, jx := i+dir[0], j+dir[1]
			if ix < 0 || jx < 0 || ix >= m || jx >= n || visited[ix*n+jx] || board[ix][jx] != word[index] {
				continue
			}
			visited[ix*n+jx] = true
			dfs(ix, jx, index+1)
			visited[ix*n+jx] = false
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visited[i*n+j] = true
				dfs(i, j, 1)
				if flag {
					return true
				}
				visited[i*n+j] = false
			}
		}
	}
	return false
}
