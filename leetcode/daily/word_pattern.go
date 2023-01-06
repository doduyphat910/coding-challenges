package daily

import "strings"

func wordPattern(pattern string, s string) bool {
	splittedS := strings.Split(s, " ")
	if len(splittedS) != len(pattern) {
		return false
	}

	var subByPatternMap = make(map[byte]string)
	var patternBySubMap = make(map[string]byte)
	for i := range pattern {
		_, ok := subByPatternMap[pattern[i]]
		if !ok {
			subByPatternMap[pattern[i]] = splittedS[i]
		}
		if subByPatternMap[pattern[i]] != splittedS[i] {
			return false
		}

		_, ok = patternBySubMap[splittedS[i]]
		if !ok {
			patternBySubMap[splittedS[i]] = pattern[i]
		}

		if pattern[i] != patternBySubMap[splittedS[i]] {
			return false
		}

	}
	return true
}
