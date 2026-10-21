func generateParenthesis(n int) []string {
	res := make([]string, 0)
	path := make([]byte, 0)
	var backtrace func(int, int)
	backtrace = func(left, right int) {
		if left == n && right == n {
			res = append(res, string(path))
			return
		} else if left > n || right > n {
			return
		} else if left < right {
			return
		}

		if left > right {
			path = append(path, ')')
			backtrace(left, right+1)
			path = path[:len(path)-1]
		} 
		path = append(path, '(')
		backtrace(left+1, right)
		path = path[:len(path)-1]
	}
	backtrace(0, 0)
	return res
}
