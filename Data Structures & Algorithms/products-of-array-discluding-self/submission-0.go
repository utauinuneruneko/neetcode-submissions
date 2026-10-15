func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	cnt := 0
	for _, v := range nums {
		if v == 0 {
			cnt++
		}
		if cnt >= 2 {
			break
		}
	}
	// more than two zero
	if cnt >= 2 {
		return res
	} else if cnt == 1 {
		// one zero
		tmp := 1
		index := 0
		for i, val := range nums {
			if val != 0 {
				tmp *= val
			} else {
				index = i
			}
		}
		res[index] = tmp
		return res
	}
	// no zero
	overFlag := false
	total := 1
	for i, val := range nums {
		if i < n-1 {
			total *= val
		} else {
			// over 32-bit
			if int64(total) * int64(val) > math.MaxInt32 {
				overFlag = true
			} else {
				// less than 32-bit
				total *= val
			}
		}
	}

	for i, val := range nums {
		if overFlag {
			res[i] = total/val*nums[n-1]
		} else {
			res[i] = total/val
		}
	}
	return res

}
