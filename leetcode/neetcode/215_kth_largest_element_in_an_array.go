package neetcode

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, l, r, k int) int {
	pivot, p := nums[r], l
	for i := l; i < r; i++ {
		if nums[i] <= pivot {
			nums[i], nums[p] = nums[p], nums[i]
			p += 1
		}
	}
	nums[r], nums[p] = nums[p], nums[r]

	if p == k {
		return nums[p]
	}
	if p > k {
		return quickSelect(nums, l, p-1, k)
	}
	return quickSelect(nums, p+1, r, k)
}
