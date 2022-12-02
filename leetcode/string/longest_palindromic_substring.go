package string

func longestPalindrome(s string) string {
	var result string
	for i := range s {
		maxString1 := maxPalindrome(s, i, i-1)
		if len(maxString1) > len(result) {
			result = maxString1
		}
		maxString2 := maxPalindrome(s, i+1, i-1)
		if len(maxString2) > len(result) {
			result = maxString2
		}
	}

	return result
}

func maxPalindrome(s string, i int, k int) string {
	if i <= 0 || k < 0 || i > len(s)-1 {
		return string(s[0])
	}

	if s[i] != s[k] {
		return string(s[i])
	}
	for i < len(s)-1 {
		if k == 0 || s[i+1] != s[k-1] {
			break
		}
		i++
		k--
	}

	return string(s[k : i+1])
}
