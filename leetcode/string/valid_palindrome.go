package string

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var newS strings.Builder
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
			newS.WriteString(string(s[i]))
		}
	}

	var j = newS.Len() - 1
	for i := 0; i < newS.Len()/2; i++ {
		if newS.String()[i] != newS.String()[j] {
			return false
		}
		j--
	}

	return true
}
