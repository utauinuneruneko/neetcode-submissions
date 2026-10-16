func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n+2), make([]int, n+2)
	heights = append([]int{0}, append(heights, 0)...)
	stack := []int{n+1}
	for i := n; i >= 1; i-- {
		
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = i+1
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{0}
	for i := 1; i <= n; i++ {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = i-1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	res := 0
	for i := 1; i <= n; i++ {
		res = max(res, (right[i] - left[i] - 1)*heights[i])
	}
	return res
}
