func productExceptSelf(nums []int) []int {
	var (
		results                 []int
		totalProductWithZero    = 1
		totalProductWithoutZero = 1
		moreZero                int
	)
	for i := range nums {
		if nums[i] == 0 {
			moreZero++
		}
		if nums[i] != 0 {
			totalProductWithZero *= nums[i]
		}
		totalProductWithoutZero *= nums[i]
	}

	for i := range nums {
		var result int
		if nums[i] == 0 {
			result = totalProductWithZero
		} else {
			result = totalProductWithoutZero / nums[i]
		}

		if moreZero > 1 {
			result = 0
		}
		results = append(results, result)
	}

	return results

	// var results []int
	// var result = 1
	// for i := range nums {
	//     for j := range nums {
	//         if i != j {
	//             result *= nums[j]
	//         }
	//         if j == len(nums) -1 {
	//             results = append(results, result)
	//             result = 1
	//         }
	//     }
	// }

	// return results
}