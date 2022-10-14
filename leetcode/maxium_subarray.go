func maxSubArray(nums []int) int {
	var (
		maxSoFar  = math.MinInt64
		maxEnding = 0
	)

	for i := range nums {
		maxEnding = maxEnding + nums[i]
		if maxSoFar < maxEnding {
			maxSoFar = maxEnding
		}
		if maxEnding < 0 {
			maxEnding = 0
		}
	}

	return maxSoFar
}

/*func maxSubArray(nums []int) int {
	var maxs []int
	for i := range nums {
		max := totalSubArray(nums, i)
		maxs = append(maxs, max)
	}

	var maxAll int
	if len(maxs) > 0 {
		maxAll = maxs[0]
	}
	for i := range maxs {
		if maxs[i] > maxAll {
			maxAll = maxs[i]
		}
	}

	return maxAll
}

func totalSubArray(nums []int, mid int) int {
	var (
		maxLeft  []int
		maxRight []int
		left     int
		right    int
	)

	for i := 0; i <= mid; i++ {
		left += nums[i]
		maxLeft = append(maxLeft, left)
	}

	for i := mid; i < len(nums); i++ {
		right += nums[i]
		maxRight = append(maxRight, right)
	}

	var max int
	if len(maxLeft) > 0 {
		max = maxLeft[0]
	}

	for i := range maxLeft {
		if maxLeft[i] > max {
			max = maxLeft[i]
		}
	}
	for i := range maxRight {
		if maxRight[i] > max {
			max = maxRight[i]
		}
	}

	return max
}*/