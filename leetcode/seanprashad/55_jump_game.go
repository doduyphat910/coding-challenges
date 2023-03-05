package seanprashad

func canJump(nums []int) bool {
	var max int
	for i := range nums {
		if max < i {
			return false
		}
		if max < i+nums[i] {
			max = i + nums[i]
		}
	}

	return true
}
