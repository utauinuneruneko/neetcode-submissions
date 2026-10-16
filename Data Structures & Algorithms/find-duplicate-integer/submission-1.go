func findDuplicate(nums []int) int {
    slow, fast := 0, 0
    
    for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = 0
	for {
		slow = nums[slow]
		fast = nums[fast]
		if slow == fast {
			break
		}
	}
	return slow
}
