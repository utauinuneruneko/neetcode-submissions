func maxArea(heights []int) int {
	left, right := 0, len(heights)-1
	res := 0
	for left < right {
		cur := min(heights[left], heights[right]) * (right-left)
		res = max(res, cur)
		if heights[left] <= heights[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
