func maxSubArray(nums []int) int {
    n := len(nums)
	arr := make([]int, n) // end as nums[i] maximum sum
	arr[0] = nums[0]
	res := nums[0]
	for i := 1; i < n; i++ {
		if arr[i-1] < 0 {
			arr[i] = nums[i]
		} else {
			arr[i] = nums[i] + arr[i-1]
		}
		res = max(res, arr[i])
	}
	return res
}
