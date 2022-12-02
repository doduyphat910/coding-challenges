package string

func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	var numsByAppearMap = make(map[string]int)
	for i := 0; i < len(s); i++ {
		nums := numsByAppearMap[string(s[i])]
		nums += 1
		numsByAppearMap[string(s[i])] = nums

		nums = numsByAppearMap[string(t[i])]
		nums -= 1
		numsByAppearMap[string(t[i])] = nums
	}

	for i := range numsByAppearMap {
		if numsByAppearMap[i] != 0 {
			return false
		}
	}

	return true
}
