func canJump(nums []int) bool {
    n := len(nums)
	arr := make([]int, n)
	for i, val := range nums {
		arr[i] = val+i // farest position can reach from current pos 
	}
	for i := 0; i < n; {
		l, r := 1, arr[i]
		if r >= n-1 {
			return true
		}
		next := i
		for j := l; j <= r; j++ {
			if arr[next] < arr[j] {
				next = j
			}
		}
		if next == i {
			return false
		}
		i = next
	}
	return false
}
