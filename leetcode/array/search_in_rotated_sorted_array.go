package array

func search(nums []int, target int) int {
	var (
		l   = 0
		mid = 0
		r   = len(nums) - 1
	)

	for l < r {
		mid = (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	s := l
	l = 0
	r = len(nums) - 1
	if target >= nums[s] && target <= nums[r] {
		l = s
	} else {
		r = s
	}

	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}
