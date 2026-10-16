func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (right-left)/2+left
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid+1
		}
	}
	return nums[right]
}
