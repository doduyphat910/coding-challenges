package graph

import "sort"

// brute-force:
// Time: O(NlogN), Space: O(1)
func longestConsecutive(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var result int
	var max = 0
	if len(nums) > 0 {
		max = 1
	}
	for i := range nums {
		if i < len(nums)-1 && nums[i+1] == nums[i]+1 {
			max += 1
			if max > result {
				result = max
			}
			continue
		} else if i < len(nums)-1 && nums[i+1] == nums[i] {
			continue
		}
		max = 1
	}
	if max > result {
		result = max
	}

	return result
}

// optimize
// Time: O(N), Space: O(N)
func longestConsecutive(nums []int) int {
	var (
		numMap = make(map[int]bool)
		result int
		max    int
	)
	for i := range nums {
		if !numMap[nums[i]] {
			numMap[nums[i]] = true
		}
	}

	for i := range nums {
		max = 1
		if !numMap[nums[i]-1] {
			j := 1
			for numMap[nums[i]+j] {
				j++
			}
			if j > result {
				result = j
			}
		}
	}

	if max > result {
		result = max
	}

	return result
}
