func trap(height []int) int {
	left, right := 0, len(height)-1
	res := 0 
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax <= rightMax {
			// caculate left
			res += leftMax-height[left]
			left++
		} else {
			res += rightMax-height[right]
			right--
		}
	}
	return res
}
