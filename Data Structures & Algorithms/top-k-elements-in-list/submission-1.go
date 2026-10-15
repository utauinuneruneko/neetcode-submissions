
func topKFrequent(nums []int, k int) []int {
	numToCnt := make(map[int]int)
	for _, num := range nums {
		numToCnt[num]++
	}
	cntToNums := make(map[int][]int)
	minCnt, maxCnt := len(nums), 0
	for num, cnt := range numToCnt {
		minCnt = min(minCnt, cnt)
		maxCnt = max(maxCnt, cnt)
		if _, ok := cntToNums[cnt]; !ok {
			cntToNums[cnt] = make([]int, 0)
		}
		cntToNums[cnt] = append(cntToNums[cnt], num)
	}
	res := make([]int, k)
	for i, cnt := 0, maxCnt; i < k && cnt >= minCnt; {
		if curNums, ok := cntToNums[cnt]; ok {
			for j := 0; j < len(curNums) && i < k; j++ {
				res[i] = curNums[j]
				i++
			}
		}
		cnt--
	}
	return res
}
