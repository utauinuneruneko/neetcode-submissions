func lengthOfLIS(nums []int) int {
	//card game manipulate
	n := len(nums)
	stack := make([]int, 0)

	var find func(int) int
	find = func(num int) int {
		left, right := 0, len(stack)-1
		for left <= right {
			mid := (right-left)/2+left
			if stack[mid] < num {
				left = mid+1
			} else {
				right = mid-1
			}
		}
		return right+1
	}

	for i := 0; i < n; i++ {
		index := find(nums[i])
		if index < len(stack) {
			stack[index] = nums[i]
		} else {
			stack = append(stack, nums[i])
		}
	}

	return len(stack)
}
