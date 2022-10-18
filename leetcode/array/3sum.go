package array

import "sort"

func threeSum(nums []int) [][]int {
	resultMap := make(map[[3]int]bool, 0)
	results := make([][]int, 0)

	remainByKeysMap := make(map[int][]int)
	for i := range nums {
		keys, ok := remainByKeysMap[nums[i]]
		if !ok {
			keys = make([]int, 0)
		}
		keys = append(keys, i)
		remainByKeysMap[nums[i]] = keys
	}

	for i := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			remain := 0 - nums[i] - nums[j]
			keys, ok1 := remainByKeysMap[remain]
			if !ok1 {
				continue
			}

			for k := range keys {
				if keys[k] > j {
					var result = []int{nums[keys[k]], nums[i], nums[j]}
					sort.Ints(result)
					if _, ok := resultMap[[3]int{result[0], result[1], result[2]}]; !ok {
						resultMap[[3]int{result[0], result[1], result[2]}] = true
						results = append(results, result)
					}
				}
			}
		}
	}

	return results
}
