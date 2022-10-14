func containsDuplicate(nums []int) bool {
	numByAppear := make(map[int]uint)
	for i := range nums {
		numByAppear[nums[i]] += 1
		if numByAppear[nums[i]] == 2 {
			return true
		}
	}

	return false
}