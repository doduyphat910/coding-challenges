package seanprashad

import (
	"strings"
	"unicode"
)

func letterCasePermutation(s string) []string {
	var permutation = make([]string, 0)

	if unicode.IsLetter(rune(s[0])) {
		permutation = append(permutation, strings.ToLower(string(s[0])))
		permutation = append(permutation, strings.ToUpper(string(s[0])))
	} else {
		permutation = append(permutation, string(s[0]))
	}

	for i := 1; i < len(s); i++ {
		var tmp = make([]string, 0)
		if unicode.IsLetter(rune(s[i])) {
			for j := range permutation {
				tmp = append(tmp, permutation[j]+strings.ToLower(string(s[i])))
				tmp = append(tmp, permutation[j]+strings.ToUpper(string(s[i])))
			}
		} else {
			for j := range permutation {
				tmp = append(tmp, permutation[j]+string(s[i]))
			}
		}
		permutation = tmp
	}

	return permutation
}
