func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtrace func(int)
	backtrace = func(index int) {
		res = append(res, append([]int{}, path...))
		if index == len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrace(i+1)
			path = path[:len(path)-1]
		}
	}
	backtrace(0)
	return res
}
