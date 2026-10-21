func combinationSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
    res := make([][]int, 0)
	path := make([]int, 0)
	var backtrace func(int, int)
	backtrace = func(index int, value int) {
		if value == target {
			res = append(res, append([]int{}, path...))
			return
		} else if value > target {
			return
		}
		for i := index; i < len(nums); i++ {
			if value + nums[i] > target {
				return
			}
			path = append(path, nums[i])
			backtrace(i, value+nums[i])
			path = path[:len(path)-1]
		}
	}
	backtrace(0, 0)
	return res
}
