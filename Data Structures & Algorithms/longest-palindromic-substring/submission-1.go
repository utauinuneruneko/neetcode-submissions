func longestPalindrome(s string) string {
    n := len(s)
	l, r, res := 0, 0, 1
	for i := 0; i <n; i++ {
		// odd case
		for a, b := i, i; a >= 0 && b < n; a, b = a-1, b+1 {
			if s[a] != s[b] {
				break
			}
			if b-a+1 > res {
				l, r = a, b
				res = b-a+1
			}
		}

		// even case
		if i == n-1 {
			continue
		}
		for a, b := i, i+1; a >= 0 && b < n; a, b = a-1, b+1 {
			if s[a] != s[b] {
				break
			}
			if b-a+1 > res {
				l, r = a, b
				res = b-a+1
			}
		}
	}
	return s[l:r+1]
}
