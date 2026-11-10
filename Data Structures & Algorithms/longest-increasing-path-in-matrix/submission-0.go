func longestIncreasingPath(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])
	nums := make([][]int, m*n)
	for i := range matrix {
		for j := range matrix[i] {
			nums[i*n+j] = []int{matrix[i][j], i, j}
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] > nums[j][0]
	})
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	dp := make([]int, m*n)
	res := 0
	for _, num := range nums {
		val, i, j := num[0], num[1], num[2]
		index := i*n+j
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni < 0 || nj < 0 || ni >= m || nj >= n || matrix[ni][nj] <= val {
				continue
			}
			dp[index] = max(dp[ni*n+nj], dp[index])
		}
		dp[index] += 1
		res = max(res, dp[index])
	}
	return res
}
