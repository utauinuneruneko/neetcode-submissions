func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
	res := make([]int, 0)
	win := make([]int, 0)
	left, right := 0, 0
	for right < k-1 {
		cur := nums[right]
		right++
		for len(win) > 0 && win[len(win)-1] < cur {
			win = win[:len(win)-1]
		}
		win = append(win, cur)
	}

	for right < n {
		cur := nums[right]
		right++
		for len(win) > 0 && win[len(win)-1] < cur {
			win = win[:len(win)-1]
		}
		win = append(win, cur)
		res = append(res, win[0])

		del := nums[left]
		left++
		if win[0] == del {
			win = win[1:]
		}
		
	}
	return res
}
