package daily

func minDeletionSize(strs []string) int {
	var count int
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				count += 1
				break
			}
		}
	}

	return count
}
