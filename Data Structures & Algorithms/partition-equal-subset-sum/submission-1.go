func canPartition(nums []int) bool {
    sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 != 0 {
		return false
	}
	target := sum/2
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	memo := make(map[int]map[int]bool)

	var dfs func(int, int) bool
	dfs = func(curSum, index int) bool {
		if _, ok1 := memo[curSum]; ok1 {
			if _, ok2 := memo[curSum][index]; ok2 {
				return memo[curSum][index]
			}
		} else {
			memo[curSum] = make(map[int]bool)
		}

		if curSum > target {
			return false
		}


		if index == len(nums) {
			memo[curSum][index] = curSum == target
			return memo[curSum][index]
		}

		memo[curSum][index] = dfs(curSum, index+1) || dfs(curSum+nums[index], index+1)
		return memo[curSum][index]
	}
	return dfs(0, 0)
}
