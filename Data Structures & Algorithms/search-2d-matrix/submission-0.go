func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (right-left)/2+left
		i, j := mid/n, mid%n
		if matrix[i][j] < target {
			left = mid+1
		} else if matrix[i][j] > target {
			right = mid-1
		} else if matrix[i][j] == target{
			return true
		}
	}
	return false
}
