package dp

func combinationSum4(nums []int, target int) int {
	var dpMap = make(map[int]int, target)
	dpMap[0] = 1

	for i := 0; i <= target; i++ {
		for j := range nums {
			dpMap[i] += dpMap[i-nums[j]]
		}
	}
	return dpMap[target]
}
