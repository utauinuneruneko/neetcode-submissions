func rob(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return max(nums[0], nums[n-1])
	}

	res := max(nums[0], nums[1])
    a, b := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		tmp := max(nums[i]+a, b)
		res = max(res, tmp)
		a = b
		b = tmp
	}
	return res
}
