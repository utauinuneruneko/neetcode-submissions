func jump(nums []int) int {
    n := len(nums)
	if n == 1 {
		return 0
	}

	arr := make([]int, n)
	for i, val := range nums {
		arr[i] = i+val
	}
	count := 0
	for i := 0; ; {
		count++
		l, r := i, arr[i]
		if r >= n-1 {
			return count
		}
		next := i
		for j := l; j <= r; j++ {
			if arr[j] > arr[next] {
				next = j
			}
		}
		if next == i {
			return -1
		}
		i = next
	}
	return -1
}
