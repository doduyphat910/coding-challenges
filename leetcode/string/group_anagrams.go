package string

import "sort"

func groupAnagrams(strs []string) [][]string {
	var (
		subStrsBySortedSMap = make(map[string][]string)
		result              = make([][]string, 0)
	)

	for i := range strs {
		s := []byte(strs[i])
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		subStrs, ok := subStrsBySortedSMap[string(s)]
		if ok {
			subStrs = append(subStrs, strs[i])
			subStrsBySortedSMap[string(s)] = subStrs
			continue
		}
		subStrs = []string{strs[i]}
		subStrsBySortedSMap[string(s)] = subStrs
	}

	for i := range subStrsBySortedSMap {
		result = append(result, subStrsBySortedSMap[i])
	}

	return result
}
