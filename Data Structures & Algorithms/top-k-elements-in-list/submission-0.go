
func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	for _, num := range nums {
		dict[num]++
	}
	tuples := make([][]int, len(dict))
	i := 0
	for num, cnt := range dict {
		tuples[i] = []int{num, cnt}
		i++
	}
	sort.Slice(tuples, func(i, j int) bool {
		if tuples[i][1] == tuples[j][1] {
			return tuples[i][0] < tuples[j][0]
		}
		return tuples[i][1] > tuples[j][1]
	})
	results := make([]int, 0)
	for i = 0; i < k; i++ {
		results = append(results, tuples[i][0])
	}
	return results
}
