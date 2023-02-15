package dp

func lengthOfLIS(nums []int) int {
	var (
		subByIndexMap = make(map[int]int)
		max           = 1
	)
	for i := len(nums) - 1; i >= 0; i-- {
		subByIndexMap[i] = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] >= nums[j] {
				continue
			}
			if 1+subByIndexMap[j] <= subByIndexMap[i] {
				continue
			}
			subByIndexMap[i] = 1 + subByIndexMap[j]
			if subByIndexMap[i] > max {
				max = subByIndexMap[i]
			}
		}
	}

	return max
}
