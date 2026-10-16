func findMin(nums []int) int {
	n := len(nums)
	var biSearch func(int, int) int
	biSearch = func(left int, right int) int {
		if left == right {
			return nums[left]
		}
		if nums[left] < nums[right] {
			return nums[left]
		}

		mid := (right-left)/2+left
		return min(biSearch(left, mid), biSearch(mid+1, right))

	}
	return biSearch(0, n-1)
}
