func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtrace func(int)
	backtrace = func(index int) {
		res = append(res, append([]int{}, path...))
		if index == len(nums) {
			return
		}
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrace(i+1)
			path = path[:len(path)-1]
		}
	}
	backtrace(0)
	return res
}
