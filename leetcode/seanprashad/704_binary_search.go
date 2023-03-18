package seanprashad

func search(nums []int, target int) int {
	var (
		l = 0
		r = len(nums) - 1
	)

	for l < r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		}

		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}

	if nums[l] == target {
		return l
	}
	return -1
}
