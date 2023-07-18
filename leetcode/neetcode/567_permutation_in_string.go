package neetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1Map = make(map[byte]int)
	var s2Map = make(map[byte]int)
	for i := range s1 {
		s1Map[s1[i]]++
		s2Map[s2[i]]++
	}

	for i := len(s1); i < len(s2); i++ {
		if compareMap(s1Map, s2Map) {
			return true
		}

		s2Map[s2[i-len(s1)]]--
		s2Map[s2[i]]++
	}

	return compareMap(s1Map, s2Map)
}

func compareMap(map1 map[byte]int, map2 map[byte]int) bool {
	for i := range map1 {
		if map1[i] != map2[i] {
			return false
		}
	}

	return true
}
