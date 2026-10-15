func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := nums[0], nums[0]
	vals := make(map[int]struct{})
	for _, val := range nums {
		left = min(left, val)
		right = max(right, val)

		vals[val] = struct{}{}
	}

	res, curLen, cur := 1, 1, left+1
	for cur <= right {
		if _, ok := vals[cur]; !ok {
			curLen = 0
			if len(nums)-res <= res {
				return res
			}
		} else {
			curLen += 1
		}
		cur++
		res = max(res, curLen)
	}
	return res
}
