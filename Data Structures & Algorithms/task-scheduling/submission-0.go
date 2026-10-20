func leastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	maxCnt := 0
	for _, task := range tasks {
		index := int(task-'A')
		count[index] += 1
		maxCnt = max(maxCnt, count[index])
	}

	maxTask := 0
	for _, val := range count {
		if val == maxCnt {
			maxTask += 1
		}
	}

	res := maxTask + (maxCnt-1) * (n+1)
	if res < len(tasks) {
		return len(tasks)
	}
	return res
}
