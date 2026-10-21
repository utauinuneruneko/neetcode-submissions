func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	var isPalindrome func(string) bool
	isPalindrome = func(str string) bool {
		l, r := 0, len(str)-1
		for l < r {
			if str[l] != str[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	var backtrack func(int)
	backtrack = func(start int) {
		if start == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := start; i < len(s); i++ {
			str := s[start:i+1]
			if !isPalindrome(str) {
				continue
			}
			path = append(path, str)
			backtrack(i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
