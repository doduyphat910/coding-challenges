package seanprashad

import "math"

func findDisappearedNumbers(nums []int) []int {
	var result = make([]int, 0)
	for i := range nums {
		nums[int(math.Abs(float64(nums[i])))-1] = -int(math.Abs(float64(nums[int(math.Abs(float64(nums[i])))-1])))
	}
	for i := range nums {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}
