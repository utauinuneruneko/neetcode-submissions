func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1
	for _, val := range piles {
		right = max(right, val)
	}

	var costTime func(int) int
	costTime = func(rate int) int {
		res := 0
		for _, val := range piles {
			res += val/rate
			if val/rate*rate < val {
				res += 1
			}
		}
		return res
	}

	for left <= right {
		mid := (right-left)/2+left
		cost := costTime(mid)
		if cost <= h {
			right = mid-1
		} else if cost > h {
			left = mid+1
		} 
	}
	return right+1
}
