package string

func countSubstrings(s string) int {
	var result int

	for i := range s {
		result += 1

		number := countPalindromicSubstring(s, i, i+1)
		result += number

		number2 := countPalindromicSubstring(s, i-1, i+1)
		result += number2
	}

	return result
}

func countPalindromicSubstring(s string, i, j int) int {
	var count int
	for i >= 0 && j <= len(s)-1 && s[i] == s[j] {
		count += 1
		i--
		j++
	}

	return count
}
