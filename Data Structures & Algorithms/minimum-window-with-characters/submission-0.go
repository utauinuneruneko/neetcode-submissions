func minWindow(s string, t string) string {
    needs := make([]int, 52)
	cnt := 0
	var getIndex func(byte) int
	getIndex = func(b byte) int {
		index := 0
		if b >= 'a' && b <= 'z' {
			index = int(b-'a')
		} else {
			index = int(b - 'A') + 26
		}
		return index
	}
	for i := 0; i < len(t); i++ {
		index := getIndex(t[i])
		if needs[index] == 0 {
			cnt++
		}
		needs[index]++
	}
	res := ""
	winCnt := 0
	win := make([]int, 52)
	left, right := 0, 0
	for right < len(s) {
		cur := getIndex(s[right])
		right++
		win[cur]++
		if win[cur] == needs[cur] {
			winCnt++
		}

		for winCnt == cnt {
			if len(res) == 0 || right-left < len(res) {
				res = s[left:right]
			}
			del := getIndex(s[left]) 
			if win[del] == needs[del] {
				winCnt--
			}
			win[del]--
			left++
		}
	}
	return res
}
