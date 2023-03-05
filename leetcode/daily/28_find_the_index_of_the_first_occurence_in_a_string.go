package daily

func strStr(haystack string, needle string) int {
	for i := range haystack {
		for j := 0; j < len(needle); j++ {
			if i+j >= len(haystack) || haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1
}
