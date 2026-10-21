func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	path := make([][]byte, n)
	for i := range path {
		path[i] = make([]byte, n)
		for j := range path[i] {
			path[i][j] = '.'
		}
	}
	col, row, diag1, diag2 := make([]bool, n), make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)

	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			str := make([]string, n)
			for j, v := range path {
				str[j] = string(v)
			}
			res = append(res, append([]string{}, str...))
			return
		}

		for j := 0; j < n; j++ {
			if col[i] || row[j] || diag1[i-j+(n-1)] || diag2[i+j] {
				continue
			}
			col[i] = true
			row[j] = true 
			diag1[i-j+(n-1)] = true
			diag2[i+j] = true
			path[i][j] = 'Q'
			backtrack(i+1)
			path[i][j] = '.'
			col[i] = false
			row[j] = false 
			diag1[i-j+(n-1)] = false
			diag2[i+j] = false
		}
	}
	backtrack(0)
	return res
}
