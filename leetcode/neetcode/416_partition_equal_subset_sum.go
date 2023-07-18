package neetcode

func canPartition(nums []int) bool {
	var total int
	for i := range nums {
		total += nums[i]
	}
	if total%2 != 0 {
		return false
	}

	var sumMap = make(map[int]bool)
	sumMap[nums[0]] = true
	backtracking(1, nums, &sumMap)

	return sumMap[total/2]
}

func backtracking(i int, nums []int, sumMap *map[int]bool) map[int]bool {
	if i >= len(nums) {
		return *sumMap
	}

	var tempMap = make(map[int]bool)
	for j := range *sumMap {
		tempMap[j] = true
		tempMap[nums[i]] = true
		tempMap[j+nums[i]] = true
	}
	*sumMap = tempMap
	i += 1
	return backtracking(i, nums, sumMap)
}

// second way
func canPartition2(nums []int) bool {
	var total int
	for i := range nums {
		total += nums[i]
	}
	if total%2 != 0 {
		return false
	}

	var sumMap = make(map[int]bool)
	sumMap[0] = true
	for i := range nums {
		var temp = make(map[int]bool)
		for j := range sumMap {
			temp[j] = true
			temp[j+nums[i]] = true
		}
		sumMap = temp
	}

	return sumMap[total/2]

}
