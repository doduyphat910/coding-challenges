package daily

import "regexp"

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	var isCap bool
	if isCapital(word[1]) {
		isCap = true
	}
	if isCap && !isCapital(word[0]) {
		return false
	}

	for i := 1; i < len(word); i++ {
		if !isCapital(word[i]) && isCap {
			return false
		}
		if isCapital(word[i]) && !isCap {
			return false
		}
	}
	return true
}

func isCapital(s byte) bool {
	reg, _ := regexp.Compile("[A-Z]")
	return reg.Match([]byte{s})
}
