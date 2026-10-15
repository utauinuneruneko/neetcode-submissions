func isValidSudoku(board [][]byte) bool {
	m := len(board)
	col, row, squa := make([][]bool, m), make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		col[i], row[i], squa[i] = make([]bool, 10), make([]bool, 10), make([]bool, 10)
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				if col[i][num] || row[j][num] || squa[(i/3)*3+j/3][num] {
					return false
				}
				col[i][num] = true
				row[j][num] = true
				squa[(i/3)*3+j/3][num] = true
			}
		}
	}
	
	return true
}
