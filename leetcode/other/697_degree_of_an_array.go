package other

import "math"

func findShortestSubArray(nums []int) int {
	var (
		numByKeysMap = make(map[int][]int)
		maxDegree    = math.MinInt
	)
	for i := range nums {
		keys, ok := numByKeysMap[nums[i]]
		if !ok {
			keys = make([]int, 0)
		}
		keys = append(keys, i)
		numByKeysMap[nums[i]] = keys
		if len(numByKeysMap[nums[i]]) > maxDegree {
			maxDegree = len(numByKeysMap[nums[i]])
		}
	}

	var min = math.MaxInt
	for i := range numByKeysMap {
		if len(numByKeysMap[i]) == maxDegree {
			if numByKeysMap[i][len(numByKeysMap[i])-1]-numByKeysMap[i][0]+1 < min {
				min = numByKeysMap[i][len(numByKeysMap[i])-1] - numByKeysMap[i][0] + 1
			}
		}
	}

	return min
}
