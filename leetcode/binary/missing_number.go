package binary

func missingNumber(nums []int) int {
	//     var numSet = make(map[int]bool, len(nums))
	//     for i := range nums {
	//         numSet[nums[i]] = true
	//     }

	//     var max int
	//     for i := range nums {
	//         if _, ok := numSet[i]; !ok {
	//             return i
	//         }
	//         if nums[i] > max {
	//             max = nums[i]
	//         }
	//     }

	//     return max + 1

	var (
		total = len(nums)
	)

	for i := range nums {
		total ^= nums[i]
		total ^= i
	}

	return total
}
