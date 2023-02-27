package daily

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, i, j, target int) int {
	if j-i <= 1 {
		if nums[j] < target {
			return j + 1
		}
		if nums[i] < target {
			return i + 1
		}
		if i-1 <= 0 {
			return 0
		}
		return i - 1
	}

	mid := (i + j) / 2
	if nums[mid] == target {
		return mid
	}
	var result int
	if nums[mid] > target {
		result = binarySearch(nums, i, mid, target)
	} else {
		result = binarySearch(nums, mid, j, target)
	}

	return result
}
