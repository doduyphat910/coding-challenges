package array

import "math"

func maxProduct(nums []int) int {
	var (
		max        = math.MinInt
		totalLeft  = 1
		totalRight = 1
	)
	maxs := maxSubProduct(nums, 0, totalLeft, totalRight)
	for i := range maxs {
		if maxs[i] > max {
			max = maxs[i]
		}
	}
	return max
}

func maxSubProduct(nums []int, i int, totalLeft int, totalRight int) (maxs []int) {
	totalLeft *= nums[i]
	totalRight *= nums[len(nums)-1-i]
	maxs = append(maxs, totalLeft)
	maxs = append(maxs, totalRight)
	if totalLeft == 0 {
		totalLeft = 1
	}
	if totalRight == 0 {
		totalRight = 1
	}

	if i < len(nums)-1 {
		i++
		maxsSub := maxSubProduct(nums, i, totalLeft, totalRight)
		maxs = append(maxs, maxsSub...)
	}

	return maxs
}

// Best solution
func maxProduct(nums []int) int {
	var (
		max        = math.MinInt
		totalLeft  = 1
		totalRight = 1
		maxs       []int
	)
	for i := range nums {
		totalLeft *= nums[i]
		totalRight *= nums[len(nums)-1-i]
		maxs = append(maxs, totalLeft)
		maxs = append(maxs, totalRight)
		if totalLeft == 0 {
			totalLeft = 1
		}
		if totalRight == 0 {
			totalRight = 1
		}
	}

	for i := range maxs {
		if maxs[i] > max {
			max = maxs[i]
		}
	}
	return max
}
