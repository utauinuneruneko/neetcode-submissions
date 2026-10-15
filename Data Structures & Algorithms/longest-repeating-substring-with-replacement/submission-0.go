func characterReplacement(s string, k int) int {
	left, right := 0, 0
	res := 0
	win := make([]int, 26)
	curMax := 0
	for right < len(s) {
		cur := s[right]-'A'
		right++
		win[cur]++

		curMax = max(win[cur], curMax)
		for right-left - curMax > k {
			del := s[left] - 'A'
			win[del] --
			left++
		}
		res = max(res, right-left)
	}
	return res
}
