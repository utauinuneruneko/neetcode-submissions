func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)

	var caculate func(s string) string
	caculate = func(s string) string {
		nums := make([]int, 26)
		for _, r := range []rune(s) {
			nums[r-'a']++
		}
		return fmt.Sprintf("%v", nums)
	}

	for _, str := range strs {
		key := caculate(str)
		if _, ok := dict[key]; !ok {
			dict[key] = make([]string, 0)
		} 
		dict[key] = append(dict[key], str)
	}

	results := make([][]string, 0)
	for _, values := range dict {
		results = append(results, append([]string{}, values...))
	}
	return results
}
