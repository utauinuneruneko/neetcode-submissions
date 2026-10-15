func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	vals := make(map[int]struct{})
	for _, val := range nums {
		vals[val] = struct{}{}
	}

	res := 0
	for _, val := range nums {
		if _, ok := vals[val]; !ok {
			continue
		}
		delete(vals, val)
		cnt := 1
		left, right := val-1, val+1
		for {
			if _, ok := vals[left]; ok {
				delete(vals, left)
				left--
				cnt++
			} else {
				break
			}
		}
		for {
			if _, ok := vals[right]; ok {
				delete(vals, right)
				right++
				cnt++
			} else {
				break
			}
		}
		res = max(res, cnt)
	}
	return res
}
