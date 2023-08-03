package neetcode

func partitionLabels(s string) []int {
	var (
		size, end       int
		result          = make([]int, 0)
		labelByIndexMap = make(map[string]int, 0)
	)
	for i := range s {
		labelByIndexMap[string(s[i])] = i
	}
	for i := range s {
		index, _ := labelByIndexMap[string(s[i])]
		size += 1
		if end < index {
			end = index
		}
		if i == end {
			result = append(result, size)
			size = 0
		}

	}
	return result
}
