func countSubstrings(s string) int {
    count := 0
	n := len(s)
	for i := 0; i < n; i++ {
		for a, b := i, i; a >= 0 && b < n; a, b = a-1, b+1 {
			if s[a] != s[b] {
				break
			}
			count++
		}

		if i == n-1 {
			continue
		}
		for a, b := i, i+1; a >= 0 && b < n; a, b = a-1, b+1 {
			if s[a] != s[b] {
				break
			}
			count++
		}
	}
	return count
}
