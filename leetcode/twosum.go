func twoSum(nums []int, target int) []int {
	var result []int
	// for i := range nums {
	//     for j := i + 1; j < len(nums); j++ {
	//         if nums[i] + nums[j] == target {
	//             result = append(result, i)
	//              result = append(result, j)
	//         }
	//     }
	// }
	// return result

	numsByKeyMap := make(map[int]int, len(nums))
	for i := range nums {
		numsByKeyMap[nums[i]] = i
	}

	for i := range nums {
		remain := target - nums[i]
		key, ok := numsByKeyMap[remain]
		if ok && key != i {
			result = append(result, i)
			result = append(result, key)

			// delete(numsByKeyMap, remain)
			delete(numsByKeyMap, nums[i])
		}
	}

	return result
}