package seanprashad

func singleNumber(nums []int) int {
	var sum int
	for i := range nums {
		sum ^= nums[i]
	}
	return sum
}
