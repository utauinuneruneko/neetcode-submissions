func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtrace func(int, int)
	backtrace = func(index int, sum int) {
		if sum == target {
			res = append(res ,append([]int{}, path...))
			return
		} else if sum > target {
			return
		}

		for i := index; i <len(candidates); i++ {
			if sum + candidates[i] > target {
				return
			}
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtrace(i+1, sum+candidates[i])
			path = path[:len(path)-1]

		}
	}
	backtrace(0, 0)
	return res
}
