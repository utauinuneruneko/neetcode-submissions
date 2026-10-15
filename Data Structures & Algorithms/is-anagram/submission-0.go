func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	nums := make([]int, 26)
	for _, r := range []rune(s) {
		nums[r-'a']++
	}
	for _, r := range []rune(t) {
		key := r-'a'
		if nums[key] == 0 {
			return false
		}
		nums[key]--
	}
	return true
}
