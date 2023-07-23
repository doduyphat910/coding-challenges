package neetcode

func partition(s string) [][]string {
	var res = make([][]string, 0)
	find(0, s, []string{}, &res)
	return res
}

func find(idx int, s string, list []string, res *[][]string) {
	if idx == len(s) {
		var temp = make([]string, len(list))
		copy(temp, list)
		*res = append(*res, temp)
		return
	}

	for i := idx; i < len(s); i++ {
		if isPalindrome(string(s[idx : i+1])) {
			find(i+1, s, append(list, string(s[idx:i+1])), res)
		}
	}
}

func isPalindrome(s string) bool {
	var j = len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}
