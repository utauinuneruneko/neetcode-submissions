func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
	totalSum := 0
	arr := make([]int, n)
	for i := range gas {
		arr[i] = gas[i]-cost[i]
		totalSum += arr[i]
	}
	if totalSum < 0 {
		return -1
	}
	start := 0
	curSum := 0
	for i := 0; i < n; i++{
		if curSum + arr[i] < 0 {
			start = i+1
			curSum = 0
		} else {
			curSum += arr[i]
		}
	}
	return start
}
