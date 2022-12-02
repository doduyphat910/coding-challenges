package string

func lengthOfLongestSubstring(s string) int {
	var (
		positionByStringMap = make(map[string]int)
		max                 = 0
		start               = -1
	)

	for i := range s {
		position, ok := positionByStringMap[string(s[i])]
		if !ok {
			position = -1
		}

		if position > start {
			start = position
		}
		if (i - start) > max {
			max = i - start
		}

		positionByStringMap[string(s[i])] = i
	}

	return max
}
